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

// 保存每个单元的数据
func SaveUnitMetaData(data any, srcPath string) {

	result, err := json.Marshal(&data)

	if err != nil {
		fmt.Println("SaveUnitMetaData", err)
		return
	}

	output_dir := filepath.Dir(srcPath)
	output_file := filepath.Join(output_dir, "meta.json")

	err = os.WriteFile(output_file, result, 0666)

	if err != nil {
		fmt.Printf("SaveUnitMetaData WriteFile err: %v\n", err)
		return
	}
}

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

// 基于暂存数据进行查找当前类目元数据
func CreateMetaByCurrentMetas(currentMetas *[]models.MetaData, path string, info os.FileInfo, order uint) (*[]models.MetaData, models.MetaData) {
	var meta models.MetaData
	if currentMetas == nil {
		cacheMetas, err := utils.ReadMetas(filepath.Join(filepath.Dir(path), "meta.json"), info)
		if err != nil {
			meta = *utils.CreateMeta(info, order)
		} else {
			meta = *utils.SearchMeta(cacheMetas, info, order)
		}
		utils.UpdateMeta(&meta, info.Name(), utils.PureFileName(info.Name()), order, false)
		return cacheMetas, meta
	} else {
		meta = *utils.SearchMeta(currentMetas, info, order)
		utils.UpdateMeta(&meta, info.Name(), utils.PureFileName(info.Name()), order, false)
		return currentMetas, meta
	}
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

// 保存元数据到本地 Mode 1专属
func WriteMetadataToLocal(
	currentCatBooks []models.Book,
	cateLocalMeta []models.MetaDataLocal,
	currentBookChapters []models.Chapter,
	bookLocalMeta []models.MetaDataLocal,
	summaryMeta []models.MetaData,
	catDatas []models.Category,
) {
	if len(currentCatBooks) > 0 {
		metas := make([]models.MetaData, 0)
		for _, data := range currentCatBooks {
			metas = append(metas, data.MetaData)
		}

		cateLocalMeta = append(cateLocalMeta, models.MetaDataLocal{
			MetaDatas: metas,
			Filepath:  currentCatBooks[0].Filepath,
		})

	}

	if len(currentBookChapters) > 0 {
		metas := make([]models.MetaData, 0)
		for _, data := range currentBookChapters {
			metas = append(metas, data.MetaData)
		}
		bookLocalMeta = append(bookLocalMeta, models.MetaDataLocal{
			MetaDatas: metas,
			Filepath:  currentBookChapters[0].Filepath,
		})
	}

	SaveUnitMetaData(summaryMeta, catDatas[0].Filepath)

	for _, data := range cateLocalMeta {
		SaveUnitMetaData(data.MetaDatas, data.Filepath)

	}
	for _, data := range bookLocalMeta {
		SaveUnitMetaData(data.MetaDatas, data.Filepath)

	}
}

func WriteMetaToDocument(docsDatas []models.Document) {

}
