package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

const DatabaseConnectionTimeoutSeconds = 3

func (config *Config) ConnectDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), DatabaseConnectionTimeoutSeconds*time.Second)
	mongoClient, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(config.MongoURL),
	)
	config.MongoClient = mongoClient
	defer func() {
		cancel()
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect database")
		}
	}()

	if err != nil {
		log.Fatalf("Failed to connect database")
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Failed to ping database")
	}
	log.Println("connected to database")

}
