package routes

import (
	"docsfly/global"
	"docsfly/handlers"

	"github.com/gin-gonic/gin"
)

func registerRoutes(engine *gin.Engine) {

	engine.POST("/"+global.AppConfig.APIVersion+"/ip", handlers.VisitorInsertLog)
	engine.POST("/"+global.AppConfig.APIVersion+"/auth/login", handlers.LoginAuth)
	engine.GET("/"+global.AppConfig.APIVersion+"/auth/token", handlers.TokenVerify)

	engine.GET("/"+global.AppConfig.APIVersion+"/book", handlers.GetBook)
	engine.GET("/"+global.AppConfig.APIVersion+"/statistic/book", handlers.GetStatisticBook)
	engine.GET("/"+global.AppConfig.APIVersion+"/book/meta", handlers.GetBookMeta)
	engine.PUT("/"+global.AppConfig.APIVersion+"/book/meta", handlers.UpdateBookMeta)

	engine.GET("/"+global.AppConfig.APIVersion+"/comment", handlers.GetComments)
	engine.POST("/"+global.AppConfig.APIVersion+"/comment", handlers.SendComment)

	engine.GET("/"+global.AppConfig.APIVersion+"/nav", handlers.GetNav)
	engine.GET("/"+global.AppConfig.APIVersion+"/statistic/home", handlers.GetStatisticHome)
	engine.GET("/"+global.AppConfig.APIVersion+"/query", handlers.Query)

	engine.GET("/"+global.AppConfig.APIVersion+"/app/version", handlers.GetAppVersion)
	engine.GET("/"+global.AppConfig.APIVersion+"/rand/nickname", handlers.GetRndName)
	engine.GET("/"+global.AppConfig.APIVersion+"/rand/poem", handlers.GetRndPoem)
	engine.GET("/"+global.AppConfig.APIVersion+"/rand/post", handlers.GetRndPost)

	engine.GET("/"+global.AppConfig.APIVersion+"/post", handlers.GetPost)
	engine.GET("/"+global.AppConfig.APIVersion+"/post/html", handlers.GetPostHtml)
	engine.POST("/"+global.AppConfig.APIVersion+"/post", handlers.SavePost)
	engine.GET("/"+global.AppConfig.APIVersion+"/post/chapter", handlers.GetChapter)

	engine.GET("/"+global.AppConfig.APIVersion+"/vendor/yiyan", handlers.GetYiYan)

}
