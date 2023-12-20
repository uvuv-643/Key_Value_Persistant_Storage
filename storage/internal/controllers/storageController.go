package controllers

import (
	"Key_Value_Persistant_Storage/internal/fs"
	"Key_Value_Persistant_Storage/internal/services"
	"Key_Value_Persistant_Storage/internal/storage"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var fileSystem fs.FileSystem

func createFileSystem(storageFilesPath string) fs.FileSystem {
	return fs.NewLocalFileSystem(storageFilesPath)
}

func GetFromStorage(c *gin.Context, config *services.Config) {
	fileSystem = createFileSystem(config.StorageFiles)
	storageItem, err := storage.Get(c, config, fileSystem)
	if errors.Is(err, services.ErrorNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
		})
		return
	} else if errors.Is(err, services.UserBadRequest) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	var fileSystem fs.FileSystem = fs.NewLocalFileSystem(config.StorageFiles)
	bytes, err := fileSystem.Read(storageItem.ValuePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.Header("Content-Type", http.DetectContentType(bytes))
	_, err = c.Writer.Write(bytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
}

func WriteToStorage(c *gin.Context, config *services.Config) {
	fileSystem = createFileSystem(config.StorageFiles)
	err := storage.Write(c, config, fileSystem)
	if errors.Is(err, services.UserBadRequest) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
}

func DeleteFromStorage(c *gin.Context, config *services.Config) {
	fileSystem = createFileSystem(config.StorageFiles)
	err := storage.Delete(c, config, fileSystem)
	if err != nil {
		if errors.Is(err, services.ErrorNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "there is no such element",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
