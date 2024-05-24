package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {

	// 设置跨域
	setCors(engine)

	// 设置路由
	setupIPRoutes(engine)
	setupAuthRoutes(engine)
	setupBookRoutes(engine)
	setupPostRoutes(engine)
	setupChapterRoutes(engine)
	setupNavRoutes(engine)
	setupStatisticRoutes(engine)
	setupHomeRoutes(engine)

}
