package biz

import (
	"strings"

	issueB "github.com/blackhorseya/ekko/entity/domain/issue/biz"
	issueM "github.com/blackhorseya/ekko/entity/domain/issue/model"
	"github.com/blackhorseya/ekko/entity/domain/user/model"
	"github.com/blackhorseya/ekko/internal/app/domain/issue/biz/repo"
	"github.com/blackhorseya/ekko/internal/pkg/errorx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/genx"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// IssueSet presents as issue biz provider set
var IssueSet = wire.NewSet(NewImpl)

type impl struct {
	repo      repo.IRepo
	generator genx.Generator
}

func NewImpl(repo repo.IRepo, generator genx.Generator) issueB.IBiz {
	return &impl{
		repo:      repo,
		generator: generator,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id int64) (info *issueM.Ticket, err error) {
	ret, err := i.repo.GetByID(ctx, id)
	if err != nil {
		ctx.Error(errorx.ErrGetTask.Error(), zap.Error(err), zap.Int64("id", id))
		return nil, errorx.ErrGetTask
	}
	if ret == nil {
		ctx.Error(errorx.ErrTaskNotExists.Error(), zap.Int64("id", id))
		return nil, errorx.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, page, size int) (info []*issueM.Ticket, total int, err error) {
	if page < 0 {
		ctx.Error(errorx.ErrInvalidPage.Error(), zap.Int("page", page), zap.Int("size", size))
		return nil, 0, errorx.ErrInvalidPage
	}

	if size < 0 {
		ctx.Error(errorx.ErrInvalidSize.Error(), zap.Int("page", page), zap.Int("size", size))
		return nil, 0, errorx.ErrInvalidSize
	}

	condition := repo.QueryTicketsCondition{
		Limit:  size,
		Offset: (page - 1) * size,
	}
	ret, err := i.repo.List(ctx, condition)
	if err != nil {
		ctx.Error(errorx.ErrListTasks.Error(), zap.Error(err), zap.Int("page", page), zap.Int("size", size))
		return nil, 0, errorx.ErrListTasks
	}
	if ret == nil {
		ctx.Error(errorx.ErrTaskNotExists.Error(), zap.Int("page", page), zap.Int("size", size))
		return nil, 0, errorx.ErrTaskNotExists
	}

	total, err = i.repo.Count(ctx, condition)
	if err != nil {
		ctx.Error(errorx.ErrCountTask.Error(), zap.Error(err))
		return nil, 0, errorx.ErrCountTask
	}

	return ret, total, nil
}

func (i *impl) Create(ctx contextx.Contextx, title string) (info *issueM.Ticket, err error) {
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		ctx.Error(errorx.ErrInvalidTitle.Error(), zap.String("title", title))
		return nil, errorx.ErrInvalidTitle
	}

	created := &issueM.Ticket{
		Id:        i.generator.Int64(),
		Title:     title,
		Status:    issueM.TicketStatus_TICKET_STATUS_TODO,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
	ret, err := i.repo.Create(ctx, created)
	if err != nil {
		ctx.Error(errorx.ErrCreateTask.Error(), zap.Error(err), zap.String("title", created.Title), zap.Int64("id", created.Id))
		return nil, errorx.ErrCreateTask
	}

	return ret, nil
}

func (i *impl) UpdateStatus(ctx contextx.Contextx, id int64, status issueM.TicketStatus) (info *issueM.Ticket, err error) {
	exists, err := i.repo.GetByID(ctx, id)
	if err != nil {
		ctx.Error(errorx.ErrGetTask.Error(), zap.Error(err))
		return nil, errorx.ErrGetTask
	}
	if exists == nil {
		ctx.Error(errorx.ErrTaskNotExists.Error())
		return nil, errorx.ErrTaskNotExists
	}

	exists.Status = status
	err = i.repo.Update(ctx, exists)
	if err != nil {
		ctx.Error(errorx.ErrUpdateStatusTask.Error(), zap.Error(err), zap.Any("updated", exists))
		return nil, errorx.ErrUpdateStatusTask
	}

	return exists, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id int64) error {
	err := i.repo.DeleteByID(ctx, id)
	if err != nil {
		ctx.Error(errorx.ErrDeleteTask.Error(), zap.Error(err), zap.Int64("id", id))
		return errorx.ErrDeleteTask
	}

	return nil
}

func (i *impl) ListTasks(ctx contextx.Contextx, who *model.Profile, condition issueB.ListTasksCondition) (tickets []*issueM.Ticket, total int, err error) {
	// todo: 2023/4/16|sean|impl me
	panic("implement me")
}
