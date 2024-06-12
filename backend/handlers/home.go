package handlers

// HomePage

import (
	"docsfly/database"
	"docsfly/models"
	"docsfly/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取顶部导航栏信息
func GetNav(c *gin.Context) {

	navs := []models.Nav{}
	db, err := database.DbManager.Connect()

	if err != nil {
		return
	}
	var cats []models.Entry
	var books []models.Entry
	db.Model(models.Entry{}).Where("depth = ?", 0).Find(&cats)
	db.Model(models.Entry{}).Where("depth = ?", 1).Find(&books)

	for _, cat := range cats {

		nav := models.Nav{}
		nav.MetaData = cat.MetaData

		for _, book := range books {

			if strings.HasPrefix(book.Filepath, cat.Filepath) {
				nav.Children = append(nav.Children, book.MetaData)
			}
		}
		navs = append(navs, nav)
	}

	c.JSON(http.StatusOK, navs)

}

func Query(c *gin.Context) {
	slug := c.Query("slug")
	keyword := c.Query("keyword")

	start := time.Now()

	db, err := database.DbManager.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	// 结果数量
	limit := 20

	// 内容截取长度
	count := 30

	var documentInfo models.Entry
	var documents []models.Entry

	if slug != "" {
		db.Model(&documentInfo).Preload("Entry").Where("filepath LIKE ?", slug+"%").Where("content LIKE ?", "%"+keyword+"%").Limit(limit).Find(&documents)
	} else {
		db.Model(&documentInfo).Preload("Entry").Where("content LIKE ?", "%"+keyword+"%").Limit(limit).Find(&documents)
	}

	if len(documents) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No documents found"})
		return
	}

	var searchResult models.SearchResult

	for _, document := range documents {

		plaintext, err := ExtractPlainText(document.Content)
		if err != nil {
			return
		}

		runeSlice := []rune(*plaintext)
		keywordIndex := IndexOfKeywordInRuneSlice(runeSlice, keyword)

		if keywordIndex == -1 {
			continue
		}

		start := keywordIndex - count
		if start < 0 {
			start = 0
		}

		// 确定截取的结束位置
		end := keywordIndex + count
		if end > len(runeSlice) {
			end = len(runeSlice)
		}

		runeText := runeSlice[start:end]

		nearbyText := string(runeText)

		cat, book, locale, ok := utils.Filepath2Params(document.Filepath)

		if ok {

			var catTitle string
			var bookTitle string
			db.Model(&models.Entry{}).Where("filepath = ?", cat).Select("title").Scan(&catTitle)
			db.Model(&models.Entry{}).Where("filepath = ?", cat+"/"+book).Select("title").Scan(&bookTitle)

			dsData := models.SearchData{
				Url:           document.WebPath,
				CategoryTitle: catTitle,
				BookTitle:     bookTitle,
				Locale:        locale,
				DocumentTitle: strings.Replace(document.Title, ".md", "", 1),
				Content:       nearbyText,
			}
			searchResult.Result = append(searchResult.Result, dsData)
		}

	}

	searchResult.SearchTime = utils.DurationToString(time.Since(start))

	c.JSON(http.StatusOK, searchResult)
}
