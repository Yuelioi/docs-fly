package handlers

import (
	"docsfly/global"
	"time"

	"github.com/gin-gonic/gin"
)

func GetRndName(c *gin.Context) {
	sendSuccessResponse(c, time.Now(), RndName())
}

func GetRndPoem(c *gin.Context) {
	sendSuccessResponse(c, time.Now(), RndPoem())
}

func GetAppVersion(c *gin.Context) {
	sendSuccessResponse(c, time.Now(), global.AppConfig.AppVersion)

}
