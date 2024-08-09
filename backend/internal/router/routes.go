package router

import (
	"docsfly/api/v1/auth"
	"docsfly/api/v1/book"
	"docsfly/api/v1/comment"
	"docsfly/api/v1/home"
	"docsfly/api/v1/other"
	"docsfly/api/v1/post"

	"github.com/gin-gonic/gin"
)

type RouteRegistrar interface {
	Register(*gin.Engine)
}

func register(engine *gin.Engine, registrars ...RouteRegistrar) {
	for _, registrar := range registrars {
		registrar.Register(engine)
	}
}

func registerRoutes(engine *gin.Engine) {

	register(engine,
		&auth.AuthRouter{},
		&book.BookRoutes{},
		&comment.CommentRouter{},
		&home.HomeRouter{},
		&other.OtherRouter{},
		&post.BookRouter{},
	)

}
