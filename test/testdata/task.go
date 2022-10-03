package testdata

import (
	"time"

	"github.com/blackhorseya/todo-app/internal/pkg/entity/ticket"
	"github.com/blackhorseya/todo-app/pb"
)

var (
	// Task1 task 1
	Task1 = &ticket.Task{
		ID:        1,
		Title:     "task 1",
		Status:    pb.TaskStatus_TASK_STATUS_TODO,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
)
