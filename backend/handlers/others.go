package handlers

import (
	"docsfly/global"
	"docsfly/models"
	"net/http"
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
