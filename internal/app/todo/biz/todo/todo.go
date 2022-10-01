package todo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/blackhorseya/todo-app/pb"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IBiz describe task business service function
//
//go:generate mockery --all --inpackage
type IBiz interface {
	// GetByID serve caller to given task's id to get a task
	GetByID(ctx contextx.Contextx, id primitive.ObjectID) (task *todo.Task, err error)

	// List serve caller to list all tasks
	List(ctx contextx.Contextx, start, end int) (tasks []*todo.Task, total int, err error)

	// Create serve caller to create a task
	Create(ctx contextx.Contextx, title string) (task *todo.Task, err error)

	// UpdateStatus serve caller to update the task's status by id
	UpdateStatus(ctx contextx.Contextx, id primitive.ObjectID, status pb.TaskStatus) (task *todo.Task, err error)

	// ChangeTitle serve caller to change the task's title by id
	ChangeTitle(ctx contextx.Contextx, id primitive.ObjectID, title string) (task *todo.Task, err error)

	// Delete serve caller to given task's id to delete the task
	Delete(ctx contextx.Contextx, id primitive.ObjectID) error
}

var (
	// ProviderSet is a provider set for wire
	ProviderSet = wire.NewSet(NewImpl, repo.ProviderSetViaDatabase)

	// ProviderSetViaHTTP is a http provider set for wire
	ProviderSetViaHTTP = wire.NewSet(NewImpl, repo.ProviderSetViaHTTP)
)
