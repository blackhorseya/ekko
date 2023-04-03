package repo

import (
	"time"

	im "github.com/blackhorseya/ekko/pkg/entity/domain/issue/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type task struct {
	ID        int64           `json:"id" db:"id"`
	Title     string          `json:"title" db:"title"`
	Status    im.TicketStatus `json:"status" db:"status"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt time.Time       `json:"updated_at" db:"updated_at"`
}

func newTask(val *im.Ticket) *task {
	return &task{
		ID:        val.Id,
		Title:     val.Title,
		Status:    val.Status,
		CreatedAt: val.CreatedAt.AsTime().UTC(),
		UpdatedAt: val.UpdatedAt.AsTime().UTC(),
	}
}

func (t *task) ToEntity() *im.Ticket {
	return &im.Ticket{
		Id:        t.ID,
		Title:     t.Title,
		Status:    t.Status,
		CreatedAt: timestamppb.New(t.CreatedAt),
		UpdatedAt: timestamppb.New(t.UpdatedAt),
	}
}
