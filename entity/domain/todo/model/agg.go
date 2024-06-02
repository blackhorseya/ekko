package model

import (
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
func NewTodo(title string) *Todo {
	return &Todo{
		Title:     title,
		UpdatedAt: time.Now(),
	}
}
