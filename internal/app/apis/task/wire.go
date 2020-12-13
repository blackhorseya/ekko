// +build wireinject

package task

import (
	"github.com/blackhorseya/todo-app/internal/app/biz/task"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateTaskHandler serve user to create task api handler
func CreateTaskHandler(taskBiz task.Biz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
