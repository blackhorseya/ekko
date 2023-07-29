//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	taskM "github.com/blackhorseya/ekko/entity/domain/task/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

type ListTicketsCondition struct {
	Page int
	Size int
}

// IBiz declare task domain interface
type IBiz interface {
	// GetTicketByID serve caller to given ticket's id to get a ticket
	GetTicketByID(ctx contextx.Contextx, id string) (ticket *taskM.Ticket, err error)

	// ListTickets serve caller to list all tickets
	ListTickets(ctx contextx.Contextx, condition ListTicketsCondition) (tickets []*taskM.Ticket, total int, err error)

	// CreateTicket serve caller to create a ticket
	CreateTicket(ctx contextx.Contextx, title string) (ticket *taskM.Ticket, err error)

	// UpdateTicketStatus serve caller to update the ticket's status by id
	UpdateTicketStatus(ctx contextx.Contextx, id string, status taskM.TicketStatus) (ticket *taskM.Ticket, err error)

	// DeleteTicket serve caller to given ticket's id to delete the ticket
	DeleteTicket(ctx contextx.Contextx, id string) error
}
