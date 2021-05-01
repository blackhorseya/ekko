// +build wireinject

package health

import (
	health2 "github.com/blackhorseya/todo-app/internal/app/todo/biz/health"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateHealthHandler serve user to create health api handler
func CreateHealthHandler(biz health2.IBiz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
