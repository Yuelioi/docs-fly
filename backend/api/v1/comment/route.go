package comment

import (
	"docsfly/internal/global"

	"github.com/gin-gonic/gin"
)

type CommentRouter struct {
}

func (cr *CommentRouter) Register(engine *gin.Engine) {
	engine.GET("/"+global.AppConfig.AppConfig.ApiVersion+"/comment", GetComments)
	engine.POST("/"+global.AppConfig.AppConfig.ApiVersion+"/comment", SendComment)
}
