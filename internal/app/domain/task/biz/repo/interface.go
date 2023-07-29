//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	taskM "github.com/blackhorseya/ekko/entity/domain/task/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

// ListTicketsCondition declare list tasks condition
type ListTicketsCondition struct {
	Limit  int
	Offset int
}

// IRepo declare task repo interface
type IRepo interface {
	// GetTicketByID serve caller to get a ticket by id
	GetTicketByID(ctx contextx.Contextx, id string) (ticket *taskM.Ticket, err error)

	// ListTickets serve caller to list all tickets
	ListTickets(ctx contextx.Contextx, condition ListTicketsCondition) (tickets []*taskM.Ticket, total int, err error)

	// CountTickets serve caller to count all tickets
	CountTickets(ctx contextx.Contextx, condition ListTicketsCondition) (total int, err error)

	// CreateTicket serve caller to create a ticket
	CreateTicket(ctx contextx.Contextx, created *taskM.Ticket) (ticket *taskM.Ticket, err error)

	// UpdateTicket serve caller to update a ticket
	UpdateTicket(ctx contextx.Contextx, updated *taskM.Ticket) error

	// DeleteTicketByID serve caller to remove a ticket by id
	DeleteTicketByID(ctx contextx.Contextx, id string) error
}
