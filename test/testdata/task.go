package testdata

import (
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// TaskOID1 task object id 1
	TaskOID1 = primitive.NewObjectID()

	// Task1 task 1
	Task1 = &todo.Task{
		ID:        TaskOID1,
		Title:     "task 1",
		Completed: false,
	}

	// TaskUpdate1 task update 1
	TaskUpdate1 = &todo.Task{
		ID:        TaskOID1,
		Title:     "task 1 update",
		Completed: true,
	}

	// TaskCreate1 task create 1
	TaskCreate1 = &todo.Task{
		ID:        primitive.ObjectID{},
		Title:     "task 1",
		Completed: false,
	}
)
