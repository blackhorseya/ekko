package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Ticket is an aggregate root that represents a ticket in the system.
type Ticket struct {
	ID          string       `json:"id,omitempty" bson:"_id"`
	Title       string       `json:"title,omitempty" bson:"title"`
	Description string       `json:"description,omitempty" bson:"description"`
	Status      TicketStatus `json:"status,omitempty" bson:"status"`
	Priority    int          `json:"priority,omitempty" bson:"priority"`
	CreateBy    string       `json:"created_by,omitempty" bson:"create_by"`
	CreatedAt   time.Time    `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at,omitempty" bson:"updated_at"`
}

// NewTicket creates a new ticket with the given title.
func NewTicket(title string) *Ticket {
	return &Ticket{
		ID:          "",
		Title:       title,
		Description: "",
		Status:      &TicketStatusBacklog{},
		Priority:    0,
		CreateBy:    "",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func (x *Ticket) UnmarshalBSON(bytes []byte) error {
	type Alias Ticket
	alias := &struct {
		*Alias `bson:",inline"`
		Status string `bson:"status"`
	}{
		Alias:  (*Alias)(x),
		Status: "",
	}

	if err := bson.Unmarshal(bytes, alias); err != nil {
		return err
	}

	status, err := UnmarshalTicketStatus(alias.Status)
	if err != nil {
		return err
	}
	x.Status = status

	return nil
}

func (x *Ticket) MarshalBSON() ([]byte, error) {
	type Alias Ticket
	alias := &struct {
		*Alias `bson:",inline"`
		Status string `bson:"status"`
	}{
		Alias:  (*Alias)(x),
		Status: "",
	}

	if x.Status != nil {
		alias.Status = x.Status.String()
	}

	return bson.Marshal(alias)
}
