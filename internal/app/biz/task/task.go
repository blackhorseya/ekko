package task

import (
	"github.com/blackhorseya/todo-app/internal/app/entities"
	"github.com/google/wire"
)

// Biz describe task business service function
type Biz interface {
	Create(t *entities.Task) (task *entities.Task, err error)
	RemoveByID(id string) (ok bool, err error)
	Complete(id string) (task *entities.Task, err error)
	Incomplete(id string) (task *entities.Task, err error)
	ModifyTitle(id, title string) (task *entities.Task, err error)
	List() (tasks []*entities.Task, err error)
}

// ProviderSet is a task provider set
var ProviderSet = wire.NewSet(NewImpl)
