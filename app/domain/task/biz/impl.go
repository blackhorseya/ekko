package biz

import (
	"github.com/blackhorseya/ekko/entity/domain/task/biz"
	"github.com/blackhorseya/ekko/entity/domain/task/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

type impl struct {
}

// NewTaskBiz is used to create a new task business logic.
func NewTaskBiz() biz.ITaskBiz {
	return &impl{}
}

func (i *impl) CreateTicket(ctx contextx.Contextx, title string) (item *model.Ticket, err error) {
	// todo: 2024/6/6|sean|add some logic here
	panic("implement me")
}

func (i *impl) GetTicketByID(ctx contextx.Contextx, id string) (item *model.Ticket, err error) {
	// todo: 2024/6/6|sean|add some logic here
	panic("implement me")
}

func (i *impl) ListTicket(
	ctx contextx.Contextx,
	options biz.ListTicketOptions,
) (items []*model.Ticket, total int, err error) {
	// todo: 2024/6/6|sean|add some logic here
	panic("implement me")
}

func (i *impl) UpdateTicket(ctx contextx.Contextx, id string, update *model.Ticket) (err error) {
	// todo: 2024/6/6|sean|add some logic here
	panic("implement me")
}
