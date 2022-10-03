//go:build wireinject

package todo

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIHandler serve user to create health api handler
func CreateIHandler(e *gin.Engine, biz todo.ITodoBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
