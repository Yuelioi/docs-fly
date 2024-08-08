package router

import (
	"docsfly/api/auth"
	"docsfly/api/book"
	"docsfly/api/comment"
	"docsfly/api/home"
	"docsfly/api/other"
	"docsfly/api/post"

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
