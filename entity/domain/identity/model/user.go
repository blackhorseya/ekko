package model

import (
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/google/uuid"
)

var defaultUser = &User{
	ID:     uuid.New().String(),
	Active: true,
	Profile: Profile{
		Name: "Anonymous",
	},
}

// User is an entity that represents a user
type User struct {
	ID     string `json:"id,omitempty"`
	Active bool   `json:"active,omitempty"`

	Profile Profile `json:"profile,omitempty"`
}

// FromContext is a factory function that creates a User from a context.
func FromContext(ctx contextx.Contextx) (*User, error) {
	user, ok := ctx.Value(contextx.KeyWho).(*User)
	if !ok {
		return defaultUser, nil
	}

	return user, nil
}
