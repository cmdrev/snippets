package serverfx

import (
	"github.com/cmdrev/snippets/config"
	server "github.com/cmdrev/snippets/sever"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

type Middleware interface {
	HandleFunc(c *gin.Context)
}

var GinDefaultModule = fx.Module("ginserver",
	fx.Provide(fx.Annotate(
		config.NewGinHttpConfig,
		fx.As(new(config.Http)),
	)),
	fx.Provide(fx.Annotate(
		server.NewGinHttpRouter,
		fx.As(new(http.Handler)),
		fx.As(new(gin.IRouter)),
	)),
	fx.Provide(NewHttpServer),
)
