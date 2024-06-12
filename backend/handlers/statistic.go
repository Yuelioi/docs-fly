package handlers

import (
	"docsfly/database"
	"docsfly/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetStatisticHome(c *gin.Context) {

	db, err := database.DbManager.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	var (
		BooksCount             int64
		DocumentsCount         int64
		HistoricalVisitorCount int64
		TodayVisitorCount      int64
	)

	db.Model(models.Entry{}).Where("depth = ?", 1).Count(&BooksCount)
	db.Model(models.Entry{}).Count(&DocumentsCount)

	db.Model(models.Visitor{}).Count(&HistoricalVisitorCount)

	// today := time.Now().Format("2006-01-02") 不能用(DATE(time))
	today := time.Now().Truncate(24 * time.Hour)
	db.Debug().Model(models.Visitor{}).Where("time > ?", today).Count(&TodayVisitorCount)

	stats := models.Statistic{
		BookCount:              BooksCount,
		DocumentCount:          DocumentsCount,
		HistoricalVisitorCount: HistoricalVisitorCount,
		TodayVisitorCount:      TodayVisitorCount,
	}

	c.JSON(http.StatusOK, stats)
}
func GetStatisticBook(c *gin.Context) {

	slug := c.Query("slug")
	locale := c.Query("locale")

	db, err := database.DbManager.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	var bookReadCount int64
	var chapterCount int64
	var documentCount int64

	//	 获取阅读量
	// db.Model(models.Visitor{}).Where("category = ?", category).Where("book = ?", book).Count(&bookReadCount)

	// 获取章节数量
	db.Model(models.Entry{}).Where("filepath Like ?", slug+"/"+locale+"%").Count(&chapterCount)

	// 获取书籍数量
	db.Model(models.Entry{}).Where("filepath Like ?", slug+"/"+locale+"%").Count(&documentCount)

	type bookCount struct {
		ReadCount     int64 `json:"read_count"`
		ChapterCount  int64 `json:"chapter_count"`
		DocumentCount int64 `json:"document_count"`
	}

	c.JSON(http.StatusOK,
		bookCount{
			ReadCount:     bookReadCount,
			ChapterCount:  chapterCount,
			DocumentCount: documentCount,
		})
}
