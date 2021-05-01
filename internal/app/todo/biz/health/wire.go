// +build wireinject

package health

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health/repo"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(
	NewImpl,
)

// CreateHealthBiz serve user to create health biz
func CreateHealthBiz(logger *zap.Logger, repo repo.IRepo) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
