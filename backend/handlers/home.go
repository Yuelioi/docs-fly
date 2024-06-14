package handlers

// HomePage

import (
	"docsfly/models"
	"docsfly/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 获取顶部导航栏信息
func GetNav(c *gin.Context) {

	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	clientTime := currentTime()
	db := dbContext.(*gorm.DB)

	var cats []models.Entry
	var books []models.Entry
	var navs []Nav

	if err := db.Scopes(BasicModel, FindCategory, FindFolder).Find(&cats).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Failed to retrieve categories")
	}
	if err := db.Scopes(BasicModel, FindBook, FindFolder).Find(&books).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Failed to retrieve books")
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

func Query(c *gin.Context) {

	dbContext, exists := c.Get("db")
	if !exists {
		return
	}

	clientTime := currentTime()
	db := dbContext.(*gorm.DB)

	fullPath := c.Query("fullPath")
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

	// 结果数量
	splitContentSize := 30

	var documents []models.Entry
	var totalEntries int64

	// 根据查询条件获取结果
	query := db.Scopes(BasicModel)
	if fullPath != "" {
		query = query.Scopes(HasPrefixPath(fullPath))
	}

	query = query.Where("content LIKE ?", "%"+keyword+"%")
	query.Model(&models.Entry{}).Count(&totalEntries)

	query = query.Offset(offset).Limit(pageSize)
	query.Find(&documents)

	if len(documents) == 0 {
		sendErrorResponse(c, http.StatusNotFound, clientTime, "No documents found")
	}

	var results []SearchData

	for _, document := range documents {

		plaintext, err := extractPlainText(document.Content)
		if err != nil {
			return
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
			db.Scopes(BasicModel).Where("filepath = ?", cat).Select("title").Scan(&catTitle)
			db.Scopes(BasicModel).Where("filepath = ?", cat+"/"+book).Select("title").Scan(&bookTitle)

			dsData := SearchData{
				Url:           document.URLPath,
				CategoryTitle: catTitle,
				BookTitle:     bookTitle,
				Locale:        locale,
				DocumentTitle: strings.Replace(document.Title, ".md", "", 1),
				Content:       nearbyText,
			}
			results = append(results, dsData)
		}

	}

	c.JSON(http.StatusOK, results)
}
