package storage

import (
	"Key_Value_Persistant_Storage/internal/fs"
	"Key_Value_Persistant_Storage/internal/models"
	"Key_Value_Persistant_Storage/internal/services"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func removeSingleFile(collection *mongo.Collection, storageItem *models.StorageItem, config *services.Config, fileSystem fs.FileSystem) error {
	err := fileSystem.Remove(storageItem.ValuePath)
	if err == nil {
		collection.DeleteOne(*config.MongoContext, bson.M{
			"key": storageItem.Key,
		})
	}
	return err
}

// todo: make it in CRON not when Get or Write called to improve performance for Get and Write methods
func removeExpired(config *services.Config, fileSystem fs.FileSystem) {
	collection := config.MongoClient.Database(config.StorageDB).Collection(config.StorageDBCollection)
	items, err := collection.Find(*config.MongoContext, bson.M{
		"expires_at": bson.M{
			"$lt": time.Now().Unix(),
		},
	})
	if err != nil {
		return
	}
	var expired models.StorageItem
	for items.Next(*config.MongoContext) {
		err := items.Decode(&expired)
		if err != nil {
			continue
		}
		// delete one by one because we have btree indexes
		// and possible situation with Network FS when removed from database but not removed from FS
		err = removeSingleFile(collection, &expired, config, fileSystem)
	}
}

func Get(c *gin.Context, config *services.Config, fileSystem fs.FileSystem) (*models.StorageItem, error) {

	removeExpired(config, fileSystem)

	key := c.Param("key")
	var storageItem models.StorageItem
	now := time.Now()
	collection := config.MongoClient.Database(config.StorageDB).Collection(config.StorageDBCollection)
	err := collection.FindOne(*config.MongoContext, bson.M{
		"key":        key,
		"expires_at": bson.M{"$gt": now.Unix()},
	}).Decode(&storageItem)
	switch {
	case err == nil:
		return &storageItem, nil
	case errors.Is(err, mongo.ErrNoDocuments):
		return nil, services.ErrorNotFound
	default:
		return nil, err
	}
}

func Write(c *gin.Context, config *services.Config, fileSystem fs.FileSystem) error {

	removeExpired(config, fileSystem)

	key := c.Param("key")
	var queryData struct {
		Extension string `json:"extension"`
		Value     string `json:"value"`
		Ttl       int64  `json:"ttl"`
	}
	if err := c.BindJSON(&queryData); err != nil {
		return services.UserBadRequest
	}
	if len(queryData.Extension) == 0 || len(queryData.Value) == 0 || queryData.Ttl <= 0 {
		return services.UserBadRequest
	}
	collection := config.MongoClient.Database(config.StorageDB).Collection(config.StorageDBCollection)
	fileName := uuid.New().String() + "." + queryData.Extension
	err := fileSystem.Write(fileName, queryData.Value)
	if err != nil {
		return err
	}

	items, err := collection.Find(*config.MongoContext, bson.M{
		"key": key,
	})
	if err != nil {
		return err
	}

	// cannot append new element while collision
	var duplicateItem models.StorageItem
	for items.Next(*config.MongoContext) {
		err := items.Decode(&duplicateItem)
		if err != nil {
			return err
		}
		// delete one by one because we have btree indexes
		// and possible situation with Network FS when removed from database but not removed from FS
		err = removeSingleFile(collection, &duplicateItem, config, fileSystem)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(*config.MongoContext, bson.M{
		"key":        key,
		"value_path": fileName,
		"extension":  queryData.Extension,
		"expires_at": time.Now().Unix() + queryData.Ttl,
	})
	return err
}

func Delete(c *gin.Context, config *services.Config, fileSystem fs.FileSystem) error {

	removeExpired(config, fileSystem)

	key := c.Param("key")
	collection := config.MongoClient.Database(config.StorageDB).Collection(config.StorageDBCollection)
	items, err := collection.Find(*config.MongoContext, bson.M{
		"key": key,
	})
	if err != nil {
		return err
	}

	var duplicateItem models.StorageItem
	deleteCnt := 0
	for items.Next(*config.MongoContext) {
		deleteCnt += 1
		err := items.Decode(&duplicateItem)
		if err != nil {
			return err
		}
		// delete one by one because we have btree indexes
		// and possible situation with Network FS when removed from database but not removed from FS
		err = removeSingleFile(collection, &duplicateItem, config, fileSystem)
		if err != nil {
			return err
		}
	}
	if deleteCnt == 0 {
		return services.ErrorNotFound
	}
	if err != nil {
		return err
	}
	return nil

}
