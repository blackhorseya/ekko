package er

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddErrorHandlingMiddleware global handle *gin.Context error middleware
func AddErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if c.Errors.Last() == nil {
				return
			}

			err := c.Errors.Last()

			switch e := err.Err.(type) {
			case *Error:
				c.AbortWithStatusJSON(e.Status, e)
				break
			default:
				appError := New(http.StatusInternalServerError, 50099, e.Error(), e.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, appError)
				break
			}
		}()

		c.Next()
	}
}
