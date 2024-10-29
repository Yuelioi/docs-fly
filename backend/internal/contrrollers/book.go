package controllers

import (
	"docsfly/internal/common"
	"docsfly/internal/config"
	"docsfly/internal/models"
	"docsfly/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
}

type BookData struct {
	Url      string          `json:"url"`
	IsDir    bool            `json:"is_dir"`
	MetaData models.MetaData `json:"metadata"`
}

type PageData struct {
	TotalCount int64         `json:"total_count"` // 总记录数
	Page       int           `json:"page"`        // 当前页码 从1开始
	PageSize   int           `json:"page_size"`   // 每页记录数
	Pages      []interface{} `json:"pages"`
}

func (BookController) Register(e *gin.Engine) {
	book := e.Group("/" + config.Instance.App.ApiVersion + "/book")

	book.GET("/", Book)
	book.GET("/statistic", BookStatistic)
	book.GET("meta", BookMeta)
	book.PUT("meta", UpdateBookMeta)
}

// 获取当前书籍章节信息,没有章节则直接获取文档信息
func Book(c *gin.Context) {
	bookPath := c.Query("bookPath")
	locale := c.Query("locale")

	if bookPath == "" || locale == "" {
		ReturnFailResponse(c, 400, "参数不全")
		return
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
		ReturnFailResponse(c, 400, "Database query error")
		return
	}

	if len(chapters) == 0 {
		ReturnFailResponse(c, 400, "Chapter does not exist or no matching documents found")
		return
	}

	err = db.Scopes(common.BasicModel, common.FindFile, common.HasPrefixUrlPath(bookPath+"/"+locale)).Where("depth > 3").Order("is_dir DESC, depth ASC, `order` ASC").Find(&docs).Error

	if err != nil {
		ReturnFailResponse(c, 400, "Database query error")
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

	ReturnSuccessResponse(c, bookDatas)
}

func BookStatistic(c *gin.Context) {

	bookPath := c.Query("bookPath")
	locale := c.Query("locale")

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
		ReturnFailResponse(c, 400, "Book Don't Exist")
		return
	}

	db.Model(models.Visitor{}).Where("category = ?", category).Where("book = ?", book).Count(&bookReadCount)

	// 获取阅读量和书籍标题
	err := db.Scopes(common.BasicModel, common.MatchUrlPath(bookPath)).Find(&bookRef).Error
	if err != nil {
		ReturnFailResponse(c, 400, "Error fetching book statistics")
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

	ReturnSuccessResponse(c,
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

	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	filePath := getFilepathByURL(db, bookPath+"/"+locale)

	if filePath == "" {
		ReturnFailResponse(c, 400, "Failed Find Target Category")
		return
	}

	metaPath := config.Instance.Database.Resource + "/" + filePath + "/" + config.Instance.Database.MetaFile

	var data map[string]interface{}
	err := utils.ReadJson(metaPath, &data)

	if err != nil {
		ReturnFailResponse(c, 400, "Failed load data")
		return
	}
	ReturnSuccessResponse(c, data)
}

func UpdateBookMeta(c *gin.Context) {
	bookPath := c.Query("bookPath")
	locale := c.Query("locale")

	metas_data := c.Query("metas")

	ok, err := common.TokenVerifyMiddleware(c)
	if !ok {
		ReturnFailResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// 解析meta数据
	var metas models.MetaDatas

	if err := json.Unmarshal([]byte(metas_data), &metas); err != nil {
		ReturnFailResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	filepath := getFilepathByURL(db, bookPath+"/"+locale)

	// 保存meta.json
	metapath := config.Instance.Database.Resource + "/" + filepath + "/" + config.Instance.Database.MetaFile
	err = utils.WriteJson(metapath, metas)
	if err != nil {
		ReturnFailResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 更新数据库
	for _, meta := range metas.Categories {
		db.Model(&models.Entry{}).Where("filepath = ?", filepath+"/"+meta.Name).Updates(meta)
	}

	for _, meta := range metas.Documents {
		db.Model(&models.Entry{}).Where("filepath = ?", filepath+"/"+meta.Name).Updates(meta)
	}
	ReturnSuccessResponse(c, "Data saved successfully")

}
