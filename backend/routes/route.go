package routes

import "github.com/gin-gonic/gin"

func SetupRouter(engine *gin.Engine) {

	// 设置跨域
	registerCors(engine)

	// 设置路由
	registerRoutes(engine)

	// 注册插件
	registerPlugins(engine)

}
