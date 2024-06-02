package database

import (
	"docsfly/models"
	"docsfly/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"

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

	go func() {
		for index, docsData := range *docsDatas {
			content, err := os.ReadFile(docsData.Filepath)
			if err != nil {
				log.Printf("%s", err)
				continue
			}
			(*docsDatas)[index].Content = string(content)
		}
	}()
}

func WriteMetaData(container FileContainer) {
	var output MetaOutput

	output.Documents = []models.MetaData{}
	output.Categorys = []models.MetaData{}

	for _, c := range container.Categorys {
		output.Categorys = append(output.Categorys, c.MetaData)
	}
	for _, d := range container.Documents {
		output.Documents = append(output.Documents, d.MetaData)
	}

	data, _ := json.MarshalIndent(output, "", "    ")
	outputPath := filepath.Join(container.Folder, "meta.json")

	os.WriteFile(outputPath, data, 0644)

}
