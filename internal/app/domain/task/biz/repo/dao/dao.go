package dao

import (
	"time"

	taskM "github.com/blackhorseya/ekko/entity/domain/task/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Ticket is the ticket model
type Ticket struct {
	ID        string             `json:"id" db:"id"`
	Title     string             `json:"title" db:"title"`
	Status    taskM.TicketStatus `json:"status" db:"status"`
	CreatedAt time.Time          `json:"created_at" db:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" db:"updated_at"`
}

// NewTicket is used to create a new ticket model
func NewTicket(t *taskM.Ticket) *Ticket {
	return &Ticket{
		ID:        t.Id,
		Title:     t.Title,
		Status:    t.Status,
		CreatedAt: t.CreatedAt.AsTime(),
		UpdatedAt: t.UpdatedAt.AsTime(),
	}
}

// ToEntity is used to convert the ticket model to ticket entity
func (t *Ticket) ToEntity() *taskM.Ticket {
	return &taskM.Ticket{
		Id:        t.ID,
		Title:     t.Title,
		Status:    t.Status,
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}

// Tickets is the slice of ticket model
type Tickets []*Ticket

// ToEntity is used to convert the slice of ticket model to slice of ticket entity
func (slice Tickets) ToEntity() []*taskM.Ticket {
	tickets := make([]*taskM.Ticket, 0, len(slice))
	for _, ticket := range slice {
		tickets = append(tickets, ticket.ToEntity())
	}

	return tickets
}
