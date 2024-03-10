package model

import (
	"github.com/google/uuid"
)

// Ticket is an entity that represents a ticket.
type Ticket struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Completed bool      `json:"completed,omitempty"`
	OwnerID   uuid.UUID `json:"owner_id,omitempty"`
}
