//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/ekko/entity/domain/task/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

// ListTicketOptions is the options for list tickets.
type ListTicketOptions struct {
	Page int
	Size int
}

// ITaskBiz is the interface for task business logic.
type ITaskBiz interface {
	// CreateTicket is used to create a ticket.
	CreateTicket(ctx contextx.Contextx, title string) (item *model.Ticket, err error)

	// GetTicketByID is used to get a ticket by ID.
	GetTicketByID(ctx contextx.Contextx, id string) (item *model.Ticket, err error)

	// ListTicket is used to list tickets.
	ListTicket(ctx contextx.Contextx, options ListTicketOptions) (items []*model.Ticket, total int, err error)

	// UpdateTicket is used to update a ticket.
	UpdateTicket(ctx contextx.Contextx, id string, update *model.Ticket) (err error)
}
