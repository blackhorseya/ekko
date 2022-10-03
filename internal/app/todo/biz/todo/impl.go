package todo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/er"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/ticket"
	"github.com/blackhorseya/todo-app/pb"
	"go.uber.org/zap"
)

type impl struct {
	repo repo.ITodoRepo
}

// NewImpl serve caller to create an ITodoBiz
func NewImpl(repo repo.ITodoRepo) ITodoBiz {
	return &impl{
		repo: repo,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id uint64) (task *ticket.Task, err error) {
	ret, err := i.repo.GetByID(ctx, id)
	if err != nil {
		ctx.Error(er.ErrGetTask.Error(), zap.Error(err))
		return nil, er.ErrGetTask
	}
	if ret == nil {
		ctx.Error(er.ErrTaskNotExists.Error())
		return nil, er.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, page, size int) (tasks []*ticket.Task, total int, err error) {
	if page < 0 {
		ctx.Error(er.ErrInvalidStart.Error(), zap.Int("page", page), zap.Int("size", size))
		return nil, 0, er.ErrInvalidStart
	}

	if size < 0 {
		ctx.Error(er.ErrInvalidEnd.Error(), zap.Int("page", page), zap.Int("size", size))
		return nil, 0, er.ErrInvalidEnd
	}

	ret, err := i.repo.List(ctx, repo.QueryTodoCondition{
		Limit:  size - page + 1,
		Offset: page,
	})
	if err != nil {
		ctx.Error(er.ErrListTasks.Error(), zap.Error(err), zap.Int("page", page), zap.Int("size", size))
		return nil, 0, er.ErrListTasks
	}
	if ret == nil {
		ctx.Error(er.ErrTaskNotExists.Error(), zap.Int("page", page), zap.Int("size", size))
		return nil, 0, er.ErrTaskNotExists
	}

	total, err = i.repo.Count(ctx)
	if err != nil {
		ctx.Error(er.ErrCountTask.Error(), zap.Error(err))
		return nil, 0, er.ErrCountTask
	}

	return ret, total, nil
}

func (i *impl) Create(ctx contextx.Contextx, title string) (task *ticket.Task, err error) {
	if len(title) == 0 {
		ctx.Error(er.ErrEmptyTitle.Error())
		return nil, er.ErrEmptyTitle
	}

	newTask := &ticket.Task{
		// todo: 2022/10/4|sean|fix me
		ID:     1,
		Title:  title,
		Status: pb.TaskStatus_TASK_STATUS_TODO,
	}
	ret, err := i.repo.Create(ctx, newTask)
	if err != nil {
		ctx.Error(er.ErrCreateTask.Error(), zap.Error(err), zap.String("title", title))
		return nil, er.ErrCreateTask
	}

	return ret, nil
}

func (i *impl) UpdateStatus(ctx contextx.Contextx, id uint64, status pb.TaskStatus) (task *ticket.Task, err error) {
	found, err := i.repo.GetByID(ctx, id)
	if err != nil {
		ctx.Error(er.ErrGetTask.Error(), zap.Error(err))
		return nil, er.ErrGetTask
	}
	if found == nil {
		ctx.Error(er.ErrTaskNotExists.Error())
		return nil, er.ErrTaskNotExists
	}

	found.Status = status
	ret, err := i.repo.Update(ctx, found)
	if err != nil {
		ctx.Error(er.ErrUpdateStatusTask.Error(), zap.Error(err), zap.Any("updated", found))
		return nil, er.ErrUpdateStatusTask
	}

	return ret, nil
}

func (i *impl) ChangeTitle(ctx contextx.Contextx, id uint64, title string) (task *ticket.Task, err error) {
	if len(title) == 0 {
		ctx.Error(er.ErrEmptyTitle.Error())
		return nil, er.ErrEmptyTitle
	}

	found, err := i.repo.GetByID(ctx, id)
	if err != nil {
		ctx.Error(er.ErrGetTask.Error(), zap.Error(err))
		return nil, er.ErrGetTask
	}
	if found == nil {
		ctx.Error(er.ErrTaskNotExists.Error())
		return nil, er.ErrTaskNotExists
	}

	found.Title = title
	ret, err := i.repo.Update(ctx, found)
	if err != nil {
		ctx.Error(er.ErrChangeTitleTask.Error(), zap.Error(err), zap.Any("updated", found))
		return nil, er.ErrChangeTitleTask
	}

	return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id uint64) error {
	err := i.repo.Remove(ctx, id)
	if err != nil {
		ctx.Error(er.ErrDeleteTask.Error(), zap.Error(err))
		return er.ErrDeleteTask
	}

	return nil
}
