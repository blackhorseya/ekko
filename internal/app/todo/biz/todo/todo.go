package todo

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo"
	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/pb"
	"github.com/google/wire"
)

// IBiz describe todo business service function
type IBiz interface {
	// GetByID serve caller to given task's id to get a task
	GetByID(ctx contextx.Contextx, id int64) (task *pb.Task, err error)

	// List serve caller to list all tasks
	List(ctx contextx.Contextx, start, end int) (tasks []*pb.Task, total int, err error)

	// Create serve caller to create a task
	Create(ctx contextx.Contextx, title string) (task *pb.Task, err error)

	// UpdateStatus serve caller to update the task's status by id
	UpdateStatus(ctx contextx.Contextx, id int64, status bool) (task *pb.Task, err error)

	// ChangeTitle serve caller to change the task's title by id
	ChangeTitle(ctx contextx.Contextx, id int64, title string) (task *pb.Task, err error)

	// Delete serve caller to given task's id to delete the task
	Delete(ctx contextx.Contextx, id int64) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
