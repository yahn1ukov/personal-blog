package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yahn1ukov/personal-blog/internal/config"
	"github.com/yahn1ukov/personal-blog/internal/http/router"
	"go.uber.org/fx"
)

func Run(lc fx.Lifecycle, cfg *config.Config, router *router.Router) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HTTP.Port),
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
