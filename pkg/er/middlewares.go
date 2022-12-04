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
			c.Errors = c.Errors[:0]

			switch err.Err.(type) {
			case *Error:
				appError := err.Err.(*Error)
				c.AbortWithStatusJSON(appError.Status, appError)
				break
			default:
				appError := New(http.StatusInternalServerError, 50099, err.Err.Error(), err.Err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, appError)
				break
			}
		}()

		c.Next()
	}
}
