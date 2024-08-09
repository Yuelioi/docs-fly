package book

import (
	"docsfly/internal/global"

	"github.com/gin-gonic/gin"
)

type BookRoutes struct {
}

func (br *BookRoutes) Register(engine *gin.Engine) {
	engine.GET("/"+global.AppConfig.AppConfig.ApiVersion+"/book", Book)
	engine.GET("/"+global.AppConfig.AppConfig.ApiVersion+"/book/statistic", BookStatistic)
	engine.GET("/"+global.AppConfig.AppConfig.ApiVersion+"/book/meta", BookMeta)
	engine.PUT("/"+global.AppConfig.AppConfig.ApiVersion+"/book/meta", UpdateBookMeta)
}
