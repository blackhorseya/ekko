package task

import (
	"github.com/blackhorseya/todo-app/internal/app/biz/task/repository"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	NewImpl,
)

func CreateTaskBiz(repo repository.TaskRepo) (Biz, error) {
	panic(wire.Build(testProviderSet))
}
