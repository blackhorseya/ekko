package models

import (
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/blackhorseya/todo-app/pb"
)

// TaskResponse declare presentation layer task response struct
type TaskResponse struct {
	ID        string        `json:"id"`
	Title     string        `json:"title"`
	Status    pb.TaskStatus `json:"status"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt string        `json:"updated_at"`
}

// NewTaskResponse serve caller to create a task response from pb.Task
func NewTaskResponse(t *todo.Task) *TaskResponse {
	return &TaskResponse{
		ID:        t.ID.Hex(),
		Title:     t.Title,
		Status:    t.Status,
		CreatedAt: t.CreatedAt.Time().UTC().Format(rfc3339Mill),
		UpdatedAt: t.UpdatedAt.Time().UTC().Format(rfc3339Mill),
	}
}
