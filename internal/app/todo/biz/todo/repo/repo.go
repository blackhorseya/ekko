package repo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/google/wire"
)

// QueryTodoCondition declare query tickets list condition
type QueryTodoCondition struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// ITodoRepo declare repository service function
//
//go:generate mockery --all --inpackage
type ITodoRepo interface {
	// GetByID serve caller to get a task by id
	GetByID(ctx contextx.Contextx, id uint64) (task *todo.Task, err error)

	// List serve caller to list all tasks
	List(ctx contextx.Contextx, condition QueryTodoCondition) (tasks []*todo.Task, err error)

	// Create serve caller to create a task with title
	Create(ctx contextx.Contextx, created *todo.Task) (task *todo.Task, err error)

	// Count serve caller to count all tasks
	Count(ctx contextx.Contextx) (total int, err error)

	// Update serve caller to update a task
	Update(ctx contextx.Contextx, updated *todo.Task) (task *todo.Task, err error)

	// Remove serve caller to remove a task by id
	Remove(ctx contextx.Contextx, id uint64) error
}

var (
	// ProviderSetMariadb is a provider set for wire
	ProviderSetMariadb = wire.NewSet(NewMariadb)

	// ProviderSetViaHTTP is a http provider set for wire
	ProviderSetViaHTTP = wire.NewSet(NewHTTP, NewOptions)
)
