package handlers

// HomePage

import (
	"docsfly/models"
	"docsfly/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 获取顶部导航栏信息
func GetNav(c *gin.Context) {
	clientTime := currentTime()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	db := dbContext.(*gorm.DB)

	var cats []models.Entry
	var books []models.Entry
	var navs []Nav

	if err := db.Scopes(BasicModel, FindCategory, FindFolder).Find(&cats).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Failed to retrieve categories")
		return
	}
	if err := db.Scopes(BasicModel, FindBook, FindFolder).Find(&books).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Failed to retrieve books")
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
	sendSuccessResponse(c, clientTime, navs)
}

func GetStatisticHome(c *gin.Context) {
	clientTime := currentTime()
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

	db.Scopes(BasicModel, FindBook, FindFolder).Count(&BooksCount)
	db.Scopes(BasicModel, FindFile).Count(&DocumentsCount)

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
	sendSuccessResponse(c, clientTime, statistic)

}

func Query(c *gin.Context) {
	clientTime := currentTime()

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
	query := db.Scopes(BasicModel)
	if bookPath != "" {
		query = query.Scopes(HasPrefixUrlPath(bookPath))
	}

	// 获取总查询结果
	query = query.Where("content LIKE ?", "%"+keyword+"%")
	query.Scopes(BasicModel).Count(&totalCount)

	if totalCount == 0 {
		sendErrorResponsePageData(c, http.StatusNotFound, clientTime, totalCount, page, pageSize, "No documents found")
		return
	}

	// 获取分页结果
	query = query.Offset(offset).Limit(pageSize)
	query.Find(&documents)

	if int64(offset) > totalCount {
		sendErrorResponsePageData(c, http.StatusBadRequest, time.Now(), totalCount, page, pageSize, "Query result exceeds maximum value")
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

		cat, book, locale, ok := utils.Filepath2Params(document.Filepath)

		if ok {

			var catTitle string
			var bookTitle string
			db.Scopes(BasicModel, MatchUrlPath(cat)).Select("title").Scan(&catTitle)
			db.Scopes(BasicModel, MatchUrlPath(cat+"/"+book)).Select("title").Scan(&bookTitle)

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
		sendErrorResponsePageData(c, http.StatusNotFound, clientTime, totalCount, page, pageSize, "Keyword Match Error")
		return
	}

	sendSuccessResponsePageData(c, clientTime, results, totalCount, page, pageSize)
}
