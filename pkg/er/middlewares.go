package er

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleError global handle *gin.Context error middleware
func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if c.Errors.Last() == nil {
				return
			}

			err := c.Errors.Last()

			switch e := err.Err.(type) {
			case *Error:
				c.AbortWithStatusJSON(e.Code, e)
				break
			default:
				appError := New(http.StatusInternalServerError, e.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, appError)
				break
			}
		}()

		c.Next()
	}
}
