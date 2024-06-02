package handlers

// HomePage

import (
	"docsfly/database"
	"docsfly/models"
	"docsfly/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetSearchOptions(c *gin.Context) {

	db, err := database.DbManager.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	var cats []models.Category

	var options []models.SearchOption

	db.Model(&models.Category{}).Preload("Books").Find(&cats)

	// 遍历分类
	for _, cat := range cats {

		option := models.SearchOption{}
		option.MetaData = cat.MetaData

		// 遍历书籍

		options = append(options, option)
	}

	c.JSON(http.StatusOK, options)

}

func Query(c *gin.Context) {
	category := c.Query("category")
	book := c.Query("book")

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

	var catInfo models.Category
	var documentInfo models.Document

	var documents []models.Document

	if category != "" && book != "" {
		db.Model(&catInfo).Where("identity = ?", category).First(&catInfo)
	} else {
		db.Model(&documentInfo).Preload("Category").Preload("Book").Preload("Chapter").Preload("Section").Where("content LIKE ?", "%"+keyword+"%").Limit(limit).Find(&documents)

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

		dsData := models.SearchData{
			Locale:        document.Locale,
			DocumentName:  document.Name,
			DocumentTitle: document.Title,
			Content:       nearbyText,
		}

		// 添加额外条件
		if category != "" && book != "" {
			dsData.CategoryName = catInfo.Name
			dsData.CategoryTitle = catInfo.Title

		} else {
			dsData.CategoryName = document.Category.Name
			dsData.CategoryTitle = document.Category.Title

		}

		searchResult.Result = append(searchResult.Result, dsData)
	}

	searchResult.SearchTime = utils.DurationToString(time.Since(start))

	c.JSON(http.StatusOK, searchResult)
}
