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

/*
初始化数据库
*/
func DBInit(db *gorm.DB) error {
	fmt.Println("初始化数据库准备中...")
	start := time.Now()

	// 如果数据库没有用户 则写入管理员数据
	var user models.User
	db.Model(models.User{}).Where("username =?", global.AppConfig.Username).Find(&user)

	if user.ID == 0 {
		CreateAdminAccount(db)
	}

	// 本地数据
	localOrderMap := make(map[string]uint)

	// 存储各个类目总数据 用于批量写入数据库
	localCats := make([]models.Category, 0)
	localDocs := make([]models.Document, 0)
	localMetas := make([]LocalMetaCache, 0)
	dbLocalMetas := make([]models.MetaDataLocal, 0)

	var dbCats []models.Category
	var dbDocs []models.Document
	db.Find(&dbCats)
	db.Find(&dbDocs)

	// dbCatsMap := make(map[string]models.Category)
	// dbDocsMap := make(map[string]models.Document)

	// 上一文件深度
	var lastDepth int = -1
	var catCount uint = 0

	s := &Stack{fileMetas: []LocalMetaCache{}}

	root := global.AppConfig.Resource

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

		if value, exists := localOrderMap[filepath.Dir(path)]; exists {
			localOrderMap[filepath.Dir(path)] = value + 1
			metaData.Order = value + 1
		} else {
			metaData.Order = 1
		}

		var fileInfo interface{}

		if info.IsDir() {
			// 文件夹赋值初始顺序为1
			localOrderMap[path] = 0

			catCount++

			cat := models.Category{
				MetaData: metaData,
				ModTime:  info.ModTime(),
				Display:  true,
			}
			localCats = append(localCats, cat)
			fileInfo = cat
		} else {

			doc := models.Document{
				MetaData:   metaData,
				ModTime:    info.ModTime(),
				Locale:     "",
				Content:    "",
				Hash:       "",
				Size:       uint(info.Size()),
				CategoryID: catCount,
			}

			localDocs = append(localDocs, doc)
			fileInfo = doc
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

	start = time.Now()
	WriteContentToDocsData(&localDocs)
	fmt.Println("读取内容用时", time.Since(start))

	for _, meta := range localMetas {
		ml := models.MetaDataLocal{
			Size:     meta.Size,
			Hash:     meta.Hash,
			Filepath: meta.Folder,
		}
		dbLocalMetas = append(dbLocalMetas, ml)
	}

	start = time.Now()
	err = WriteIntoDatabase(db,
		interface{}(localCats),
		interface{}(localDocs),
		interface{}(dbLocalMetas))

	// err = compareAndSync(db, localCats, localDocs)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	// fmt.Println("数据库生成成功")
	fmt.Println("保存数据库用时", time.Since(start))

	return nil
}
