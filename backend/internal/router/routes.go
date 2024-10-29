package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

type RouteRegistrar interface {
	Register(*gin.Engine)
}

func build(engine *gin.Engine, registrars ...RouteRegistrar) {
	for _, registrar := range registrars {
		registrar.Register(engine)
	}
}

func registerRoutes(engine *gin.Engine) {
	build(engine)
	log.Println("路由注册成功")
}
