package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig interface {
	ConnectionURI() string
}

func NewMongoDbSession(ctx context.Context, config MongoDBConfig) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.ConnectionURI()))
	if err != nil {
		return nil, fmt.Errorf("could not connect to MongoDB : %w", err)
	}
	return client, nil
}

func CloseMongoDbSession(ctx context.Context, client *mongo.Client) error {
	err := client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("could not close MongoDB connection: %w", err)
	}
	return nil
}
