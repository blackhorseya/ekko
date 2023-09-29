package biz

import (
	"strings"
	"time"

	taskB "github.com/blackhorseya/ekko/entity/domain/task/biz"
	taskM "github.com/blackhorseya/ekko/entity/domain/task/model"
	taskR "github.com/blackhorseya/ekko/internal/app/domain/task/biz/repo"
	"github.com/blackhorseya/ekko/internal/pkg/errorx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/google/uuid"
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TaskBizSet will be used by wire
var TaskBizSet = wire.NewSet(NewImpl, taskR.NewMariadb)

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

func (i *impl) ListTickets(ctx contextx.Contextx, passCondition taskB.ListTicketsCondition) (tickets []*taskM.Ticket, total int, err error) {
	if passCondition.Page < 1 {
		ctx.Error("page is less than 1 then error", zap.Int("page", passCondition.Page))
		return nil, 0, errorx.ErrInvalidPage
	}

	if passCondition.Size < 1 {
		ctx.Error("size is less than 1 then error", zap.Int("size", passCondition.Size))
		return nil, 0, errorx.ErrInvalidSize
	}

	condition := taskR.ListTicketsCondition{
		Limit:  passCondition.Size,
		Offset: (passCondition.Page - 1) * passCondition.Size,
	}
	ret, total, err := i.repo.ListTickets(ctx, condition)
	if err != nil {
		ctx.Error("list tickets from repo failed", zap.Error(err), zap.Any("pass_condition", passCondition), zap.Any("condition", condition))
		return nil, 0, err
	}
	if ret == nil {
		ctx.Error("not found any tickets", zap.Any("pass_condition", passCondition), zap.Any("condition", condition))
		return nil, 0, errorx.ErrTicketNotExists
	}

	return ret, total, nil
}

func (i *impl) CreateTicket(ctx contextx.Contextx, title string) (ticket *taskM.Ticket, err error) {
	title = strings.Trim(title, " ")
	if title == "" {
		ctx.Error("title is empty then error")
		return nil, errorx.ErrInvalidTitle
	}

	now := time.Now()
	ticket = &taskM.Ticket{
		Id:          uuid.New().String(),
		Title:       title,
		Description: "",
		Status:      taskM.TicketStatus_TICKET_STATUS_TODO,
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
	}
	ret, err := i.repo.CreateTicket(ctx, ticket)
	if err != nil {
		ctx.Error("create ticket from repo failed", zap.Error(err), zap.Any("ticket", ticket))
		return nil, err
	}

	return ret, nil
}

func (i *impl) UpdateTicketStatus(ctx contextx.Contextx, id string, status taskM.TicketStatus) (ticket *taskM.Ticket, err error) {
	id = strings.Trim(id, " ")
	if id == "" {
		ctx.Error("id is empty then error")
		return nil, errorx.ErrInvalidID
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.Error("parse id to uuid failed", zap.Error(err), zap.String("id", id))
		return nil, errorx.ErrInvalidID
	}

	got, err := i.GetTicketByID(ctx, uid.String())
	if err != nil {
		return nil, err
	}

	now := time.Now()
	got.Status = status
	got.UpdatedAt = timestamppb.New(now)
	err = i.repo.UpdateTicket(ctx, got)
	if err != nil {
		ctx.Error("update ticket from repo failed", zap.Error(err), zap.Any("ticket", got))
		return nil, err
	}

	return got, nil
}

func (i *impl) DeleteTicket(ctx contextx.Contextx, id string) error {
	id = strings.Trim(id, " ")
	if id == "" {
		ctx.Error("id is empty then error")
		return errorx.ErrInvalidID
	}
	uid, err := uuid.Parse(id)
	if err != nil {
		ctx.Error("parse id to uuid failed", zap.Error(err), zap.String("id", id))
		return errorx.ErrInvalidID
	}

	err = i.repo.DeleteTicketByID(ctx, uid.String())
	if err != nil {
		ctx.Error("delete ticket by id from repo failed", zap.Error(err), zap.String("id", uid.String()))
		return err
	}

	return nil
}
