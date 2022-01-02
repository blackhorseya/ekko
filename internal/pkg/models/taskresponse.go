package models

import (
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
)

// TaskResponse declare presentation layer task response struct
type TaskResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// NewTaskResponse serve caller to create a task response from pb.Task
func NewTaskResponse(t *todo.Task) *TaskResponse {
	return &TaskResponse{
		ID:        t.ID.Hex(),
		Title:     t.Title,
		Completed: t.Completed,
		CreatedAt: t.CreatedAt.Time().UTC().Format(rfc3339Mill),
		UpdatedAt: t.UpdatedAt.Time().UTC().Format(rfc3339Mill),
	}
}
