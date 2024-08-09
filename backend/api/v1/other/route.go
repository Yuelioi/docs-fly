package other

import (
	"docsfly/internal/global"

	"github.com/gin-gonic/gin"
)

type OtherRouter struct{}

func (*OtherRouter) Register(engine *gin.Engine) {
	engine.POST("/"+global.AppConfig.APIVersion+"/ip", VisitorInsertLog)

	engine.GET("/"+global.AppConfig.APIVersion+"/app/version", GetAppVersion)
	engine.GET("/"+global.AppConfig.APIVersion+"/rand/nickname", GetRndName)
	engine.GET("/"+global.AppConfig.APIVersion+"/rand/poem", GetRndPoem)
	engine.GET("/"+global.AppConfig.APIVersion+"/rand/post", GetRndPost)

	engine.GET("/"+global.AppConfig.APIVersion+"/vendor/yiyan", GetYiYan)
}
