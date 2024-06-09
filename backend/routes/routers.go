package routes

import (
	"docsfly/handlers"

	"github.com/gin-gonic/gin"
)

func registerRoutes(engine *gin.Engine) {
	engine.GET("/rand/nickname", handlers.GetRndName) // √
	engine.GET("/rand/poem", handlers.GetRndPoem)     // √

	engine.GET("/ip", handlers.GetClientIP)
	engine.POST("/ip", handlers.VisitorInsertLog)

	engine.POST("/auth/login", handlers.LoginAuth)  // √
	engine.GET("/auth/token", handlers.TokenVerify) // √

	engine.GET("/book", handlers.GetBook)

	engine.GET("/book/meta", handlers.GetBookMeta)
	engine.PUT("/book/meta", handlers.SaveBookMeta)
	engine.POST("/book/meta", handlers.SaveBookMeta)

	engine.GET("/post", handlers.GetPost)
	engine.GET("/post/html", handlers.GetPostHtml)
	engine.POST("/post", handlers.SavePost)

	engine.GET("/chapter", handlers.GetChapter)

	engine.GET("/nav", handlers.GetNav) // √

	engine.GET("/statistic/home", handlers.GetStatisticHome)
	engine.GET("/statistic/book", handlers.GetStatisticBook)

	engine.GET("/query", handlers.Query) // √

}
