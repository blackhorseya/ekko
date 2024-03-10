package model

import (
	"time"
)

// Ticket is an entity that represents a ticket.
type Ticket struct {
	ID        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
	OwnerID   string `json:"owner_id,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
