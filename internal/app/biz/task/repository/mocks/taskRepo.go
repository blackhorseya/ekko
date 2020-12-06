package mocks

import (
	"github.com/blackhorseya/todo-app/internal/app/entities"
	"github.com/stretchr/testify/mock"
)

type TaskRepo struct {
	mock.Mock
}

func (m *TaskRepo) QueryTaskList(limit, offset int32) (tasks []*entities.Task, err error) {
	panic("implement me")
}

func (m *TaskRepo) CreateTask(newTask *entities.Task) (task *entities.Task, err error) {
	ret := m.Called(newTask)

	var r0 *entities.Task
	if rf, ok := ret.Get(0).(func(task2 *entities.Task) *entities.Task); ok {
		r0 = rf(newTask)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(task2 *entities.Task) error); ok {
		r1 = rf(newTask)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
