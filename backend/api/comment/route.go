package comment

import (
	"docsfly/global"

	"github.com/gin-gonic/gin"
)

type CommentRouter struct {
}

func (cr *CommentRouter) Register(engine *gin.Engine) {
	engine.GET("/"+global.AppConfig.APIVersion+"/comment", GetComments)
	engine.POST("/"+global.AppConfig.APIVersion+"/comment", SendComment)
}
