package post

import (
	"docsfly/internal/global"

	"github.com/gin-gonic/gin"
)

type BookRouter struct{}

func (*BookRouter) Register(engine *gin.Engine) {
	engine.GET("/"+global.AppConfig.AppConfig.ApiVersion+"/post", GetPost)
	engine.GET("/"+global.AppConfig.AppConfig.ApiVersion+"/post/html", GetPostHtml)
	engine.POST("/"+global.AppConfig.AppConfig.ApiVersion+"/post", SavePost)
	engine.GET("/"+global.AppConfig.AppConfig.ApiVersion+"/post/chapter", GetChapter)
}
