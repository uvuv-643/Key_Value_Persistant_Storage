package routes

import (
	"Key_Value_Persistant_Storage/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupDefaultEndpoints(r *gin.Engine, conf *services.Config) {
	r.GET("/status", func(c *gin.Context) {
		var msg string
		msg = fmt.Sprintf("Pong! I am %s. Port is %s.", conf.ServiceName, conf.ServicePort)
		c.JSON(200, gin.H{
			"success": true,
			"message": msg,
		})
	})
}

func AddRoutes(r *gin.Engine, config *services.Config) {
	r.GET("/:key", func(c *gin.Context) {
		c.JSON(200, "Success")
	})
	r.PUT("/:key", func(c *gin.Context) {
		c.JSON(200, "Success")
	})
	r.DELETE("/:key", func(c *gin.Context) {
		c.JSON(200, "Success")
	})
}
