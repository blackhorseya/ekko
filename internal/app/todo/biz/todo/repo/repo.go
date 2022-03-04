package repo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IRepo declare repository service function
//go:generate mockery --name=IRepo
type IRepo interface {
	// GetByID serve caller to get a task by id
	GetByID(ctx contextx.Contextx, id primitive.ObjectID) (task *todo.Task, err error)

	// List serve caller to list all tasks
	List(ctx contextx.Contextx, limit, offset int) (tasks []*todo.Task, err error)

	// Create serve caller to create a task with title
	Create(ctx contextx.Contextx, newTask *todo.Task) (task *todo.Task, err error)

	// Count serve caller to count all tasks
	Count(ctx contextx.Contextx) (total int, err error)

	// Update serve caller to update a task
	Update(ctx contextx.Contextx, updated *todo.Task) (task *todo.Task, err error)

	// Remove serve caller to remove a task by id
	Remove(ctx contextx.Contextx, id primitive.ObjectID) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
