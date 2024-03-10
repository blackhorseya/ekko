package model

// Profile is a value object that represents a user's profile
type Profile struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
