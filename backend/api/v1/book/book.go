package book

import (
	"docsfly/internal/common"
	"docsfly/internal/global"
	"docsfly/models"
	"docsfly/pkg/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 获取当前书籍章节信息,没有章节则直接获取文档信息
func Book(c *gin.Context) {
	bookPath := c.Query("bookPath")
	locale := c.Query("locale")
	clientTime := time.Now()

	if bookPath == "" || locale == "" {
		common.Responser.Fail(c, http.StatusNotFound, clientTime, "参数不全")
	}

	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var chapters []models.Entry
	var docs []models.Entry
	var bookDatas []BookData

	// 查询分类章节和文章章节
	err := db.Scopes(common.BasicModel, common.FindChapter, common.HasPrefixUrlPath(bookPath+"/"+locale)).Order("is_dir DESC, depth ASC, `order` ASC").Find(&chapters).Error

	if err != nil {
		common.Responser.Fail(c, http.StatusInternalServerError, clientTime, "Database query error")
		return
	}

	if len(chapters) == 0 {
		common.Responser.Fail(c, http.StatusNotFound, clientTime, "Chapter does not exist or no matching documents found")
		return
	}

	err = db.Scopes(common.BasicModel, common.FindFile, common.HasPrefixUrlPath(bookPath+"/"+locale)).Where("depth > 3").Order("is_dir DESC, depth ASC, `order` ASC").Find(&docs).Error

	if err != nil {
		common.Responser.Fail(c, http.StatusInternalServerError, clientTime, "Database query error")
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
	common.Responser.Success(c, clientTime, bookDatas)
}

func BookStatistic(c *gin.Context) {

	bookPath := c.Query("bookPath")
	locale := c.Query("locale")

	clientTime := time.Now()
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
		common.Responser.Fail(c, http.StatusBadRequest, clientTime, "Book Don't Exist")
		return
	}

	db.Model(models.Visitor{}).Where("category = ?", category).Where("book = ?", book).Count(&bookReadCount)

	// 获取阅读量和书籍标题
	err := db.Scopes(common.BasicModel, common.MatchUrlPath(bookPath)).Find(&bookRef).Error
	if err != nil {
		common.Responser.Fail(c, http.StatusInternalServerError, clientTime, "Error fetching book statistics")
		return
	}

	// 获取章节数量
	db.Scopes(common.BasicModel, common.FindChapter, common.HasPrefixUrlPath(bookPath+"/"+locale)).Count(&chapterCount)

	// 获取书籍数量
	db.Scopes(common.BasicModel, common.FindFile, common.HasPrefixUrlPath(bookPath+"/"+locale)).Count(&documentCount)

	type bookStatistic struct {
		BookCover     string `json:"book_cover"`
		BookTitle     string `json:"book_title"`
		ReadCount     int64  `json:"read_count"`
		ChapterCount  int64  `json:"chapter_count"`
		DocumentCount int64  `json:"document_count"`
	}

	common.Responser.Success(c, clientTime,
		bookStatistic{
			BookCover:     bookRef.Icon,
			BookTitle:     bookRef.Title,
			ReadCount:     bookReadCount,
			ChapterCount:  chapterCount,
			DocumentCount: documentCount,
		})
}

func BookMeta(c *gin.Context) {
	bookPath := c.Query("bookPath")
	locale := c.Query("locale")

	clientTime := time.Now()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	filePath := getFilepathByURL(db, bookPath+"/"+locale)

	if filePath == "" {
		common.Responser.Fail(c, http.StatusInternalServerError, clientTime, "Failed Find Target Category")
		return
	}

	metaPath := global.AppConfig.DBConfig.Resource + "/" + filePath + "/" + global.AppConfig.DBConfig.MetaFile

	var data map[string]interface{}
	err := utils.ReadJson(metaPath, &data)

	if err != nil {
		common.Responser.Fail(c, http.StatusInternalServerError, clientTime, "Failed load data")
		return
	}
	common.Responser.Success(c, clientTime, data)
}

func UpdateBookMeta(c *gin.Context) {
	bookPath := c.Query("bookPath")
	locale := c.Query("locale")

	metas_data := c.Query("metas")

	clientTime := time.Now()

	ok, err := common.TokenVerifyMiddleware(c)
	if !ok {
		common.Responser.Fail(c, http.StatusUnauthorized, clientTime, err.Error())
		return
	}

	// 解析meta数据
	var metas models.MetaDatas

	if err := json.Unmarshal([]byte(metas_data), &metas); err != nil {
		common.Responser.Fail(c, http.StatusBadRequest, clientTime, err.Error())
		return
	}

	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	filepath := getFilepathByURL(db, bookPath+"/"+locale)

	// 保存meta.json
	metapath := global.AppConfig.DBConfig.Resource + "/" + filepath + "/" + global.AppConfig.DBConfig.MetaFile
	err = utils.WriteJson(metapath, metas)
	if err != nil {
		common.Responser.Fail(c, http.StatusInternalServerError, clientTime, "Failed Save Data")
		return
	}

	// 更新数据库
	for _, meta := range metas.Categories {
		db.Model(&models.Entry{}).Where("filepath = ?", filepath+"/"+meta.Name).Updates(meta)
	}

	for _, meta := range metas.Documents {
		db.Model(&models.Entry{}).Where("filepath = ?", filepath+"/"+meta.Name).Updates(meta)
	}
	common.Responser.Success(c, clientTime, "Data saved successfully")

}
