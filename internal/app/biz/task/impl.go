package task

import (
	"fmt"

	"github.com/blackhorseya/todo-app/internal/app/biz/task/repository"
	"github.com/blackhorseya/todo-app/internal/app/entities"
)

type impl struct {
	TaskRepo repository.TaskRepo
}

// NewImpl is a constructor of implement business
func NewImpl(repo repository.TaskRepo) Biz {
	return &impl{
		TaskRepo: repo,
	}
}

// Create serve user to create a task
func (i *impl) Create(newTask *entities.Task) (task *entities.Task, err error) {
	if len(newTask.Title) == 0 {
		return nil, fmt.Errorf("title must be NOT empty")
	}

	task, err = i.TaskRepo.CreateTask(newTask)
	if err != nil {
		return nil, err
	}

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
func (i *impl) List(page, size int32) (tasks []*entities.Task, err error) {
	tasks, err = i.TaskRepo.QueryTaskList(size, (page-1)*size)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
