package responsex

import (
	"github.com/gin-gonic/gin"
)

// AddErrorHandlingMiddleware is used to add error handling middleware.
func AddErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if c.Errors.Last() == nil {
				return
			}

			err := c.Errors.Last().Err

			Err(c, err)
			c.Abort()
		}()

		c.Next()
	}
}
