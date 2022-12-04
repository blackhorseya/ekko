package contextx

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AddContextxWitLoggerMiddleware add custom contextx middleware
func AddContextxWitLoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(string(KeyCtx), BackgroundWithLogger(logger))

		c.Next()
	}
}
