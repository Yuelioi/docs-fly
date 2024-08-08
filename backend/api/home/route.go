package home

import (
	"docsfly/global"

	"github.com/gin-gonic/gin"
)

type HomeRouter struct {
}

func (HomeRouter) Register(engine *gin.Engine) {
	engine.GET("/"+global.AppConfig.APIVersion+"/nav", GetNav)
	engine.GET("/"+global.AppConfig.APIVersion+"/statistic/home", GetStatisticHome)
	engine.GET("/"+global.AppConfig.APIVersion+"/query", Query)
}
