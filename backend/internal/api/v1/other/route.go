package other

import (
	"docsfly/internal/global"

	"github.com/gin-gonic/gin"
)

type OtherRouter struct{}

func (*OtherRouter) Register(engine *gin.Engine) {
	engine.POST("/"+global.AppConfig.AppConfig.ApiVersion+"/ip", VisitorInsertLog)

	engine.GET("/"+global.AppConfig.AppConfig.ApiVersion+"/app/version", GetAppVersion)
	engine.GET("/"+global.AppConfig.AppConfig.ApiVersion+"/rand/nickname", GetRndName)
	engine.GET("/"+global.AppConfig.AppConfig.ApiVersion+"/rand/poem", GetRndPoem)
	engine.GET("/"+global.AppConfig.AppConfig.ApiVersion+"/rand/post", GetRndPost)

	engine.GET("/"+global.AppConfig.AppConfig.ApiVersion+"/vendor/yiyan", GetYiYan)
}
