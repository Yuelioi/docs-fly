package post

import (
	"docsfly/internal/global"

	"github.com/gin-gonic/gin"
)

type BookRouter struct{}

func (*BookRouter) Register(engine *gin.Engine) {
	engine.GET("/"+global.AppConfig.APIVersion+"/post", GetPost)
	engine.GET("/"+global.AppConfig.APIVersion+"/post/html", GetPostHtml)
	engine.POST("/"+global.AppConfig.APIVersion+"/post", SavePost)
	engine.GET("/"+global.AppConfig.APIVersion+"/post/chapter", GetChapter)
}
