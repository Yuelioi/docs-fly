package database

import (
	"docsfly/global"
	"docsfly/models"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
)

type FileInfo struct {
	Length   int
	FileType string
	Document models.Document
	Category models.Category
}

type FileContainer struct {
	Length    int               `json:"-"`
	Folder    string            `json:"-"`
	Documents []models.Document `json:"documents"`
	Categorys []models.Category `json:"categorys"`
}

type MetaOutput struct {
	Documents []models.MetaData `json:"documents"`
	Categorys []models.MetaData `json:"categorys"`
}

type Stack struct {
	elements []FileContainer
}

func newStack() *Stack {
	return &Stack{elements: []FileContainer{}}
}

func (s *Stack) Push(element FileContainer) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Add(element FileInfo) {
	if len(s.elements) == 0 {
		return
	}

	lastContainer := &s.elements[len(s.elements)-1]

	if element.FileType == "category" {
		lastContainer.Categorys = append(lastContainer.Categorys, element.Category)
	} else if element.FileType == "document" {
		lastContainer.Documents = append(lastContainer.Documents, element.Document)
	}

}

func (s *Stack) Pop() *FileContainer {
	if len(s.elements) == 0 {
		return nil
	}
	element := &s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

/*
初始化数据库: 仅在无数据库, 或者手动修改meta信息后重置数据库

	@params:root
		markdown文件存放位置
	@params:Mode
	1: init模式,第一次创建数据库, 会创建meta信息
	2: update模式, 会按照目录下的meta.json读取设置
*/
func DBInit(db *gorm.DB) error {
	fmt.Println("初始化数据库准备中...")

	counterMap := make(map[string]uint)
	root := global.AppConfig.Resource
	Mode := global.AppConfig.DBMode

	if Mode == 0 {
		println("注意!!!当前为初始化模式, 并且会修改文件内容")
	} else if Mode <= 1 {
		println("当前为初始化模式,只会生成meta.json")

	} else if Mode == 2 {
		println("当前为更新模式: 会基于本地meta信息修改")
	}

	start := time.Now()

	// 写入管理员数据
	CreateAdminAccount(db)

	// 存储各个类目总数据 用于写入数据库
	cats := make([]models.Category, 0)

	docs := make([]models.Document, 0)

	var lastLength int = -1
	var catCount uint = 0

	s := newStack()

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == root {
			return nil
		}

		if info.IsDir() && (info.Name() == ".git" || info.Name() == ".vscode" || info.Name() == "Ue") {
			return filepath.SkipDir
		}

		if info.Name() == "meta.json" {
			return nil
		}

		path = strings.ReplaceAll(path, root+"\\", "")
		length := strings.Count(path, "\\")

		fmt.Println(catCount, lastLength, length, info.Name(), path)

		metaData := models.MetaData{
			Name:     info.Name(),
			Title:    info.Name(),
			Filepath: path,
		}

		if value, exists := counterMap[filepath.Dir(path)]; exists {
			counterMap[filepath.Dir(path)] = value + 1
			metaData.Order = value + 1
		} else {
			metaData.Order = 1
		}

		fileInfo := FileInfo{
			Length: length,
		}

		if info.IsDir() {
			// 文件夹赋值初始顺序为1
			counterMap[path] = 0
			fileInfo.FileType = "category"

			catCount++

			cat := models.Category{
				MetaData: metaData,
				ModTime:  info.ModTime(),
				Display:  true,
			}
			cats = append(cats, cat)
			fileInfo.Category = cat
		} else {
			fileInfo.FileType = "document"

			doc := models.Document{
				MetaData:   metaData,
				ModTime:    info.ModTime(),
				Locale:     "",
				Content:    "",
				Hash:       "",
				Size:       uint(info.Size()),
				CategoryID: catCount,
			}

			docs = append(docs, doc)
			fileInfo.Document = doc
		}

		if length > lastLength {
			fc := FileContainer{
				Folder:    filepath.Dir(root + "\\" + path),
				Length:    length,
				Documents: []models.Document{},
				Categorys: []models.Category{},
			}
			s.Push(fc)
			s.Add(fileInfo)
		}

		if length == lastLength {
			s.Add(fileInfo)
		}
		// 跳出层级
		if length < lastLength {

			for {

				countainer := s.Pop()

				WriteMetaData(*countainer)
				if countainer.Length-1 == length {
					s.Add(fileInfo)
					break
				}

			}

		}
		lastLength = length
		return nil
	})

	if err != nil {
		return err
	}

	// 把剩余的(根目录)的meta.json写出来
	for {
		container := s.Pop()
		if container != nil {
			WriteMetaData(*container)
		}
		break
	}

	fmt.Println("正在写入数据库...")

	WriteContentToDocsData(&docs)

	fmt.Println("保存数据中...")
	err = WriteIntoDatabase(db,
		interface{}(cats),
		interface{}(docs))
	if err != nil {
		return err
	}

	// fmt.Println("数据库生成成功")
	fmt.Println("用时", time.Since(start))

	if Mode <= 1 {

		global.AppConfig.DBMode = 2
		err := global.WriteConfigToFile("DBMode")
		if err != nil {
			println(err)
			fmt.Println("数据库模式切换失败,请手动切换")
		} else {
			fmt.Println("数据库生成模式已切换为更新模式")
		}

	}

	return nil
}
