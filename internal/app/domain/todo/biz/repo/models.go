package repo

import (
	"time"

	"github.com/blackhorseya/todo-app/pkg/entity/domain/todo/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type task struct {
	ID        int64            `json:"id" db:"id"`
	Title     string           `json:"title" db:"title"`
	Status    model.TaskStatus `json:"status" db:"status"`
	CreatedAt time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt time.Time        `json:"updated_at" db:"updated_at"`
}

func newTask(val *model.Task) *task {
	return &task{
		ID:        val.Id,
		Title:     val.Title,
		Status:    val.Status,
		CreatedAt: val.CreatedAt.AsTime().UTC(),
		UpdatedAt: val.UpdatedAt.AsTime().UTC(),
	}
}

func (t *task) ToEntity() *model.Task {
	return &model.Task{
		Id:        t.ID,
		Title:     t.Title,
		Status:    t.Status,
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}
