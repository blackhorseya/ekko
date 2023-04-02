package testdata

import (
	"github.com/blackhorseya/ekko/pkg/entity/domain/task/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	// Task1 task 1
	Task1 = &model.Task{
		Id:        1,
		Title:     "task 1",
		Status:    model.TaskStatus_TASK_STATUS_TODO,
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
	}
)
