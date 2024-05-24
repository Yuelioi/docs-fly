package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func registerCors(engine *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // 允许的源
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Content-Length"}
	engine.Use(cors.New(config))
}
