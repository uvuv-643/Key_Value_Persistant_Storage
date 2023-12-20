package logging

import (
	"github.com/blendle/zapdriver"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	RequestIDKey    = "X-Request-ID"
	LogRequestIDKey = "req-id"
)

func GetLogger() *zap.SugaredLogger {
	// The sugared logger allows untyped arguments, similar to "fmt.Printf"
	return zap.S()
}

func GetRequestLogger(c *gin.Context) *zap.SugaredLogger {
	l := zap.S()
	if c != nil {
		if ctxRqId, ok := c.Value(RequestIDKey).(string); ok {
			l = l.With(zap.String(LogRequestIDKey, ctxRqId))
		}
	}
	return l
}

func SetupLogger(environment string) {
	logger, err := zapdriver.NewProduction()
	if environment == "local" {
		logger, err = zap.NewDevelopment()
	}
	if environment == "test" {
		testConfig := zap.NewDevelopmentConfig()
		testConfig.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
		logger, err = testConfig.Build()
	}
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	GetLogger().Info("Logger configured")
}
