package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (config *Config) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
