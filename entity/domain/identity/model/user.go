package model

// User is an entity that represents a user
type User struct {
	ID     string `json:"id,omitempty"`
	Active bool   `json:"active,omitempty"`

	Profile *Profile `json:"profile,omitempty"`
}
