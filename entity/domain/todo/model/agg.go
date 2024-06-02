package model

import (
	"errors"
	"time"
)

// Todo is the aggregate root of the todo domain
type Todo struct {
	ID        string    `json:"id,omitempty" bson:"_id"`
	Title     string    `json:"title,omitempty" bson:"title"`
	Done      bool      `json:"done,omitempty" bson:"done"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

// NewTodo is to create a new todo
func NewTodo(title string) (*Todo, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	return &Todo{
		Title:     title,
		UpdatedAt: time.Now(),
	}, nil
}
