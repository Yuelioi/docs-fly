package handlers

// BookPage

import (
	"docsfly/database"
	"docsfly/models"
	"docsfly/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取当前书籍章节信息,没有章节则直接获取文档信息
func GetBook(c *gin.Context) {
	slug := c.Query("slug")
	locale := c.Query("locale")

	db, err := database.DbManager.Connect()

	var cats []models.Category
	db.Model(models.Category{}).Preload("Documents").Where("filepath like ?", slug+"/"+locale+"%").Where("depth = ?", 3).Find(&cats)

	var docs []models.Document
	db.Model(models.Document{}).Where("filepath like ?", slug+"/"+locale+"%").Where("depth = ?", 3).Find(&docs)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	var bookDatas []models.BookData

	// 1.分类数据
	for _, cat := range cats {

		var minOrderDocument models.Document

		minOrderDocument.Order = 9999

		// TODO 如果是UE 也许没有第一篇文章
		for _, doc := range cat.Documents {
			if doc.Order < minOrderDocument.Order {
				minOrderDocument = doc
			}
		}

		chapter := models.BookData{
			Url:         minOrderDocument.Filepath,
			ChapterType: "category",
			MetaData:    cat.MetaData,
		}
		bookDatas = append(bookDatas, chapter)
	}

	// 2.文章数据
	for _, doc := range docs {

		chapter := models.BookData{
			Url:         doc.Filepath,
			ChapterType: "document",
			MetaData:    doc.MetaData,
		}
		bookDatas = append(bookDatas, chapter)
	}

	c.JSON(http.StatusOK, bookDatas)
}

func GetBookMeta(c *gin.Context) {
	category := c.Query("category")

	filepath := category + "/" + "meta.json"

	var data interface{}
	err := utils.ReadJson(filepath, data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load data"})
		return
	}

	c.JSON(http.StatusOK, &data)
}

func SaveBookMeta(c *gin.Context) {
	category := c.Query("category")

	metas_data := c.Query("metas")

	ok, err := TokenVerifyMiddleware(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	// 解析meta数据
	var metas []models.MetaData

	if err := json.Unmarshal([]byte(metas_data), &metas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	db, err := database.DbManager.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Parse Data"})
		return
	}

	var catInfo models.Category

	db.Model(&catInfo).Where("identity = ?", category).First(&catInfo)

	// 保存meta.json
	filepath := catInfo.Filepath + "/" + "meta.json"
	err = utils.WriteJson(filepath, metas)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed Save Data"})
		return
	}

	// 更新数据库

	// 使用一个map来快速查找meta数据
	metaMap := make(map[string]models.MetaData)
	for _, meta := range metas {
		metaMap[meta.Name] = meta
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

	var catInfo models.Category

	db.Model(&catInfo).Where("identity = ?", category).First(&catInfo)

	// 保存meta.json
	filepath := catInfo.Filepath + "/" + "meta.json"

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
