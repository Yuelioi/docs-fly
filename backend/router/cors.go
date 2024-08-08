package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func registerCors(engine *gin.Engine) {
	config := cors.DefaultConfig()
	// 指定允许访问的源
	config.AllowOrigins = []string{"http://localhost:5173", "https://docs.yuelili.com"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Content-Length"}
	// 允许携带凭证（例如 Cookies）
	config.AllowCredentials = true
	engine.Use(cors.New(config))
}
