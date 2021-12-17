package repo

import (
	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/pb"
	"github.com/google/wire"
)

// IRepo declare repository service function
type IRepo interface {
	// GetByID serve caller to get a task by id
	GetByID(ctx contextx.Contextx, id int64) (task *pb.Task, err error)

	// List serve caller to list all tasks
	List(ctx contextx.Contextx, limit, offset int) (tasks []*pb.Task, err error)

	// Create serve caller to create a task with title
	Create(ctx contextx.Contextx, newTask *pb.Task) (task *pb.Task, err error)

	// Count serve caller to count all tasks
	Count(ctx contextx.Contextx) (total int, err error)

	// Update serve caller to update a task
	Update(ctx contextx.Contextx, updated *pb.Task) (task *pb.Task, err error)

	// Remove serve caller to remove a task by id
	Remove(ctx contextx.Contextx, id int64) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
