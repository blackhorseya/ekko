// +build wireinject

package health

import (
	"github.com/blackhorseya/todo-app/internal/app/biz/health"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateTaskHandler serve user to create task api handler
func CreateHealthHandler(biz health.Biz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
