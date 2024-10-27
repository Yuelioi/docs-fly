package router

import (
	"docsfly/internal/api/v1/book"
	"docsfly/internal/api/v1/comment"
	"docsfly/internal/api/v1/home"
	"docsfly/internal/api/v1/other"
	"docsfly/internal/api/v1/post"
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
	build(engine,
		&home.HomeRouter{},
		&book.BookRoutes{},
		&comment.CommentRouter{},
		&other.OtherRouter{},
		&post.BookRouter{},
	)
	log.Println("路由注册成功")
}
