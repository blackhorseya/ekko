// +build wireinject

package health

import (
	repository2 "github.com/blackhorseya/todo-app/internal/app/todo/biz/health/repository"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	NewImpl,
)

// CreateHealthBiz serve user to create health biz
func CreateHealthBiz(repo repository2.HealthRepo) (Biz, error) {
	panic(wire.Build(testProviderSet))
}
