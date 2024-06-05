package model

import (
	"encoding/json"
	"fmt"

	"github.com/blackhorseya/ekko/pkg/contextx"
)

// TicketStatus is an interface that represents the status of a ticket.
type TicketStatus interface {
	fmt.Stringer
	json.Marshaler

	Execute(ctx contextx.Contextx, ticket *Ticket) error
}

var _ TicketStatus = &TicketStatusBacklog{}

// TicketStatusBacklog is a type that represents the backlog status of a ticket.
type TicketStatusBacklog struct {
}

func (s *TicketStatusBacklog) String() string {
	return "backlog"
}

func (s *TicketStatusBacklog) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *TicketStatusBacklog) Execute(ctx contextx.Contextx, ticket *Ticket) error {
	// todo: 2024/6/6|sean|implement me

	ticket.Status = &TicketStatusTodo{}
	return nil
}

var _ TicketStatus = &TicketStatusTodo{}

// TicketStatusTodo is a type that represents the todo status of a ticket.
type TicketStatusTodo struct {
}

func (s *TicketStatusTodo) String() string {
	return "todo"
}

func (s *TicketStatusTodo) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *TicketStatusTodo) Execute(ctx contextx.Contextx, ticket *Ticket) error {
	// todo: 2024/6/6|sean|implement me

	ticket.Status = &TicketStatusInProgress{}
	return nil
}

var _ TicketStatus = &TicketStatusInProgress{}

// TicketStatusInProgress is a type that represents the in progress status of a ticket.
type TicketStatusInProgress struct {
}

func (s *TicketStatusInProgress) String() string {
	return "in_progress"
}

func (s *TicketStatusInProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *TicketStatusInProgress) Execute(ctx contextx.Contextx, ticket *Ticket) error {
	// todo: 2024/6/6|sean|implement me

	ticket.Status = &TicketStatusDone{}
	return nil
}

var _ TicketStatus = &TicketStatusDone{}

// TicketStatusDone is a type that represents the done status of a ticket.
type TicketStatusDone struct {
}

func (s *TicketStatusDone) String() string {
	return "done"
}

func (s *TicketStatusDone) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *TicketStatusDone) Execute(ctx contextx.Contextx, ticket *Ticket) error {
	// todo: 2024/6/6|sean|implement me
	return nil
}
