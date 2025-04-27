package config

import (
	"fmt"
	"os"
)

type MongoDBConfig struct {
	URI string
}

func NewMongoDBConfig() (*MongoDBConfig, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return nil, fmt.Errorf("MONGO_URI not set")
	}

	return &MongoDBConfig{
		URI: uri,
	}, nil
}
