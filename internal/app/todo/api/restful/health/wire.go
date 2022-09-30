//go:build wireinject

package health

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve user to create health api handler
func CreateIHandler(e *gin.Engine, biz health.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
