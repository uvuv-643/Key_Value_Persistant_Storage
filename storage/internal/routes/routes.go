package routes

import (
	"Key_Value_Persistant_Storage/internal/controllers"
	"Key_Value_Persistant_Storage/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupDefaultEndpoints(r *gin.Engine, conf *services.Config) {
	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})
}

func AddRoutes(r *gin.Engine, config *services.Config) {
	r.GET("/:key", func(c *gin.Context) {
		controllers.GetFromStorage(c, config)
	})
	r.POST("/:key", func(c *gin.Context) {
		controllers.WriteToStorage(c, config)
	})
	r.DELETE("/:key", func(c *gin.Context) {
		controllers.DeleteFromStorage(c, config)
	})
}
