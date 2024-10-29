package controllers

import (
	"docsfly/internal/common"
	"docsfly/internal/config"
	"docsfly/internal/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HomeController struct {
}

func (HomeController) Register(engine *gin.Engine) {
	engine.GET("/"+config.Instance.App.ApiVersion+"/nav", GetNav)
	engine.GET("/"+config.Instance.App.ApiVersion+"/statistic/home", GetStatisticHome)
	engine.GET("/"+config.Instance.App.ApiVersion+"/query", Query)
}

type SearchResult struct {
	Result []SearchData `json:"result"`
}

// 主页搜索显示数据
type SearchData struct {
	Index         int    `json:"index"` // 从1开始!
	Url           string `json:"url"`
	Locale        string `json:"locale"`
	CategoryTitle string `json:"category_title"`
	BookTitle     string `json:"book_title"`
	DocumentTitle string `json:"document_title"`
	Content       string `json:"content"`
}

// 主页统计信息
type HomeStatistic struct {
	BookCount              int64 `json:"book_count"`
	DocumentCount          int64 `json:"document_count"`
	HistoricalVisitorCount int64 `json:"historical_visitor_count"`
	TodayVisitorCount      int64 `json:"today_visitor_count"`
}

type Nav struct {
	MetaData models.MetaData   `json:"metadata"`
	Children []models.MetaData `json:"children"`
}

// 获取顶部导航栏信息
func GetNav(c *gin.Context) {

	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var cats []models.Entry
	var books []models.Entry
	var navs []Nav

	if err := db.Scopes(common.BasicModel, common.FindCategory, common.FindFolder).Find(&cats).Error; err != nil {
		ReturnFailResponse(c, http.StatusInternalServerError, "Failed to retrieve categories")
		return
	}
	if err := db.Scopes(common.BasicModel, common.FindBook, common.FindFolder).Find(&books).Error; err != nil {
		ReturnFailResponse(c, http.StatusInternalServerError, "Failed to retrieve books")
		return
	}

	for _, cat := range cats {
		nav := Nav{}
		nav.MetaData = cat.MetaData
		for _, book := range books {
			if strings.HasPrefix(book.Filepath, cat.Filepath) {
				nav.Children = append(nav.Children, book.MetaData)
			}
		}
		navs = append(navs, nav)
	}
	ReturnSuccessResponse(c, navs)
}

func GetStatisticHome(c *gin.Context) {

	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var (
		BooksCount             int64
		DocumentsCount         int64
		HistoricalVisitorCount int64
		TodayVisitorCount      int64
	)

	db.Scopes(common.BasicModel, common.FindBook, common.FindFolder).Count(&BooksCount)
	db.Scopes(common.BasicModel, common.FindFile).Count(&DocumentsCount)

	db.Model(models.Visitor{}).Count(&HistoricalVisitorCount)

	// today := time.Now().Format("2006-01-02") 不能用(DATE(time))
	today := time.Now().Truncate(24 * time.Hour)
	db.Model(models.Visitor{}).Where("time > ?", today).Count(&TodayVisitorCount)

	statistic := HomeStatistic{
		BookCount:              BooksCount,
		DocumentCount:          DocumentsCount,
		HistoricalVisitorCount: HistoricalVisitorCount,
		TodayVisitorCount:      TodayVisitorCount,
	}
	ReturnSuccessResponse(c, statistic)

}

func Query(c *gin.Context) {

	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	bookPath := c.Query("bookPath")
	keyword := c.Query("keyword")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	// 解析分页参数
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 50 {
		pageSize = 20
	}

	// 计算 offset
	offset := (page - 1) * pageSize

	// 显示关键词周围文字数量
	splitContentSize := 30

	var documents []models.Entry
	var totalCount int64
	var results []SearchData

	// 根据查询条件获取结果
	query := db.Scopes(common.BasicModel)
	if bookPath != "" {
		query = query.Scopes(common.HasPrefixUrlPath(bookPath))
	}

	// 获取总查询结果
	query = query.Where("content LIKE ?", "%"+keyword+"%")
	query.Scopes(common.BasicModel).Count(&totalCount)

	if totalCount == 0 {
		ReturnFailResponse(c, http.StatusNotFound, "No documents found")
		return
	}

	// 获取分页结果
	query = query.Offset(offset).Limit(pageSize)
	query.Find(&documents)

	if int64(offset) > totalCount {
		ReturnFailResponse(c, http.StatusBadRequest, "Query result exceeds maximum value")
		return
	}

	for i, document := range documents {

		plaintext, err := extractPlainText(document.Content)
		if err != nil {
			continue
		}

		runeSlice := []rune(*plaintext)
		keywordIndex := indexOfKeywordInRuneSlice(runeSlice, keyword)

		if keywordIndex == -1 {
			continue
		}

		start := keywordIndex - splitContentSize
		if start < 0 {
			start = 0
		}

		// 确定截取的结束位置
		end := keywordIndex + splitContentSize
		if end > len(runeSlice) {
			end = len(runeSlice)
		}

		runeText := runeSlice[start:end]

		nearbyText := string(runeText)

		cat, book, locale, ok := filepath2Params(document.Filepath)

		if ok {

			var catTitle string
			var bookTitle string
			db.Scopes(common.BasicModel, common.MatchUrlPath(cat)).Select("title").Scan(&catTitle)
			db.Scopes(common.BasicModel, common.MatchUrlPath(cat+"/"+book)).Select("title").Scan(&bookTitle)

			dsData := SearchData{
				Index:         offset + i + 1,
				Url:           document.URL,
				CategoryTitle: catTitle,
				BookTitle:     bookTitle,
				Locale:        locale,
				DocumentTitle: strings.ReplaceAll(document.Title, ".md", ""),
				Content:       nearbyText,
			}
			results = append(results, dsData)
		}

	}

	if len(results) == 0 {
		ReturnFailResponse(c, http.StatusNotFound, "Keyword Match Error")
		return
	}

	ReturnSuccessResponse(c, results)
}
