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

type DBDatas struct {
	Cats struct {
		Creates []models.Category
		Updates []models.Category
		Deletes []models.Category
	}
	Docs struct {
		Creates []models.Document
		Updates []models.Document
		Deletes []models.Document
	}
}

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

	// 存储各个类目总数据 用于批量写入数据库
	dbDatas := DBDatas{}

	// 数据库映射表
	var dbCats []models.Category
	var dbDocs []models.Document
	db.Find(&dbCats)
	db.Find(&dbDocs)

	dbCatsMap := make(map[string]models.Category)
	dbDocsMap := make(map[string]models.Document)

	for _, cat := range dbCats {
		dbCatsMap[cat.Filepath] = cat
	}

	for _, doc := range dbDocs {
		dbDocsMap[doc.Filepath] = doc
	}

	// 本地文件元数据映射表
	localMetas := make(map[string]LocalMetaDatasCache)

	// 上一文件深度
	var catCount uint = 0

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

		if info.IsDir() && (info.Name() == "Ue") {
			return filepath.SkipDir
		}

		if info.Name() == "meta.json" {
			return nil
		}

		path = strings.ReplaceAll(path, root+"\\", "")
		Depth := strings.Count(path, "\\")
		parent := filepath.Dir(path)

		metaData := models.MetaData{
			Name:     info.Name(),
			Title:    info.Name(),
			Depth:    Depth,
			Icon:     "",
			Status:   1,
			Filepath: path,
		}

		if info.IsDir() {

			catCount++

			cat := models.Category{
				MetaData: metaData,
				ModTime:  info.ModTime(),
				Display:  true,
			}

			// 储存本地Meta数据
			if value, exists := localMetas[parent]; exists {
				value.Categorys = append(value.Categorys, cat)
				cat.Order = uint(len(value.Categorys))
				localMetas[parent] = value
			} else {
				cat.Order = 1
				localMetas[parent] = LocalMetaDatasCache{
					ParentFolder: parent,
					Categorys:    []models.Category{cat},
					Documents:    []models.Document{},
				}
			}

			// 储存数据库数据
			if value, exists := dbCatsMap[path]; exists {

				if value.ModTime.Equal(info.ModTime()) {
					// 如果数据库有 并且修改时间没变则跳过
					return nil
				} else {
					// 如果数据库有 并且修改时间变化, 则更新
					dbDatas.Cats.Updates = append(dbDatas.Cats.Updates, cat)
				}

				// 删除暂存数据
				delete(dbCatsMap, path)

			} else {
				// 数据库没有 则追加
				dbDatas.Cats.Creates = append(dbDatas.Cats.Creates, cat)
			}

		} else {

			doc := models.Document{
				MetaData:   metaData,
				ModTime:    info.ModTime(),
				Locale:     "",
				Content:    "",
				CategoryID: catCount,
			}

			if value, exists := localMetas[parent]; exists {
				value.Documents = append(value.Documents, doc)
				doc.Order = uint(len(value.Documents))
				localMetas[parent] = value
			} else {
				doc.Order = 1
				localMetas[parent] = LocalMetaDatasCache{
					ParentFolder: parent,
					Categorys:    []models.Category{},
					Documents:    []models.Document{doc},
				}
			}

			if value, exists := dbDocsMap[path]; exists {

				if value.ModTime.Equal(info.ModTime()) {
					return nil
				} else {
					dbDatas.Docs.Updates = append(dbDatas.Docs.Updates, doc)
				}
				delete(dbDocsMap, path)

			} else {
				dbDatas.Docs.Creates = append(dbDatas.Docs.Creates, doc)
			}

		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("读取数据用时", time.Since(start))

	start = time.Now()
	// WriteLocalMetaData(localMetas)
	fmt.Println("写入meta数据用时", time.Since(start))

	start = time.Now()
	// WriteContentToDocsData(&localDocs)
	fmt.Println("读取内容用时", time.Since(start))

	start = time.Now()

	for _, cat := range dbCatsMap {
		dbDatas.Cats.Deletes = append(dbDatas.Cats.Deletes, cat)
	}
	for _, doc := range dbDocsMap {
		dbDatas.Docs.Deletes = append(dbDatas.Docs.Deletes, doc)
	}

	collections := Collections{
		Creates: []interface{}{dbDatas.Cats.Creates, dbDatas.Docs.Creates},
		Updates: []interface{}{dbDatas.Cats.Updates, dbDatas.Docs.Updates},
		Deletes: []interface{}{dbDatas.Cats.Deletes, dbDatas.Docs.Deletes},
		Models:  []interface{}{models.Category{}, models.Document{}},
	}

	err = DBUpdate(db, collections)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	fmt.Println("保存数据库用时", time.Since(start))

	return nil
}
