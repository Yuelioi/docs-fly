package main

import (
	"github.com/gin-gonic/gin"

	"docsfly/internal/global"
	"docsfly/internal/router"
)

func main() {

	engine := gin.Default()
	router.SetupRouter(engine)
	engine.Run(":" + global.AppConfig.Addr)

}
