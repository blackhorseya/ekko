package biz

import (
	"github.com/blackhorseya/todo-app/pkg/contextx"
	tm "github.com/blackhorseya/todo-app/pkg/entity/domain/task/model"
)

// IBiz declare task domain interface
//
//go:generate mockery --all --inpackage
type IBiz interface {
	Liveness(ctx contextx.Contextx) error

	Readiness(ctx contextx.Contextx) error

	// GetByID serve caller to given task's id to get a task
	GetByID(ctx contextx.Contextx, id int64) (info *tm.Task, err error)

	// List serve caller to list all tasks
	List(ctx contextx.Contextx, page, size int) (info []*tm.Task, total int, err error)

	// Create serve caller to create a task
	Create(ctx contextx.Contextx, title string) (info *tm.Task, err error)

	// UpdateStatus serve caller to update the task's status by id
	UpdateStatus(ctx contextx.Contextx, id int64, status tm.TaskStatus) (info *tm.Task, err error)

	// ChangeTitle serve caller to change the task's title by id
	ChangeTitle(ctx contextx.Contextx, id int64, title string) (info *tm.Task, err error)

	// Delete serve caller to given task's id to delete the task
	Delete(ctx contextx.Contextx, id int64) error
}
