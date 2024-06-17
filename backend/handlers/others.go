package handlers

import (
	"docsfly/global"
	"docsfly/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRndName(c *gin.Context) {
	sendSuccessResponse(c, time.Now(), RndName())
}

func GetRndPost(c *gin.Context) {
	clientTime := currentTime()
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}
	var doc models.Entry

	db := dbContext.(*gorm.DB)
	if err := db.Scopes(BasicModel, FindFile).Order("RANDOM()").First(&doc).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Could not retrieve a random post")
		return
	}

	sendSuccessResponse(c, time.Now(), doc.MetaData)
}

func GetRndPoem(c *gin.Context) {
	sendSuccessResponse(c, time.Now(), RndPoem())
}

func GetAppVersion(c *gin.Context) {
	sendSuccessResponse(c, time.Now(), global.AppConfig.AppVersion)

}

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
	db.Scopes(BasicModel, MatchUrlPath(url)).Count(&count)

	if count == 0 {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Can't find target link")
		return
	}

	urlList := strings.Split(url, "/")

	var category, book, locale string

	if len(urlList) > 2 {
		category = urlList[0]
		book = urlList[1]
		locale = urlList[2]
	} else {
		sendErrorResponse(c, http.StatusInternalServerError, clientTime, "Can't find target path")
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
	sendSuccessResponse(c, clientTime, gin.H{"message": "success"})
}
