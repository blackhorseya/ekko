//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/ekko/entity/domain/task/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

// ListCondition is the condition for listing tickets.
type ListCondition struct {
	CreatedBy string
	Limit     int
	Offset    int
}

// ITicketRepo is the interface that defines the methods of the ticket repository.
type ITicketRepo interface {
	Create(ctx contextx.Contextx, ticket *model.Ticket) (err error)
	GetByID(ctx contextx.Contextx, id string) (item *model.Ticket, err error)
	List(ctx contextx.Contextx, condition ListCondition) (items []*model.Ticket, total int, err error)
	Update(ctx contextx.Contextx, ticket *model.Ticket) (err error)
	Delete(ctx contextx.Contextx, id string) (err error)
}
