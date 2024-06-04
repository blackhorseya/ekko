package model

import (
	"errors"

	"github.com/blackhorseya/ekko/pkg/contextx"
)

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
		return nil, errors.New("user not found in context")
	}

	return user, nil
}
