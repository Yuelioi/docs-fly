package book

import (
	"docsfly/internal/global"

	"github.com/gin-gonic/gin"
)

type BookRoutes struct {
}

func (br *BookRoutes) Register(engine *gin.Engine) {
	engine.GET("/"+global.AppConfig.APIVersion+"/book", Book)
	engine.GET("/"+global.AppConfig.APIVersion+"/book/statistic", BookStatistic)
	engine.GET("/"+global.AppConfig.APIVersion+"/book/meta", BookMeta)
	engine.PUT("/"+global.AppConfig.APIVersion+"/book/meta", UpdateBookMeta)
}
