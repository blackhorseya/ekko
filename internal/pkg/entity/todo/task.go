package todo

import (
	"github.com/blackhorseya/gocommon/pkg/timex"
	"github.com/blackhorseya/todo-app/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task declare a task information
type Task struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	Status    pb.TaskStatus      `json:"status" bson:"status"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

// NewTaskResponse return *pb.Task
func NewTaskResponse(t *Task) *pb.Task {
	return &pb.Task{
		Id:        t.ID.Hex(),
		Title:     t.Title,
		Status:    t.Status,
		CreatedAt: t.CreatedAt.Time().Format(timex.RFC3339Mill),
		UpdatedAt: t.UpdatedAt.Time().Format(timex.RFC3339Mill),
	}
}
