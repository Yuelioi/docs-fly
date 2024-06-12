package handlers

// BookPage

import (
	"docsfly/database"
	"docsfly/global"
	"docsfly/models"
	"docsfly/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取当前书籍章节信息,没有章节则直接获取文档信息
func GetBook(c *gin.Context) {
	slug := c.Query("slug")
	locale := c.Query("locale")

	db, err := database.DbManager.Connect()

	var cats []models.Entry
	db.Model(models.Entry{}).Where("webpath like ?", slug+"/"+locale+"%").Where("file_type", 1).Where("depth = ?", 3).Find(&cats)

	var docs []models.Entry
	db.Model(models.Entry{}).Where("webpath like ?", slug+"/"+locale+"%").Where("file_type", 0).Where("depth = ?", 3).Find(&docs)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	var bookDatas []models.BookData

	// 1.分类数据
	for _, cat := range cats {

		var minOrderDocument models.Entry

		minOrderDocument.Order = 9999

		// // TODO 如果是UE 也许没有第一篇文章
		// for _, doc := range cat.Documents {
		// 	if doc.Order < minOrderDocument.Order {
		// 		minOrderDocument = doc
		// 	}
		// }

		chapter := models.BookData{
			Url:         minOrderDocument.WebPath,
			ChapterType: "category",
			MetaData:    cat.MetaData,
		}
		// TODO 这里找不到应该log 个 warning
		if chapter.Url != "" {
			bookDatas = append(bookDatas, chapter)
		} else {
			fmt.Printf("chapter: %v\n", chapter)
		}
	}

	// 2.文章数据
	for _, doc := range docs {

		chapter := models.BookData{
			Url:         doc.WebPath,
			ChapterType: "document",
			MetaData:    doc.MetaData,
		}
		if chapter.Url != "" {
			bookDatas = append(bookDatas, chapter)
		} else {
			fmt.Printf("chapter: %v\n", chapter)
		}
	}

	c.JSON(http.StatusOK, bookDatas)
}

func GetBookMeta(c *gin.Context) {
	slug := c.Query("slug")
	locale := c.Query("locale")

	db, err := database.DbManager.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	filepath := GetFilepathByWebpath(db, "category", slug+"/"+locale)

	if filepath == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Find Target Category"})
		return
	}

	metapath := global.AppConfig.Resource + "/" + filepath + "/" + global.AppConfig.MetaFile

	var data map[string]interface{}
	err = utils.ReadJson(metapath, &data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load data"})
		return
	}

	c.JSON(http.StatusOK, &data)
}

func SaveBookMeta(c *gin.Context) {
	slug := c.Query("slug")
	locale := c.Query("locale")

	metas_data := c.Query("metas")

	ok, err := TokenVerifyMiddleware(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	// 解析meta数据
	var metas models.MetaDatas

	if err := json.Unmarshal([]byte(metas_data), &metas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.DbManager.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Parse Data"})
		return
	}

	filepath := GetFilepathByWebpath(db, "category", slug+"/"+locale)

	// 保存meta.json
	metapath := global.AppConfig.Resource + "/" + filepath + "/" + global.AppConfig.MetaFile
	err = utils.WriteJson(metapath, metas)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Save Data"})
		return
	}

	// 更新数据库
	for _, meta := range metas.Categorys {
		db.Model(&models.Entry{}).Where("filepath = ?", filepath+"/"+meta.Name).Updates(meta)
	}

	for _, meta := range metas.Documents {
		db.Model(&models.Entry{}).Where("filepath = ?", filepath+"/"+meta.Name).Updates(meta)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data saved successfully"})
}

func UpdateBookMeta(c *gin.Context) {
	category := c.Query("category")

	db, err := database.DbManager.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Parse Data"})
		return
	}

	var catInfo models.Entry

	db.Model(&catInfo).Where("identity = ?", category).First(&catInfo)

	// 保存meta.json
	filepath := catInfo.Filepath + "/" + global.AppConfig.MetaFile

	var metas interface{}

	err = utils.ReadJson(filepath, metas)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load data"})
		return
	}

	// 更新数据库

	metaMap := make(map[string]models.MetaData)

	metadataSlice, ok := metas.([]models.MetaData)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error converting to []Metadata"})
		return
	}
	for _, meta := range metadataSlice {
		metaMap[meta.Name] = meta
	}

	c.JSON(http.StatusOK, &metas)
}
