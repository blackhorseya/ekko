package contextx

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// WithContextx add custom contextx middleware
func WithContextx(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(string(KeyCtx), WithLogger(logger))

		c.Next()
	}
}
