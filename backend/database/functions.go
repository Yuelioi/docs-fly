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

// 将完整数据转为本地要保存的数据
func convertMetaData(meta models.MetaData) models.LocalMetaData {
	return models.LocalMetaData{
		Name:   meta.Name,
		Title:  meta.Title,
		Order:  meta.Order,
		Icon:   meta.Icon,
		Status: meta.Status,
	}
}

func searchLocalMetaDatasCache(cache LocalMetaDatasCache, isCat bool, name string) *models.LocalMetaData {

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

func WriteLocalMetaData(metas map[string]LocalMetaDatasCache) {
	var wg sync.WaitGroup

	for _, meta := range metas {
		wg.Add(1)

		go func(meta LocalMetaDatasCache) {

			defer wg.Done()

			output := models.LocalMetaDatas{
				Categorys: meta.Categorys,
				Documents: meta.Documents,
			}

			data, _ := json.MarshalIndent(output, "", "    ")
			outputPath := filepath.Join(meta.ParentFolder, "meta.json")

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

	if info.Name() == "meta.json" {
		return ErrSkip
	}
	if info.Name() == "main.db" {
		return ErrSkip
	}

	return nil
}
