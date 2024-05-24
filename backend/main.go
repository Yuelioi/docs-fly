package main

import (
	"github.com/gin-gonic/gin"

	"docsfly/global"
	"docsfly/routes"
)

func main() {

	engine := gin.Default()

	routes.SetupRouter(engine)
	routes.SetupRouterPlugins(engine)

	engine.Run(":" + global.AppConfig.Addr)

}
