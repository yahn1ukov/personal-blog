package app

import (
	"github.com/yahn1ukov/personal-blog/internal/config"
	"github.com/yahn1ukov/personal-blog/internal/database"
	"github.com/yahn1ukov/personal-blog/internal/http"
	"github.com/yahn1ukov/personal-blog/internal/http/handler"
	"github.com/yahn1ukov/personal-blog/internal/http/router"
	"github.com/yahn1ukov/personal-blog/internal/repository"
	"github.com/yahn1ukov/personal-blog/internal/service"
	"go.uber.org/fx"
)

func New(configPath string) *fx.App {
	return fx.New(
		fx.Provide(
			func() (*config.Config, error) {
				return config.New(configPath)
			},
			database.New,
		),

		fx.Provide(
			fx.Annotate(repository.New, fx.As(new(repository.Repository))),
			fx.Annotate(service.New, fx.As(new(service.Service))),
		),

		fx.Provide(handler.New, router.New),

		fx.Invoke(database.Run, http.Run),
	)
}
