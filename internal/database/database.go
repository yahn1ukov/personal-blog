package database

import (
	"context"

	"github.com/yahn1ukov/personal-blog/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.uber.org/fx"
)

type Database struct {
	*mongo.Client
}

func New(cfg *config.Config) (*Database, error) {
	client, err := mongo.Connect(options.Client().ApplyURI(cfg.Database.Mongo.URL))
	if err != nil {
		return nil, err
	}

	return &Database{
		client,
	}, nil
}

func Run(lc fx.Lifecycle, database *Database) {
	fx.Annotate(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return database.Ping(ctx, readpref.Primary())
		},
		OnStop: func(ctx context.Context) error {
			return database.Disconnect(ctx)
		},
	})
}
