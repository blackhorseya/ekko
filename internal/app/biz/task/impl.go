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

// Create serve user to create a task
func (i *impl) Create(newTask *entities.Task) (task *entities.Task, err error) {
	if len(newTask.Title) == 0 {
		return nil, fmt.Errorf("title must be NOT empty")
	}

	// todo: 2020-12-06|12:34|doggy|implement it

	return task, nil
}

// UpdateStatus serve user to update complete status of task by id
func (i *impl) UpdateStatus(id string, completed bool) (task *entities.Task, err error) {
	panic("implement me")
}

// Remove serve user to remove a task by id
func (i *impl) Remove(id string) (ok bool, err error) {
	panic("implement me")
}

// ChangeTitle serve user to update title of task
func (i *impl) ChangeTitle(id, newTitle string) (task *entities.Task, err error) {
	panic("implement me")
}

// List all tasks
func (i *impl) List() (tasks []*entities.Task, err error) {
	panic("implement me")
}
