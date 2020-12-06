package task

import "github.com/blackhorseya/todo-app/internal/app/entities"

// Biz describe task business service function
type Biz interface {
	Create(t *entities.Task) (task *entities.Task, err error)
	RemoveByID(id string) (ok bool, err error)
	ModifyTitle(id, title string) (task *entities.Task, err error)
	List() (tasks []*entities.Task, err error)
}
