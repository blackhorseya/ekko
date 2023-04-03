//go:generate mockgen -destination=./mock_${GOFILE} -package=biz -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/ekko/pkg/contextx"
	im "github.com/blackhorseya/ekko/pkg/entity/domain/issue/model"
)

// IBiz declare issue domain interface
type IBiz interface {
	// GetByID serve caller to given issue's id to get a issue
	GetByID(ctx contextx.Contextx, id int64) (info *im.Ticket, err error)

	// List serve caller to list all tasks
	List(ctx contextx.Contextx, page, size int) (info []*im.Ticket, total int, err error)

	// Create serve caller to create a issue
	Create(ctx contextx.Contextx, title string) (info *im.Ticket, err error)

	// UpdateStatus serve caller to update the issue's status by id
	UpdateStatus(ctx contextx.Contextx, id int64, status im.TicketStatus) (info *im.Ticket, err error)

	// Delete serve caller to given issue's id to delete the issue
	Delete(ctx contextx.Contextx, id int64) error
}
