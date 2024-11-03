package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"docsfly/internal/config"
	_ "docsfly/internal/database"
	"docsfly/internal/router"
	"docsfly/pkg/logger"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func main() {
	r := gin.Default()
	router.SetupRouter(r)

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	user := r.Group("User")
	{
		user.GET("/user", func(c *gin.Context) {

			session := sessions.Default(c)

			if session.Get("hello") != "world" {
				session.Set("hello", "world")
				session.Save()
			}
		})
	}

	r.Run(fmt.Sprintf("%s:%d", config.Instance.System.Host, config.Instance.System.Port))

}
