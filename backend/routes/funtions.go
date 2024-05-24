package routes

import (
	"docsfly/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setCors(engine *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // 允许的源
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Content-Length"}
	engine.Use(cors.New(config))
}
func setupIPRoutes(engine *gin.Engine) {
	engine.GET("/ip", handlers.GetClientIP)
	engine.POST("/ip", handlers.VisitorInsertLog)
}

func setupAuthRoutes(engine *gin.Engine) {
	engine.POST("/auth/login", handlers.LoginAuth)
	engine.GET("/auth/token", handlers.TokenVerify)
}

func setupBookRoutes(engine *gin.Engine) {
	engine.GET("/book", handlers.GetBook)
}

func setupPostRoutes(engine *gin.Engine) {
	engine.GET("/post", handlers.GetPost)
	engine.GET("/post/html", handlers.GetPostHtml)
	engine.POST("/post", handlers.SavePost)
}

func setupChapterRoutes(engine *gin.Engine) {
	engine.GET("/chapter", handlers.GetChapter)
}

func setupNavRoutes(engine *gin.Engine) {
	engine.GET("/nav", handlers.GetNav)
}

func setupStatisticRoutes(engine *gin.Engine) {
	engine.GET("/statistic/home", handlers.GetStatisticHome)
	engine.GET("/statistic/book", handlers.GetStatisticBook)

}
func setupHomeRoutes(engine *gin.Engine) {
	engine.GET("/search_options", handlers.GetSearchOptions)
	engine.GET("/query", handlers.Query)

}
