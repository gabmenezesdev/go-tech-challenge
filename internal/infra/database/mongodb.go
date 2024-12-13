package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL      = "MONGODB_URL"
	MONGODB_DATABASE = "MONGODB_DATABASE"
)

var client *mongo.Client

func InitConnection() (*mongo.Database, error) {
	mongodbUri := os.Getenv(MONGODB_URL)
	mongodbDatabase := os.Getenv(MONGODB_DATABASE)
	fmt.Println(mongodbUri)
	fmt.Println(mongodbDatabase)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodbDatabase), nil
}

func CloseConnection() error {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return client.Disconnect(ctx)
	}
	return nil
}
