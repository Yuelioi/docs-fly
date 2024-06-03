package database

import (
	"docsfly/global"
	"docsfly/models"
	"docsfly/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"sync"

	"gorm.io/gorm"
)

// 初始化管理员账号
func CreateAdminAccount(db *gorm.DB) {
	hashedPassword, err := utils.HashPassword("admin")
	if err != nil {
		fmt.Println("初始化管理员数据失败")
		return
	}
	userData := models.User{
		Username: "admin",
		Password: hashedPassword,
	}
	db.Create(&userData)
}

// 把所有数据写入数据库
func WriteIntoDatabase(db *gorm.DB, datas ...interface{}) (err error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	batchSize := 100

	for _, data := range datas {
		value := reflect.ValueOf(data)
		length := value.Len()
		for i := 0; i < length; i += batchSize {
			endIndex := i + batchSize
			if endIndex > length {
				endIndex = length
			}
			batch := value.Slice(i, endIndex).Interface()
			if err = tx.Create(batch).Error; err != nil {
				return
			}
		}
	}

	return nil
}

// 读写Markdown内容 保存回文档数据
func WriteContentToDocsData(docsDatas *[]models.Document) {

	const maxGoroutines = 500
	guard := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup

	for index, docsData := range *docsDatas {
		wg.Add(1)
		guard <- struct{}{}
		go func(index int, docsData models.Document) {
			defer wg.Done()

			defer func() { <-guard }()

			docsPath := filepath.Join(global.AppConfig.Resource, docsData.MetaData.Filepath)
			content, err := os.ReadFile(docsPath)
			if err != nil {
				log.Printf("%s", err)
				return
			}
			(*docsDatas)[index].Content = string(content)
		}(index, docsData)
	}

	wg.Wait()
}

func WriteLocalMetaData(metas []LocalMetaCache) {
	var wg sync.WaitGroup

	for _, meta := range metas {
		wg.Add(1)

		go func(meta LocalMetaCache) {

			defer wg.Done()

			var output LocalMeta
			output.Documents = []models.MetaData{}
			output.Categorys = []models.MetaData{}

			for _, c := range meta.Categorys {
				output.Categorys = append(output.Categorys, c.MetaData)
			}
			for _, d := range meta.Documents {
				output.Documents = append(output.Documents, d.MetaData)
			}
			data, _ := json.MarshalIndent(output, "", "    ")
			outputPath := filepath.Join(meta.Folder, "meta.json")

			os.WriteFile(outputPath, data, 0644)
		}(meta)
	}

	wg.Wait()

}

func compareAndSync(db *gorm.DB, localCats []models.Category, localDocs []models.Document) error {
	// 从数据库中读取现有数据
	var dbCats []models.Category
	var dbDocs []models.Document
	db.Find(&dbCats)
	db.Find(&dbDocs)

	// 使用映射存储文件路径对应的数据
	dbCatsMap := make(map[string]models.Category)
	dbDocsMap := make(map[string]models.Document)

	for _, dbCat := range dbCats {
		dbCatsMap[dbCat.MetaData.Filepath] = dbCat
	}
	for _, dbDoc := range dbDocs {
		dbDocsMap[dbDoc.MetaData.Filepath] = dbDoc
	}

	// 存储需要进行的操作
	var catsToCreate []models.Category
	var catsToUpdate []models.Category
	var catsToDelete []models.Category
	var docsToCreate []models.Document
	var docsToUpdate []models.Document
	var docsToDelete []models.Document

	// 比较并同步类别数据
	for _, localCat := range localCats {
		if dbCat, exists := dbCatsMap[localCat.MetaData.Filepath]; exists {
			if localCat.ModTime != dbCat.ModTime {
				catsToUpdate = append(catsToUpdate, localCat)
			}
			delete(dbCatsMap, localCat.MetaData.Filepath)
		} else {
			catsToCreate = append(catsToCreate, localCat)
		}
	}

	// 剩下的 dbCatsMap 中的就是需要删除的
	for _, dbCat := range dbCatsMap {
		catsToDelete = append(catsToDelete, dbCat)
	}

	// 比较并同步文档数据
	for _, localDoc := range localDocs {
		if dbDoc, exists := dbDocsMap[localDoc.MetaData.Filepath]; exists {
			if localDoc.ModTime != dbDoc.ModTime {
				docsToUpdate = append(docsToUpdate, localDoc)
			}
			delete(dbDocsMap, localDoc.MetaData.Filepath)
		} else {
			docsToCreate = append(docsToCreate, localDoc)
		}
	}

	// 剩下的 dbDocsMap 中的就是需要删除的
	for _, dbDoc := range dbDocsMap {
		docsToDelete = append(docsToDelete, dbDoc)
	}

	// 执行数据库操作
	err := db.Transaction(func(tx *gorm.DB) error {
		if len(catsToCreate) > 0 {
			if err := tx.Create(&catsToCreate).Error; err != nil {
				return err
			}
		}
		if len(catsToUpdate) > 0 {
			for _, cat := range catsToUpdate {
				if err := tx.Model(&models.Category{}).Where("id = ?", cat.ID).Updates(cat).Error; err != nil {
					return err
				}
			}
		}
		if len(catsToDelete) > 0 {
			for _, cat := range catsToDelete {
				if err := tx.Delete(&cat).Error; err != nil {
					return err
				}
			}
		}
		if len(docsToCreate) > 0 {
			if err := tx.Create(&docsToCreate).Error; err != nil {
				return err
			}
		}
		if len(docsToUpdate) > 0 {
			for _, doc := range docsToUpdate {
				if err := tx.Model(&models.Document{}).Where("id = ?", doc.ID).Updates(doc).Error; err != nil {
					return err
				}
			}
		}
		if len(docsToDelete) > 0 {
			for _, doc := range docsToDelete {
				if err := tx.Delete(&doc).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to sync database: %w", err)
	}

	fmt.Println("数据库与本地文件同步完成")
	return nil
}
