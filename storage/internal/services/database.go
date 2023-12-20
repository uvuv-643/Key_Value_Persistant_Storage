package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func (config *Config) ConnectDatabase(ctx context.Context) {
	mongoClient, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(config.MongoURL),
	)
	config.MongoClient = mongoClient
	if err != nil {
		log.Fatalf("Failed to connect database")
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Failed to ping database")
	}
	log.Println("connected to database")

	_, err = mongoClient.Database(config.StorageDB).Collection(config.StorageDBCollection).Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"expires_at": 1},
		Options: options.Index(),
	})
	if err != nil {
		fmt.Println("Cannot create index for expires_at")
	}

	_, err = mongoClient.Database(config.StorageDB).Collection(config.StorageDBCollection).Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"key": "hashed"},
		Options: options.Index(),
	})
	if err != nil {
		fmt.Println("Cannot create index for key")
	}

}
