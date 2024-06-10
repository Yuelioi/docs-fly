package handlers

import (
	"docsfly/database"
	"docsfly/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type VisitorInfo struct {
	IP string
}

func GetClientIP(c *gin.Context) {
	// 获取客户端的 IP 地址
	clientIP := c.ClientIP()

	// 返回 IP 地址给客户端
	c.JSON(http.StatusOK, gin.H{"ip": clientIP})
}

func VisitorInsertLog(c *gin.Context) {
	// 获取客户端的 IP 地址
	clientIP := c.ClientIP()

	url := c.Query("url")

	urlList := strings.Split(url, "/")

	var category, book, locale string

	if len(urlList) > 2 {
		category = urlList[0]
		book = urlList[1]
		locale = urlList[2]
	} else {
		return
	}

	db, err := database.DbManager.Connect()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed load database"})
		return
	}

	today := time.Now().Local()

	vs := models.Visitor{
		IP:       clientIP,
		URL:      url,
		Time:     today,
		Category: category,
		Book:     book,
		Locale:   locale,
	}

	db.Model(&models.Visitor{}).Create(&vs)

	result := VisitorInfo{}
	result.IP = clientIP

	// 返回 IP 地址给客户端
	c.JSON(http.StatusOK, result)
}
