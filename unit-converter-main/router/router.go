package router

import (
	"log/slog"

	"github.com/fentezi/unit-converter/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	log     *slog.Logger
	handler *handler.Handler
}

func NewRouter(log *slog.Logger, handler *handler.Handler) *Router {
	return &Router{
		log:     log,
		handler: handler,
	}
}

func (r *Router) initRouter() *gin.Engine {
	g := gin.Default()

	g.POST("/convert", r.handler.Convert)

	return g
}

func (r *Router) StartServer() {
	r.initRouter().Run(":8080")
}
