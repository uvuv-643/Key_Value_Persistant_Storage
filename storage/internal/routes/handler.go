package routes

import (
	"Key_Value_Persistant_Storage/internal/middleware"
	"Key_Value_Persistant_Storage/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewHandler(conf *services.Config) *gin.Engine {
	h := gin.New()

	h.RedirectTrailingSlash = false
	h.RedirectFixedPath = false

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	h.Use(cors.New(corsConfig))

	h.Use(gin.CustomRecovery(conf.PanicRecovery))

	requestLogger := middleware.NewRequestLogger(conf.Logger)
	h.Use(requestLogger.Middleware)

	SetupDefaultEndpoints(h, conf)
	AddRoutes(h, conf)

	return h
}
