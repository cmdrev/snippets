package databasefx

import (
	"context"
	"fmt"
	"github.com/cmdrev/snippets/config"
	"github.com/cmdrev/snippets/database"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
	"go.uber.org/fx"
)

func NewMongoDBClient(lc fx.Lifecycle, config database.MongoDBConfig) (*mongo.Client, error) {
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

func MongoClientWithDatabase(client *mongo.Client, config database.MongoDBConfig) (*mongo.Database, error) {
	cs, err := connstring.ParseAndValidate(config.ConnectionURI())
	if err != nil {
		return nil, fmt.Errorf("invalid mongodb connection uri: %w", err)
	}
	return client.Database(cs.Database), nil
}

var MongoDBModule = fx.Module("mongodb",
	fx.Provide(config.NewMongoDBConfig),
	fx.Provide(NewMongoDBClient),
	fx.Provide(MongoClientWithDatabase),
)
