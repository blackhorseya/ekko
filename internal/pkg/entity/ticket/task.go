package ticket

import (
	"time"

	"github.com/blackhorseya/gocommon/pkg/timex"
	"github.com/blackhorseya/todo-app/pb"
)

// Task declare a task information
type Task struct {
	ID        uint64        `json:"id"`
	Title     string        `json:"title"`
	Status    pb.TaskStatus `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

// NewTaskResponse serve caller to given Task to pb.Task
func NewTaskResponse(val *Task) *pb.Task {
	if val == nil {
		return nil
	}

	return &pb.Task{
		Id:        val.ID,
		Title:     val.Title,
		Status:    val.Status,
		CreatedAt: val.CreatedAt.UTC().Format(timex.RFC3339Mill),
		UpdatedAt: val.UpdatedAt.UTC().Format(timex.RFC3339Mill),
	}
}
