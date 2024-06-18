package routes

import (
	"docsfly/global"
	"docsfly/handlers"

	"github.com/gin-gonic/gin"
)

func registerRoutes(engine *gin.Engine) {

	engine.POST("/api/"+global.AppConfig.APIVersion+"/ip", handlers.VisitorInsertLog)
	engine.POST("/api/"+global.AppConfig.APIVersion+"/auth/login", handlers.LoginAuth)
	engine.GET("/api/"+global.AppConfig.APIVersion+"/auth/token", handlers.TokenVerify)

	engine.GET("/api/"+global.AppConfig.APIVersion+"/book", handlers.GetBook)
	engine.GET("/api/"+global.AppConfig.APIVersion+"/statistic/book", handlers.GetStatisticBook)
	engine.GET("/api/"+global.AppConfig.APIVersion+"/book/meta", handlers.GetBookMeta)
	engine.PUT("/api/"+global.AppConfig.APIVersion+"/book/meta", handlers.UpdateBookMeta)

	engine.GET("/api/"+global.AppConfig.APIVersion+"/comment", handlers.GetComments)
	engine.POST("/api/"+global.AppConfig.APIVersion+"/comment", handlers.SendComment)

	engine.GET("/api/"+global.AppConfig.APIVersion+"/nav", handlers.GetNav)
	engine.GET("/api/"+global.AppConfig.APIVersion+"/statistic/home", handlers.GetStatisticHome)
	engine.GET("/api/"+global.AppConfig.APIVersion+"/query", handlers.Query)

	engine.GET("/api/"+global.AppConfig.APIVersion+"/app/version", handlers.GetAppVersion)
	engine.GET("/api/"+global.AppConfig.APIVersion+"/rand/nickname", handlers.GetRndName)
	engine.GET("/api/"+global.AppConfig.APIVersion+"/rand/poem", handlers.GetRndPoem)
	engine.GET("/api/"+global.AppConfig.APIVersion+"/rand/post", handlers.GetRndPost)

	engine.GET("/api/"+global.AppConfig.APIVersion+"/post", handlers.GetPost)
	engine.GET("/api/"+global.AppConfig.APIVersion+"/post/html", handlers.GetPostHtml)
	engine.POST("/api/"+global.AppConfig.APIVersion+"/post", handlers.SavePost)
	engine.GET("/api/"+global.AppConfig.APIVersion+"/post/chapter", handlers.GetChapter)

	engine.GET("/api/"+global.AppConfig.APIVersion+"/vendor/yiyan", handlers.GetYiYan)

}
