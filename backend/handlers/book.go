package handlers

// BookPage

import (
	"docsfly/database"
	"docsfly/global"
	"docsfly/models"
	"docsfly/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetStatisticBook(c *gin.Context) {

	bookPath := c.Query("bookPath")

	clientTime := currentTime()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var bookTitle string
	var bookReadCount int64
	var chapterCount int64
	var documentCount int64

	category, book := getCategoryAndBookByUrl(bookPath)

	if category == "" || book == "" {
		sendErrorResponse(c, http.StatusBadRequest, clientTime, "Book Don't Exist")
		return
	}

	db.Model(models.Visitor{}).Where("category = ?", category).Where("book = ?", book).Count(&bookReadCount)

	// 获取阅读量和书籍标题
	err := db.Scopes(BasicModel, MatchPath(bookPath)).
		Select("title").Scan(&bookTitle).Error
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Error fetching book statistics")
		return
	}

	// 获取章节数量
	db.Scopes(BasicModel, FindChapter, HasPrefixPath(bookPath)).Count(&chapterCount)

	// 获取书籍数量
	db.Scopes(BasicModel, FindFile, HasPrefixPath(bookPath)).Count(&documentCount)

	type bookStatistic struct {
		BookTitle     string `json:"book_title"`
		ReadCount     int64  `json:"read_count"`
		ChapterCount  int64  `json:"chapter_count"`
		DocumentCount int64  `json:"document_count"`
	}

	sendSuccessResponse(c, clientTime,
		bookStatistic{
			BookTitle:     bookTitle,
			ReadCount:     bookReadCount,
			ChapterCount:  chapterCount,
			DocumentCount: documentCount,
		})
}

// 获取当前书籍章节信息,没有章节则直接获取文档信息
func GetBook(c *gin.Context) {
	bookPath := c.Query("bookPath")
	locale := c.Query("locale")

	clientTime := currentTime()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var entries []models.Entry
	var bookDatas []BookData

	// 查询分类章节和文章章节
	err := db.Scopes(BasicModel, HasPrefixPath(bookPath+"/"+locale), FindChapter).
		Order("is_dir DESC, depth ASC, `order` ASC").
		Find(&entries).Error

	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Database query error")
		return
	}

	if len(entries) == 0 {
		sendErrorResponse(c, http.StatusNotFound, clientTime, "Chapter does not exist or no matching documents found")
		return
	}

	for _, entry := range entries {
		if entry.IsDir {
			var closeDoc models.Entry
			db.Scopes(BasicModel, FindFile, HasPrefixPath(entry.Filepath)).
				Order("depth ASC, `order` ASC").
				Limit(1).Find(&closeDoc)

			if closeDoc.URLPath != "" {
				chapter := BookData{
					Url:      closeDoc.URLPath,
					IsDir:    true,
					MetaData: entry.MetaData,
				}
				bookDatas = append(bookDatas, chapter)
			}
		} else {
			bookDatas = append(bookDatas, BookData{
				Url:      entry.URLPath,
				IsDir:    false,
				MetaData: entry.MetaData,
			})
		}
	}

	sendSuccessResponse(c, clientTime, bookDatas)
}

func GetBookMeta(c *gin.Context) {
	bookPath := c.Query("bookPath")
	locale := c.Query("locale")

	clientTime := currentTime()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	filePath := getFilepathByURLPath(db, bookPath+"/"+locale)

	if filePath == "" {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Failed Find Target Category")
		return
	}

	metaPath := global.AppConfig.Resource + "/" + filePath + "/" + global.AppConfig.MetaFile

	var data map[string]interface{}
	err := utils.ReadJson(metaPath, &data)

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
