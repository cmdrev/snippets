package serverfx

import (
	"context"
	"errors"
	"github.com/cmdrev/snippets/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

type (
	Result struct {
		fx.Out
		ServerRunError <-chan error `name:"ServerRunError"`
		HttpServer     *http.Server
	}

	Params struct {
		fx.In
		LC      fx.Lifecycle
		Logger  *zap.Logger
		Config  config.Http
		Handler http.Handler
	}
)

func NewHttpServer(params Params) (Result, error) {
	errChan := make(chan error, 1)
	server := http.Server{Addr: params.Config.Address(), Handler: params.Handler}
	params.LC.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			params.Logger.Info("starting HTTP server at: " + params.Config.Address())
			go func() {
				defer close(errChan)
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					errChan <- err
				}
			}()
			select {
			case err := <-errChan:
				return err
			default:
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
	return Result{
		ServerRunError: errChan,
		HttpServer:     &server,
	}, nil
}
