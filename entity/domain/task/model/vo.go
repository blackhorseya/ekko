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
	json.Unmarshaler

	Execute(ctx contextx.Contextx, ticket *Ticket) error
}

var _ TicketStatus = &TicketStatusBacklog{}

// TicketStatusBacklog is a type that represents the backlog status of a ticket.
type TicketStatusBacklog struct {
}

func (s *TicketStatusBacklog) String() string {
	// todo: 2024/6/6|sean|implement me
	panic("implement me")
}

func (s *TicketStatusBacklog) MarshalJSON() ([]byte, error) {
	// todo: 2024/6/6|sean|implement me
	panic("implement me")
}

func (s *TicketStatusBacklog) UnmarshalJSON(bytes []byte) error {
	// todo: 2024/6/6|sean|implement me
	panic("implement me")
}

func (s *TicketStatusBacklog) Execute(ctx contextx.Contextx, ticket *Ticket) error {
	// todo: 2024/6/6|sean|implement me
	panic("implement me")
}

var _ TicketStatus = &TicketStatusTodo{}

// TicketStatusTodo is a type that represents the todo status of a ticket.
type TicketStatusTodo struct {
}

func (s *TicketStatusTodo) String() string {
	// TODO implement me
	panic("implement me")
}

func (s *TicketStatusTodo) MarshalJSON() ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

func (s *TicketStatusTodo) UnmarshalJSON(bytes []byte) error {
	// TODO implement me
	panic("implement me")
}

func (s *TicketStatusTodo) Execute(ctx contextx.Contextx, ticket *Ticket) error {
	// TODO implement me
	panic("implement me")
}

var _ TicketStatus = &TicketStatusInProgress{}

// TicketStatusInProgress is a type that represents the in progress status of a ticket.
type TicketStatusInProgress struct {
}

func (s *TicketStatusInProgress) String() string {
	// TODO implement me
	panic("implement me")
}

func (s *TicketStatusInProgress) MarshalJSON() ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

func (s *TicketStatusInProgress) UnmarshalJSON(bytes []byte) error {
	// TODO implement me
	panic("implement me")
}

func (s *TicketStatusInProgress) Execute(ctx contextx.Contextx, ticket *Ticket) error {
	// TODO implement me
	panic("implement me")
}

var _ TicketStatus = &TicketStatusDone{}

// TicketStatusDone is a type that represents the done status of a ticket.
type TicketStatusDone struct {
}

func (s *TicketStatusDone) String() string {
	// TODO implement me
	panic("implement me")
}

func (s *TicketStatusDone) MarshalJSON() ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

func (s *TicketStatusDone) UnmarshalJSON(bytes []byte) error {
	// TODO implement me
	panic("implement me")
}

func (s *TicketStatusDone) Execute(ctx contextx.Contextx, ticket *Ticket) error {
	// TODO implement me
	panic("implement me")
}
