package response

import (
	"net/http"

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

			c.AbortWithStatusJSON(http.StatusInternalServerError, Err.WithMessage(err.Error()))
		}()

		c.Next()
	}
}
