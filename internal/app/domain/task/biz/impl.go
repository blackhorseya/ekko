package biz

import (
	"strings"

	taskB "github.com/blackhorseya/ekko/entity/domain/task/biz"
	taskM "github.com/blackhorseya/ekko/entity/domain/task/model"
	taskR "github.com/blackhorseya/ekko/internal/app/domain/task/biz/repo"
	"github.com/blackhorseya/ekko/internal/pkg/errorx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"go.uber.org/zap"
)

type impl struct {
	repo taskR.IRepo
}

// NewImpl will create an object that implement IBiz interface
func NewImpl(repo taskR.IRepo) taskB.IBiz {
	return &impl{
		repo: repo,
	}
}

func (i *impl) GetTicketByID(ctx contextx.Contextx, id string) (ticket *taskM.Ticket, err error) {
	id = strings.Trim(id, " ")
	if id == "" {
		ctx.Error("id is empty then error")
		return nil, errorx.ErrInvalidID
	}

	ret, err := i.repo.GetTicketByID(ctx, id)
	if err != nil {
		ctx.Error("get ticket by id from repo failed", zap.Error(err), zap.String("id", id))
		return nil, err
	}
	if ret == nil {
		ctx.Error("ticket is not exists", zap.String("id", id))
		return nil, errorx.ErrTicketNotExists
	}

	return ret, nil
}

func (i *impl) ListTickets(ctx contextx.Contextx, condition taskB.ListTicketsCondition) (tickets []*taskM.Ticket, total int, err error) {
	// todo: 2023/7/31|sean|implement me
	panic("implement me")
}

func (i *impl) CreateTicket(ctx contextx.Contextx, title string) (ticket *taskM.Ticket, err error) {
	// todo: 2023/7/31|sean|implement me
	panic("implement me")
}

func (i *impl) UpdateTicketStatus(ctx contextx.Contextx, id string, status taskM.TicketStatus) (ticket *taskM.Ticket, err error) {
	// todo: 2023/7/31|sean|implement me
	panic("implement me")
}

func (i *impl) DeleteTicket(ctx contextx.Contextx, id string) error {
	// todo: 2023/7/31|sean|implement me
	panic("implement me")
}
