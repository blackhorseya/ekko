//go:build wireinject

package health

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health/repo"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	NewImpl,
)

// CreateIBiz serve user to create health biz
func CreateIBiz(repo repo.IHealthRepo) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
