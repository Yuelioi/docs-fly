package handlers

import (
	"docsfly/database"
	"docsfly/models"
	"net/http"
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

	category := c.Query("category")
	book := c.Query("book")
	url := c.Query("url")

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
	}

	db.Model(&models.Visitor{}).Create(&vs)

	result := VisitorInfo{}
	result.IP = clientIP

	// 返回 IP 地址给客户端
	c.JSON(http.StatusOK, result)
}
