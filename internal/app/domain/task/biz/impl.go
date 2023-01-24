package biz

import (
	"github.com/blackhorseya/todo-app/internal/app/domain/task/biz/repo"
	"github.com/blackhorseya/todo-app/internal/pkg/errorx"
	"github.com/blackhorseya/todo-app/pkg/contextx"
	tb "github.com/blackhorseya/todo-app/pkg/entity/domain/task/biz"
	"github.com/blackhorseya/todo-app/pkg/entity/domain/task/model"
	"github.com/blackhorseya/todo-app/pkg/genx"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var TaskSet = wire.NewSet(NewImpl)

type impl struct {
	repo      repo.IRepo
	generator genx.Generator
}

func NewImpl(repo repo.IRepo, generator genx.Generator) tb.IBiz {
	return &impl{
		repo:      repo,
		generator: generator,
	}
}

func (i *impl) Liveness(ctx contextx.Contextx) error {
	return nil
}

func (i *impl) Readiness(ctx contextx.Contextx) error {
	return nil
}

func (i *impl) GetByID(ctx contextx.Contextx, id int64) (info *model.Task, err error) {
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

func (i *impl) List(ctx contextx.Contextx, page, size int) (info []*model.Task, total int, err error) {
	if page < 0 {
		ctx.Error(errorx.ErrInvalidPage.Error(), zap.Int("page", page), zap.Int("size", size))
		return nil, 0, errorx.ErrInvalidPage
	}

	if size < 0 {
		ctx.Error(errorx.ErrInvalidSize.Error(), zap.Int("page", page), zap.Int("size", size))
		return nil, 0, errorx.ErrInvalidSize
	}

	condition := repo.QueryTasksCondition{
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

func (i *impl) Create(ctx contextx.Contextx, title string) (info *model.Task, err error) {
	created := &model.Task{
		Id:        i.generator.Int64(),
		Title:     title,
		Status:    model.TaskStatus_TASK_STATUS_TODO,
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

func (i *impl) UpdateStatus(ctx contextx.Contextx, id int64, status model.TaskStatus) (info *model.Task, err error) {
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
	ret, err := i.repo.Update(ctx, exists)
	if err != nil {
		ctx.Error(errorx.ErrUpdateStatusTask.Error(), zap.Error(err), zap.Any("updated", exists))
		return nil, errorx.ErrUpdateStatusTask
	}

	return ret, nil
}

func (i *impl) ChangeTitle(ctx contextx.Contextx, id int64, title string) (info *model.Task, err error) {
	if len(title) == 0 {
		ctx.Error(errorx.ErrInvalidTitle.Error())
		return nil, errorx.ErrInvalidTitle
	}

	exists, err := i.repo.GetByID(ctx, id)
	if err != nil {
		ctx.Error(errorx.ErrGetTask.Error(), zap.Error(err))
		return nil, errorx.ErrGetTask
	}
	if exists == nil {
		ctx.Error(errorx.ErrTaskNotExists.Error())
		return nil, errorx.ErrTaskNotExists
	}

	exists.Title = title
	ret, err := i.repo.Update(ctx, exists)
	if err != nil {
		ctx.Error(errorx.ErrChangeTitleTask.Error(), zap.Error(err), zap.Any("updated", exists))
		return nil, errorx.ErrChangeTitleTask
	}

	return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id int64) error {
	err := i.repo.DeleteByID(ctx, id)
	if err != nil {
		ctx.Error(errorx.ErrDeleteTask.Error(), zap.Error(err), zap.Int64("id", id))
		return errorx.ErrDeleteTask
	}

	return nil
}
