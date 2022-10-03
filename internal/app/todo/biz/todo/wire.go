//go:build wireinject

package todo

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateIBiz serve user to create health biz
func CreateIBiz(repo repo.IRepo) (ITodoBiz, error) {
	panic(wire.Build(testProviderSet))
}
