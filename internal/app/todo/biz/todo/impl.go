package todo

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/er"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/blackhorseya/todo-app/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type impl struct {
	repo repo.IRepo
}

// NewImpl serve caller to create an ITodoBiz
func NewImpl(repo repo.IRepo) ITodoBiz {
	return &impl{
		repo: repo,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id primitive.ObjectID) (task *todo.Task, err error) {
	if id == primitive.NilObjectID {
		ctx.Error(er.ErrEmptyID.Error())
		return nil, er.ErrEmptyID
	}

	ret, err := i.repo.GetByID(ctx, id)
	if err != nil {
		ctx.Error(er.ErrGetTask.Error(), zap.Error(err), zap.String("id", id.Hex()))
		return nil, er.ErrGetTask
	}
	if ret == nil {
		ctx.Error(er.ErrTaskNotExists.Error(), zap.String("id", id.Hex()))
		return nil, er.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, start, end int) (tasks []*todo.Task, total int, err error) {
	if start < 0 {
		ctx.Error(er.ErrInvalidStart.Error(), zap.Int("start", start), zap.Int("end", end))
		return nil, 0, er.ErrInvalidStart
	}

	if end < 0 {
		ctx.Error(er.ErrInvalidEnd.Error(), zap.Int("start", start), zap.Int("end", end))
		return nil, 0, er.ErrInvalidEnd
	}

	ret, err := i.repo.List(ctx, end-start+1, start)
	if err != nil {
		ctx.Error(er.ErrListTasks.Error(), zap.Error(err), zap.Int("start", start), zap.Int("end", end))
		return nil, 0, er.ErrListTasks
	}
	if ret == nil {
		ctx.Error(er.ErrTaskNotExists.Error(), zap.Int("start", start), zap.Int("end", end))
		return nil, 0, er.ErrTaskNotExists
	}

	total, err = i.repo.Count(ctx)
	if err != nil {
		ctx.Error(er.ErrCountTask.Error(), zap.Error(err))
		return nil, 0, er.ErrCountTask
	}

	return ret, total, nil
}

func (i *impl) Create(ctx contextx.Contextx, title string) (task *todo.Task, err error) {
	if len(title) == 0 {
		ctx.Error(er.ErrEmptyTitle.Error())
		return nil, er.ErrEmptyTitle
	}

	newTask := &todo.Task{
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

func (i *impl) UpdateStatus(ctx contextx.Contextx, id primitive.ObjectID, status pb.TaskStatus) (task *todo.Task, err error) {
	if id == primitive.NilObjectID {
		ctx.Error(er.ErrEmptyID.Error())
		return nil, er.ErrEmptyID
	}

	found, err := i.repo.GetByID(ctx, id)
	if err != nil {
		ctx.Error(er.ErrGetTask.Error(), zap.Error(err), zap.String("id", id.Hex()))
		return nil, er.ErrGetTask
	}
	if found == nil {
		ctx.Error(er.ErrTaskNotExists.Error(), zap.String("id", id.Hex()))
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

func (i *impl) ChangeTitle(ctx contextx.Contextx, id primitive.ObjectID, title string) (task *todo.Task, err error) {
	if id == primitive.NilObjectID {
		ctx.Error(er.ErrEmptyID.Error())
		return nil, er.ErrEmptyID
	}

	if len(title) == 0 {
		ctx.Error(er.ErrEmptyTitle.Error())
		return nil, er.ErrEmptyTitle
	}

	found, err := i.repo.GetByID(ctx, id)
	if err != nil {
		ctx.Error(er.ErrGetTask.Error(), zap.Error(err), zap.String("id", id.Hex()))
		return nil, er.ErrGetTask
	}
	if found == nil {
		ctx.Error(er.ErrTaskNotExists.Error(), zap.String("id", id.Hex()))
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

func (i *impl) Delete(ctx contextx.Contextx, id primitive.ObjectID) error {
	if id == primitive.NilObjectID {
		ctx.Error(er.ErrEmptyID.Error())
		return er.ErrEmptyID
	}

	err := i.repo.Remove(ctx, id)
	if err != nil {
		ctx.Error(er.ErrDeleteTask.Error(), zap.Error(err), zap.String("id", id.Hex()))
		return er.ErrDeleteTask
	}

	return nil
}
