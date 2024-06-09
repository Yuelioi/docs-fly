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
	"strings"
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

	batchSize := 128

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
				return err
			}
		}
	}

	return nil
}

// 读写Markdown内容 保存回文档数据
func WriteContentToDocsData(datas ...*[]models.Document) {

	const maxGoroutines = 500
	guard := make(chan struct{}, maxGoroutines)

	var wg sync.WaitGroup

	for _, docsDatas := range datas {
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
	}

	wg.Wait()
}

func searchMetaDatasCache(cache MetaDatasCache, isCat bool, name string) *models.MetaData {
	if isCat {
		for _, meta := range cache.Categorys {
			if meta.Name == name {
				return &meta
			}
		}
	} else {
		for _, meta := range cache.Documents {
			if meta.Name == name {
				return &meta
			}
		}
	}
	return nil
}

func searchDBCatMetaDatas(cache []models.Category, path string) *models.MetaData {
	for _, item := range cache {
		if item.Filepath == path {
			return &item.MetaData
		}
	}
	return nil
}

func searchDBDocMetaDatas(cache []models.Document, path string) *models.MetaData {
	for _, item := range cache {
		if item.Filepath == path {
			return &item.MetaData
		}
	}
	return nil
}

func compare(localMeta *models.MetaData, dbMeta *models.MetaData) bool {
	if localMeta == nil || dbMeta == nil {
		return false
	}
	return *localMeta == *dbMeta
}

// 写入本地meta.json
// @param: rebuild false:只写入修改的 true:全部重写
func WriteMetaData(
	metas map[string]MetaDatasCache,
	rebuild bool,
) {

	update_metas := make([]MetaDatasCache, 0)

	for _, meta := range metas {
		if meta.NeedWrite || rebuild {
			update_metas = append(update_metas, meta)
		}

	}

	var wg sync.WaitGroup

	for _, meta := range update_metas {
		wg.Add(1)

		go func(meta MetaDatasCache) {

			defer wg.Done()

			output := models.MetaDatas{
				Categorys: meta.Categorys,
				Documents: meta.Documents,
			}

			data, _ := json.MarshalIndent(output, "", "    ")
			outputPath := filepath.Join(meta.ParentFolder, global.AppConfig.MetaFile)

			os.WriteFile(outputPath, data, 0644)

		}(meta)
	}

	wg.Wait()

}

func WalkSkip(root string, info os.FileInfo, path string, err error) error {

	if path == root {
		return ErrSkip
	}

	if info.IsDir() && strings.HasPrefix(info.Name(), ".") {
		return filepath.SkipDir
	}

	if info.IsDir() && strings.HasPrefix(info.Name(), "_") {
		return filepath.SkipDir
	}

	if info.IsDir() && strings.ToLower(info.Name()) == "ue" {
		return filepath.SkipDir
	}

	if !info.IsDir() && strings.HasPrefix(info.Name(), "_") {
		return ErrSkip
	}

	if info.Name() == global.AppConfig.MetaFile {
		return ErrSkip
	}
	if info.Name() == "main.db" {
		return ErrSkip
	}

	return nil
}
