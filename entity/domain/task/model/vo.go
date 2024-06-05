package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/blackhorseya/ekko/pkg/contextx"
)

// TicketStatus is an interface that represents the status of a ticket.
type TicketStatus interface {
	fmt.Stringer
	json.Marshaler

	Execute(ctx contextx.Contextx, ticket *Ticket) error
}

// UnmarshalTicketStatus unmarshals a ticket status from a string.
func UnmarshalTicketStatus(name string) (TicketStatus, error) {
	switch name {
	case (&TicketStatusBacklog{}).String():
		return &TicketStatusBacklog{}, nil
	case (&TicketStatusTodo{}).String():
		return &TicketStatusTodo{}, nil
	case (&TicketStatusInProgress{}).String():
		return &TicketStatusInProgress{}, nil
	case (&TicketStatusDone{}).String():
		return &TicketStatusDone{}, nil
	case "":
		return &TicketStatusUnknown{}, nil
	default:
		return nil, fmt.Errorf("invalid status: %s", name)
	}
}

var _ TicketStatus = &TicketStatusUnknown{}

// TicketStatusUnknown is a type that represents an unknown status of a ticket.
type TicketStatusUnknown struct {
}

func (s *TicketStatusUnknown) String() string {
	return "unknown"
}

func (s *TicketStatusUnknown) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *TicketStatusUnknown) Execute(ctx contextx.Contextx, ticket *Ticket) error {
	return errors.New("unknown status")
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
	return errors.New("ticket is already done")
}
