package handlers

// BookPage

import (
	"docsfly/global"
	"docsfly/models"
	"docsfly/utils"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetStatisticBook(c *gin.Context) {

	bookPath := c.Query("bookPath")
	locale := c.Query("locale")

	clientTime := currentTime()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var bookRef models.Entry
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
	err := db.Scopes(BasicModel, MatchUrlPath(bookPath)).Find(&bookRef).Error
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Error fetching book statistics")
		return
	}

	// 获取章节数量
	db.Scopes(BasicModel, FindChapter, HasPrefixUrlPath(bookPath+"/"+locale)).Count(&chapterCount)

	// 获取书籍数量
	db.Scopes(BasicModel, FindFile, HasPrefixUrlPath(bookPath+"/"+locale)).Count(&documentCount)

	type bookStatistic struct {
		BookCover     string `json:"book_cover"`
		BookTitle     string `json:"book_title"`
		ReadCount     int64  `json:"read_count"`
		ChapterCount  int64  `json:"chapter_count"`
		DocumentCount int64  `json:"document_count"`
	}

	sendSuccessResponse(c, clientTime,
		bookStatistic{
			BookCover:     bookRef.Icon,
			BookTitle:     bookRef.Title,
			ReadCount:     bookReadCount,
			ChapterCount:  chapterCount,
			DocumentCount: documentCount,
		})
}

func findClosestDoc(chapter models.Entry, docs []models.Entry) models.Entry {
	var closestDoc models.Entry
	minOrder := int(^uint(0) >> 1) // 初始化为最大值

	for _, doc := range docs {
		if strings.HasPrefix(doc.URL, chapter.URL) && int(doc.Order) < minOrder {
			closestDoc = doc
			minOrder = int(doc.Order)
		}
	}

	return closestDoc
}

// 获取当前书籍章节信息,没有章节则直接获取文档信息
func GetBook(c *gin.Context) {
	bookPath := c.Query("bookPath")
	locale := c.Query("locale")
	clientTime := currentTime()

	if bookPath == "" || locale == "" {
		sendErrorResponse(c, http.StatusNotFound, clientTime, "参数不全")
	}

	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var chapters []models.Entry
	var docs []models.Entry
	var bookDatas []BookData

	ok, cachedData := getCache(bookPath + locale)
	if ok {
		sendSuccessResponse(c, clientTime, cachedData)
		return
	}

	// 查询分类章节和文章章节
	err := db.Scopes(BasicModel, FindChapter, HasPrefixUrlPath(bookPath+"/"+locale)).Order("is_dir DESC, depth ASC, `order` ASC").Find(&chapters).Error

	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Database query error")
		return
	}

	if len(chapters) == 0 {
		sendErrorResponse(c, http.StatusNotFound, clientTime, "Chapter does not exist or no matching documents found")
		return
	}

	err = db.Scopes(BasicModel, FindFile, HasPrefixUrlPath(bookPath+"/"+locale)).Where("depth > 3").Order("is_dir DESC, depth ASC, `order` ASC").Find(&docs).Error

	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Database query error")
		return
	}

	for _, chapter := range chapters {
		if chapter.IsDir {

			closeDoc := findClosestDoc(chapter, docs)

			if closeDoc.URL != "" {
				chapter := BookData{
					Url:      closeDoc.URL,
					IsDir:    true,
					MetaData: chapter.MetaData,
				}
				bookDatas = append(bookDatas, chapter)
			}
		} else {
			bookDatas = append(bookDatas, BookData{
				Url:      chapter.URL,
				IsDir:    false,
				MetaData: chapter.MetaData,
			})
		}
	}
	saveCache(bookPath+locale, bookDatas)
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

	filePath := getFilepathByURL(db, bookPath+"/"+locale)

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

func UpdateBookMeta(c *gin.Context) {
	bookPath := c.Query("bookPath")
	locale := c.Query("locale")

	metas_data := c.Query("metas")

	clientTime := currentTime()

	ok, err := TokenVerifyMiddleware(c)
	if !ok {
		sendErrorResponse(c, http.StatusUnauthorized, clientTime, err.Error())
		return
	}

	// 解析meta数据
	var metas models.MetaDatas

	if err := json.Unmarshal([]byte(metas_data), &metas); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, clientTime, err.Error())
		return
	}

	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	filepath := getFilepathByURL(db, bookPath+"/"+locale)

	// 保存meta.json
	metapath := global.AppConfig.Resource + "/" + filepath + "/" + global.AppConfig.MetaFile
	err = utils.WriteJson(metapath, metas)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Failed Save Data")
		return
	}

	// 更新数据库
	for _, meta := range metas.Categorys {
		db.Model(&models.Entry{}).Where("filepath = ?", filepath+"/"+meta.Name).Updates(meta)
	}

	for _, meta := range metas.Documents {
		db.Model(&models.Entry{}).Where("filepath = ?", filepath+"/"+meta.Name).Updates(meta)
	}
	sendSuccessResponse(c, clientTime, "Data saved successfully")

}
