package services

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Config struct {
	Env                 string `required:"true"`
	ServicePort         string `required:"true"`
	MongoURL            string `required:"true"`
	StorageDB           string `required:"true"`
	StorageDBCollection string `required:"true"`
	StorageFiles        string `required:"true"`

	Logger       *zap.SugaredLogger `ignored:"true"`
	MongoClient  *mongo.Client      `ignored:"true"`
	MongoContext *context.Context   `ignored:"true"`
}
