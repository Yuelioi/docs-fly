package handlers

// BookPage

import (
	"docsfly/database"
	"docsfly/global"
	"docsfly/models"
	"docsfly/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取当前书籍章节信息,没有章节则直接获取文档信息
func GetBook(c *gin.Context) {
	slug := c.Query("slug")
	locale := c.Query("locale")

	clientTime := time.Now()

	db, err := database.DbManager.Connect()

	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Failed load database")
		return
	}

	var cats, docs []models.Entry
	var bookDatas []BookData

	db.Model(models.Entry{}).Scopes(HasPrefixPath(slug+"/"+locale), FindChapter, FindFolder).Find(&cats)
	db.Model(models.Entry{}).Scopes(HasPrefixPath(slug+"/"+locale), FindChapter, FindFile).Where("depth = ?", 3).Find(&docs)

	// 1.分类数据
	for _, cat := range cats {

		var closeDoc models.Entry
		db.Model(models.Entry{}).Scopes(FindFile, HasPrefixPath(cat.Filepath)).
			Order("depth ASC, `order` ASC").
			Limit(1).Find(&closeDoc)

		if closeDoc.URLPath != "" {
			chapter := BookData{
				Url:      closeDoc.URLPath,
				IsDir:    true,
				MetaData: cat.MetaData,
			}
			bookDatas = append(bookDatas, chapter)
		}

	}

	// 2.文章数据
	for _, doc := range docs {
		if doc.URLPath != "" {
			bookDatas = append(bookDatas, BookData{
				Url:      doc.URLPath,
				IsDir:    false,
				MetaData: doc.MetaData,
			})
		}
	}

	sendSuccessResponse(c, clientTime, bookDatas)
}

func GetBookMeta(c *gin.Context) {
	slug := c.Query("slug")
	locale := c.Query("locale")

	clientTime := time.Now()

	db, err := database.DbManager.Connect()
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Failed load database")
		return
	}

	filepath := getFilepathByURLPath(db, slug+"/"+locale)

	if filepath == "" {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Failed Find Target Category")
		return
	}

	metapath := global.AppConfig.Resource + "/" + filepath + "/" + global.AppConfig.MetaFile

	var data map[string]interface{}
	err = utils.ReadJson(metapath, &data)

	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Failed load data")
		return
	}
	sendSuccessResponse(c, clientTime, data)
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

	filepath := getFilepathByURLPath(db, slug+"/"+locale)

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
