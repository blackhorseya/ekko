package model

import (
	"github.com/google/uuid"
)

// User is an entity that represents a user
type User struct {
	ID     uuid.UUID `json:"id,omitempty"`
	Active bool      `json:"active,omitempty"`

	Profile Profile `json:"profile,omitempty"`
}
