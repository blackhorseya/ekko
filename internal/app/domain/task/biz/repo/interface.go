package repo

import (
	"github.com/blackhorseya/todo-app/pkg/contextx"
	tm "github.com/blackhorseya/todo-app/pkg/entity/domain/task/model"
	"github.com/google/wire"
)

// QueryTasksCondition declare list tasks condition
type QueryTasksCondition struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// IRepo declare task repo interface
//
//go:generate mockery --all --inpackage
type IRepo interface {
	// GetByID serve caller to get a task by id
	GetByID(ctx contextx.Contextx, id int64) (info *tm.Task, err error)

	// List serve caller to list all tasks
	List(ctx contextx.Contextx, condition QueryTasksCondition) (info []*tm.Task, err error)

	// Create serve caller to create a task with title
	Create(ctx contextx.Contextx, created *tm.Task) (info *tm.Task, err error)

	// Count serve caller to count all tasks
	Count(ctx contextx.Contextx, condition QueryTasksCondition) (total int, err error)

	// Update serve caller to update a task
	Update(ctx contextx.Contextx, updated *tm.Task) (info *tm.Task, err error)

	// DeleteByID serve caller to remove a task by id
	DeleteByID(ctx contextx.Contextx, id int64) error
}

// ProviderMariadbSet is a provider set for mariadb implementation
var ProviderMariadbSet = wire.NewSet(NewMariadb)
