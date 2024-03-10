package model

import (
	"time"

	"github.com/google/uuid"
)

// Ticket is an entity that represents a ticket.
type Ticket struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Completed bool      `json:"completed,omitempty"`
	OwnerID   uuid.UUID `json:"owner_id,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
