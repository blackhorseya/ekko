package task

import (
	"github.com/blackhorseya/todo-app/internal/app/entities"
)

type impl struct {
}

// NewImpl is a constructor of implement business
func NewImpl() Biz {
	return &impl{}
}

// Create a task
func (i *impl) Create(t *entities.Task) (task *entities.Task, err error) {
	panic("implement me")
}

// RemoveByID a task by id
func (i *impl) RemoveByID(id string) (ok bool, err error) {
	panic("implement me")
}

// ModifyTitle a title of task
func (i *impl) ModifyTitle(id, title string) (task *entities.Task, err error) {
	panic("implement me")
}

// List all tasks
func (i *impl) List() (tasks []*entities.Task, err error) {
	panic("implement me")
}
