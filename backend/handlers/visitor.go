package handlers

import (
	"docsfly/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type VisitorInfo struct {
	IP string
}

func VisitorInsertLog(c *gin.Context) {

	clientTime := currentTime()
	url := c.Query("url")
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}
	db := dbContext.(*gorm.DB)

	var count int64
	db.Model(models.Entry{}).Scopes(MatchUrlPath(url)).First(count)

	if count == 0 {
		return
	}

	urlList := strings.Split(url, "/")

	var category, book, locale string

	if len(urlList) > 2 {
		category = urlList[0]
		book = urlList[1]
		locale = urlList[2]
	} else {
		return
	}

	today := time.Now().Local()

	vs := models.Visitor{
		IP:       c.ClientIP(),
		URL:      url,
		Time:     today,
		Category: category,
		Book:     book,
		Locale:   locale,
	}

	db.Model(&models.Visitor{}).Create(&vs)

	// 返回 IP 地址给客户端
	sendSuccessResponse(c, clientTime, "")
}
