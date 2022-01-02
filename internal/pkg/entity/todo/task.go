package todo

import (
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
