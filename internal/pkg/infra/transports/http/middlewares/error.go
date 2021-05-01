package middlewares

import (
	"net/http"

	"github.com/blackhorseya/todo-app/internal/pkg/entity/er"
	"github.com/gin-gonic/gin"
)

// ErrorMiddleware serve caller to format api response
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if c.Errors.Last() == nil {
				return
			}

			err := c.Errors.Last()
			c.Errors = c.Errors[:0]

			switch err.Err.(type) {
			case *er.APPError:
				appError := err.Err.(*er.APPError)
				c.AbortWithStatusJSON(appError.Status, appError)
				break
			default:
				c.AbortWithStatus(http.StatusInternalServerError)
				break
			}
		}()

		c.Next()

	}
}
