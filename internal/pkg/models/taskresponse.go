package models

import (
	"github.com/blackhorseya/todo-app/pb"
)

// TaskResponse declare presentation layer task response struct
type TaskResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	CreatedAt string `json:"created_at"`
}

// NewTaskResponse serve caller to create a task response from pb.Task
func NewTaskResponse(t *pb.Task) *TaskResponse {
	return &TaskResponse{
		ID:        t.Id,
		Title:     t.Title,
		Completed: t.Completed,
		CreatedAt: t.CreateAt.AsTime().UTC().Format(rfc3339Mill),
	}
}
