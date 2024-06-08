package contextx

import (
	"github.com/gin-gonic/gin"
)

// AddContextxMiddleware is used to add contextx middleware.
func AddContextxMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(KeyCtx, WithContext(c.Request.Context()))

		c.Next()
	}
}
