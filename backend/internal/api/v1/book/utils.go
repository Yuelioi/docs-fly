package book

import (
	"docsfly/internal/models"
	"strings"
	"sync"

	"gorm.io/gorm"
)

var globalCache sync.Map

func getCategoryAndBookByUrl(url string) (string, string) {
	urlList := strings.Split(url, "/")
	if len(urlList) < 2 {
		return "", ""
	}

	return urlList[0], urlList[1]
}

func getFilepathByURL(db *gorm.DB, url string) string {
	var filepath string
	db.Model(models.Entry{}).Where("url = ?", url).Select("filepath").Scan(&filepath)
	return filepath
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
