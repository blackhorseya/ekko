package middlewares

import (
	trace2 "github.com/blackhorseya/todo-app/internal/pkg/base/trace"
	"github.com/blackhorseya/todo-app/internal/pkg/ctx"
	"github.com/gin-gonic/gin"
)

// ContextMiddleware added context into gin
func ContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		context := ctx.WithValue(ctx.Background(), "traceID", trace2.NewTraceID())
		c.Set("ctx", context)

		// process request
		c.Next()
	}
}
