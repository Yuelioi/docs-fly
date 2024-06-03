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
	Depth    int
	FileType string
	Document models.Document
	Category models.Category
}

type LocalMetaCache struct {
	Depth     int               `json:"-"`
	Folder    string            `json:"-"`
	Documents []models.Document `json:"documents"`
	Categorys []models.Category `json:"categorys"`
}

type LocalMeta struct {
	Documents []models.MetaData `json:"documents"`
	Categorys []models.MetaData `json:"categorys"`
}

type Stack struct {
	elements []LocalMetaCache
}

func newStack() *Stack {
	return &Stack{elements: []LocalMetaCache{}}
}

func (s *Stack) Push(element LocalMetaCache) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Add(element FileInfo) {
	if len(s.elements) == 0 {
		return
	}

	lastLocalMetaCache := &s.elements[len(s.elements)-1]

	if element.FileType == "category" {
		lastLocalMetaCache.Categorys = append(lastLocalMetaCache.Categorys, element.Category)
	} else if element.FileType == "document" {
		lastLocalMetaCache.Documents = append(lastLocalMetaCache.Documents, element.Document)
	}

}

func (s *Stack) Pop() *LocalMetaCache {
	if len(s.elements) == 0 {
		return nil
	}
	element := &s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

/*
初始化数据库
*/
func DBInit(db *gorm.DB) error {
	fmt.Println("初始化数据库准备中...")

	// 记录所有层级序号
	orderMap := make(map[string]uint)
	root := global.AppConfig.Resource

	start := time.Now()

	// 如果数据库没有用户 则写入管理员数据

	var user models.User
	db.Model(models.User{}).Where("username =?", global.AppConfig.Username).Find(&user)

	if user.ID == 0 {
		CreateAdminAccount(db)
	}

	// 存储各个类目总数据 用于批量写入数据库
	cats := make([]models.Category, 0)
	docs := make([]models.Document, 0)
	localMetas := make([]LocalMetaCache, 0)

	// 上一文件深度
	var lastDepth int = -1
	var catCount uint = 0

	s := newStack()

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == root {
			return nil
		}

		if info.IsDir() && (info.Name() == ".git" || info.Name() == ".vscode") {
			return filepath.SkipDir
		}

		// if info.IsDir() && (info.Name() == "Ue") {
		// 	return filepath.SkipDir
		// }

		if info.Name() == "meta.json" {
			return nil
		}

		path = strings.ReplaceAll(path, root+"\\", "")
		Depth := strings.Count(path, "\\")

		metaData := models.MetaData{
			Name:     info.Name(),
			Title:    info.Name(),
			Depth:    Depth,
			Icon:     "",
			Status:   1,
			Filepath: path,
		}

		if value, exists := orderMap[filepath.Dir(path)]; exists {
			orderMap[filepath.Dir(path)] = value + 1
			metaData.Order = value + 1
		} else {
			metaData.Order = 1
		}

		fileInfo := FileInfo{
			Depth: Depth,
		}

		if info.IsDir() {
			// 文件夹赋值初始顺序为1
			orderMap[path] = 0
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

		if Depth > lastDepth {
			fc := LocalMetaCache{
				Folder:    filepath.Dir(root + "\\" + path),
				Depth:     Depth,
				Documents: []models.Document{},
				Categorys: []models.Category{},
			}
			s.Push(fc)
			s.Add(fileInfo)
		}

		if Depth == lastDepth {
			s.Add(fileInfo)
		}
		// 跳出层级
		if Depth < lastDepth {

			for {

				LocalMetaCache := s.Pop()

				localMetas = append(localMetas, *LocalMetaCache)

				if LocalMetaCache.Depth-1 == Depth {
					s.Add(fileInfo)
					break
				}

			}

		}
		lastDepth = Depth
		return nil
	})

	if err != nil {
		return err
	}

	// 把剩余的(根目录)的meta.json写出来
	for {
		LocalMetaCache := s.Pop()

		if LocalMetaCache == nil {
			break
		}
		localMetas = append(localMetas, *LocalMetaCache)
	}

	fmt.Println("读取数据用时", time.Since(start))

	start = time.Now()
	WriteLocalMetaData(localMetas)
	fmt.Println("写入meta数据用时", time.Since(start))

	// start = time.Now()
	// WriteContentToDocsData(&docs)
	// fmt.Println("读取内容用时", time.Since(start))

	// start = time.Now()
	// err = WriteIntoDatabase(db,
	// 	interface{}(cats),
	// 	interface{}(docs))
	// if err != nil {
	// 	return err
	// }

	// // fmt.Println("数据库生成成功")
	// fmt.Println("保存数据库用时", time.Since(start))

	return nil
}
