package main

import (
	"github.com/gin-gonic/gin"

	"docsfly/internal/config"
	_ "docsfly/internal/database"
	"docsfly/internal/router"
)

func main() {
	engine := gin.Default()
	router.SetupRouter(engine)
	engine.Run(":" + config.AppConfig.ServerConfig.Addr)
}
