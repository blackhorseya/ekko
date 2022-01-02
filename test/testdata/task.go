package testdata

import (
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/blackhorseya/todo-app/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	// TaskOID1 task object id 1
	TaskOID1 = primitive.NewObjectID()

	// Task1 task 1
	Task1 = &todo.Task{
		ID:    TaskOID1,
		Title: "task 1",
	}

	// TaskUpdate1 task update 1
	TaskUpdate1 = &todo.Task{
		ID:    TaskOID1,
		Title: "task 1 update",
	}

	// TaskCreate1 task create 1
	TaskCreate1 = &todo.Task{
		ID:     primitive.ObjectID{},
		Title:  "task 1",
		Status: pb.TaskStatus_TASK_STATUS_TODO,
	}
)
