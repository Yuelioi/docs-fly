package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"docsfly/internal/config"
	_ "docsfly/internal/database"
	"docsfly/internal/router"
)

func main() {
	engine := gin.Default()
	router.SetupRouter(engine)

	engine.Run(fmt.Sprintf("%s:%d", config.Instance.System.Host, config.Instance.System.Port))

	user := engine.Group("User")
	{
		user.GET("/user", func(c *gin.Context) {
		})
	}
}
