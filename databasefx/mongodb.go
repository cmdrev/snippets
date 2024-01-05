package databasefx

import (
	"context"
	"github.com/cmdrev/snippets/config"
	"github.com/cmdrev/snippets/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

func NewMongoDBSession(lc fx.Lifecycle, config database.MongoDBConfig) (*mongo.Client, error) {
	client, err := database.NewMongoDbSession(context.TODO(), config)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return database.CloseMongoDbSession(ctx, client)
		},
	})
	return client, nil
}

var MongoDBSession = fx.Module("mongodb-session",
	fx.Provide(config.NewMongoDBConfig),
	fx.Provide(NewMongoDBSession),
)
