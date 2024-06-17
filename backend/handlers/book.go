package handlers

// BookPage

import (
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
	locale := c.Query("locale")

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
	err := db.Scopes(BasicModel, MatchUrlPath(bookPath)).
		Select("title").Scan(&bookTitle).Error
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Error fetching book statistics")
		return
	}

	// 获取章节数量
	db.Scopes(BasicModel, FindChapter, HasPrefixUrlPath(bookPath+"/"+locale)).Count(&chapterCount)

	// 获取书籍数量
	db.Scopes(BasicModel, FindFile, HasPrefixUrlPath(bookPath+"/"+locale)).Count(&documentCount)

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
	err := db.Scopes(BasicModel, HasPrefixUrlPath(bookPath+"/"+locale), FindChapter).
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
			db.Scopes(BasicModel, FindFile, HasPrefixUrlPath(entry.Filepath)).
				Order("depth ASC, `order` ASC").
				Limit(1).Find(&closeDoc)

			if closeDoc.URL != "" {
				chapter := BookData{
					Url:      closeDoc.URL,
					IsDir:    true,
					MetaData: entry.MetaData,
				}
				bookDatas = append(bookDatas, chapter)
			}
		} else {
			bookDatas = append(bookDatas, BookData{
				Url:      entry.URL,
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
