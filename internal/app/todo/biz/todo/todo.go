package todo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/ticket"
	"github.com/blackhorseya/todo-app/pb"
	"github.com/google/wire"
)

// ITodoBiz describe task business service function
//
//go:generate mockery --all --inpackage
type ITodoBiz interface {
	// GetByID serve caller to given task's id to get a task
	GetByID(ctx contextx.Contextx, id int64) (task *ticket.Task, err error)

	// List serve caller to list all tasks
	List(ctx contextx.Contextx, page, size int) (tasks []*ticket.Task, total int, err error)

	// Create serve caller to create a task
	Create(ctx contextx.Contextx, title string) (task *ticket.Task, err error)

	// UpdateStatus serve caller to update the task's status by id
	UpdateStatus(ctx contextx.Contextx, id int64, status pb.TaskStatus) (task *ticket.Task, err error)

	// ChangeTitle serve caller to change the task's title by id
	ChangeTitle(ctx contextx.Contextx, id int64, title string) (task *ticket.Task, err error)

	// Delete serve caller to given task's id to delete the task
	Delete(ctx contextx.Contextx, id int64) error
}

var (
	// ProviderSet is a provider set for wire
	ProviderSet = wire.NewSet(NewImpl, repo.ProviderSetMariadb)

	// ProviderSetViaHTTP is a http provider set for wire
	ProviderSetViaHTTP = wire.NewSet(NewImpl, repo.ProviderSetViaHTTP)
)
