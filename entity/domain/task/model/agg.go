package model

import (
	"time"
)

// Ticket is an aggregate root that represents a ticket in the system.
type Ticket struct {
	ID          string       `json:"id,omitempty"`
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	Status      TicketStatus `json:"status,omitempty"`
	Priority    int          `json:"priority,omitempty"`
	CreateBy    string       `json:"created_by,omitempty"`
	CreatedAt   time.Time    `json:"created_at,omitempty"`
	UpdatedAt   time.Time    `json:"updated_at,omitempty"`
}
