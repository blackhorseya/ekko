// +build wireinject

package todo

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo"
	"github.com/bwmarrin/snowflake"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIBiz serve user to create health biz
func CreateIBiz(logger *zap.Logger, repo repo.IRepo, generator *snowflake.Node) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
