package task

import (
	"fmt"

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
	if len(t.Title) == 0 {
		return nil, fmt.Errorf("title must be NOT empty")
	}

	// todo: 2020-12-06|12:34|doggy|implement it

	return task, nil
}

// Complete a task which change complete status of task to true
func (i *impl) Complete(id string) (task *entities.Task, err error) {
	panic("implement me")
}

// Incomplete a task which change complete status of task to false
func (i *impl) Incomplete(id string) (task *entities.Task, err error) {
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
