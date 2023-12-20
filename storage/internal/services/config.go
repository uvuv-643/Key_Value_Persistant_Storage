package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Config struct {
	Env         string `required:"true"`
	ServiceName string `required:"true"`
	ServicePort string `required:"true"`
	MongoURL    string `required:"true"`

	Logger      *zap.SugaredLogger `ignored:"true"`
	MongoClient *mongo.Client      `ignored:"true"`
}
