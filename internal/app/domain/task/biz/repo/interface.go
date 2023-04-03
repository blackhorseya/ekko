//go:generate mockgen -destination=./mock_${GOFILE} -package=repo -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/ekko/pkg/contextx"
	tm "github.com/blackhorseya/ekko/pkg/entity/domain/task/model"
	"github.com/google/wire"
)

// QueryTicketsCondition declare list tasks condition
type QueryTicketsCondition struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// IRepo declare task repo interface
type IRepo interface {
	// GetByID serve caller to get a task by id
	GetByID(ctx contextx.Contextx, id int64) (info *tm.Ticket, err error)

	// List serve caller to list all tasks
	List(ctx contextx.Contextx, condition QueryTicketsCondition) (info []*tm.Ticket, err error)

	// Create serve caller to create a task with title
	Create(ctx contextx.Contextx, created *tm.Ticket) (info *tm.Ticket, err error)

	// Count serve caller to count all tasks
	Count(ctx contextx.Contextx, condition QueryTicketsCondition) (total int, err error)

	// Update serve caller to update a task
	Update(ctx contextx.Contextx, updated *tm.Ticket) (info *tm.Ticket, err error)

	// DeleteByID serve caller to remove a task by id
	DeleteByID(ctx contextx.Contextx, id int64) error
}

// ProvideMariadb is a provider set for mariadb implementation
var ProvideMariadb = wire.NewSet(NewMariadb)
