package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (config *Config) PanicRecovery(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		config.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
		})
		return
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}
