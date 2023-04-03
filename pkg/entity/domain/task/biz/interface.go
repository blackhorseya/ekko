//go:generate mockgen -destination=./mock_${GOFILE} -package=biz -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/ekko/pkg/contextx"
	tm "github.com/blackhorseya/ekko/pkg/entity/domain/task/model"
)

// IBiz declare task domain interface
type IBiz interface {
	// GetByID serve caller to given task's id to get a task
	GetByID(ctx contextx.Contextx, id int64) (info *tm.Ticket, err error)

	// List serve caller to list all tasks
	List(ctx contextx.Contextx, page, size int) (info []*tm.Ticket, total int, err error)

	// Create serve caller to create a task
	Create(ctx contextx.Contextx, title string) (info *tm.Ticket, err error)

	// UpdateStatus serve caller to update the task's status by id
	UpdateStatus(ctx contextx.Contextx, id int64, status tm.TicketStatus) (info *tm.Ticket, err error)

	// Delete serve caller to given task's id to delete the task
	Delete(ctx contextx.Contextx, id int64) error
}
