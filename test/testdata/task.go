package testdata

import (
	"time"

	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// Task1 task 1
	Task1 = &todo.Task{
		ID:        primitive.NewObjectID(),
		Title:     "task 1",
		Completed: false,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
)
