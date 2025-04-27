package db

import (
	"context"
	"fmt"
	"quizy-be/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitializeMongoDB() (func(), error) {
	cfg, err := config.NewMongoDBConfig()
	if err != nil {
		return nil, fmt.Errorf("database config error: %w", err)
	}

	if err := connectMongo(cfg.URI); err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	return disconnectMongo, nil
}

func connectMongo(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	return Client.Ping(ctx, nil)
}

func disconnectMongo() {
	if Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = Client.Disconnect(ctx)
	}
}
