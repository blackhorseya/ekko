package repository

import (
	"github.com/blackhorseya/todo-app/internal/app/entities"
)

type impl struct {
}

// NewImpl is a constructor task of implement repository
func NewImpl() TaskRepo {
	return &impl{}
}

// QueryTaskList handle query task list by limit and offset
func (i *impl) QueryTaskList(limit, offset int32) (tasks []*entities.Task, err error) {
	panic("implement me")
}

// CreateTask handle create a task
func (i *impl) CreateTask(newTask *entities.Task) (task *entities.Task, err error) {
	panic("implement me")
}
