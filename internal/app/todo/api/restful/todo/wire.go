//go:build wireinject

package todo

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve user to create health api handler
func CreateIHandler(e *gin.Engine, logger *zap.Logger, biz todo.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
