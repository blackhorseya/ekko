// +build wireinject

package health

import (
	"github.com/blackhorseya/todo-app/internal/app/biz/health/repository"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	NewImpl,
)

// CreateHealthBiz serve user to create health biz
func CreateHealthBiz(repo repository.HealthRepo) (Biz, error) {
	panic(wire.Build(testProviderSet))
}
