package testdata

import (
	"time"

	"github.com/blackhorseya/todo-app/internal/pkg/entity/ticket"
	"github.com/blackhorseya/todo-app/pb"
	"github.com/blackhorseya/todo-app/pkg/entity/domain/task/model"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	// Task2 task 2
	Task2 = &model.Task{
		Id:        1,
		Title:     "task 1",
		Status:    model.TaskStatus_TASK_STATUS_TODO,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}
)
