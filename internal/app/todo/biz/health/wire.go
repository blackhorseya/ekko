// +build wireinject

package health

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health/repo"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	NewImpl,
)

// CreateHealthBiz serve user to create health biz
func CreateHealthBiz(repo repo.IRepo) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
