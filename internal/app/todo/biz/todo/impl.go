package todo

import (
	"fmt"

	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo"
	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/er"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/blackhorseya/todo-app/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
}

// NewImpl serve caller to create an IBiz
func NewImpl(logger *zap.Logger, repo repo.IRepo) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "TodoBiz")),
		repo:   repo,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id primitive.ObjectID) (task *todo.Task, err error) {
	if id == primitive.NilObjectID {
		i.logger.Error(er.ErrEmptyID.Error())
		return nil, er.ErrEmptyID
	}

	ret, err := i.repo.GetByID(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrGetTask.Error(), zap.Error(err), zap.String("id", id.Hex()))
		return nil, er.ErrGetTask
	}
	if ret == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.String("id", id.Hex()))
		return nil, er.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, start, end int) (tasks []*pb.Task, total int, err error) {
	if start < 0 {
		i.logger.Error(er.ErrInvalidStart.Error(), zap.Int("start", start), zap.Int("end", end))
		return nil, 0, er.ErrInvalidStart
	}

	if end < 0 {
		i.logger.Error(er.ErrInvalidEnd.Error(), zap.Int("start", start), zap.Int("end", end))
		return nil, 0, er.ErrInvalidEnd
	}

	ret, err := i.repo.List(ctx, end-start+1, start)
	if err != nil {
		i.logger.Error(er.ErrListTasks.Error(), zap.Error(err), zap.Int("start", start), zap.Int("end", end))
		return nil, 0, er.ErrListTasks
	}
	if ret == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.Int("start", start), zap.Int("end", end))
		return nil, 0, er.ErrTaskNotExists
	}

	total, err = i.repo.Count(ctx)
	if err != nil {
		i.logger.Error(er.ErrCountTask.Error(), zap.Error(err))
		return nil, 0, er.ErrCountTask
	}

	// todo: 2021-12-19|00:50|Sean|impl me
	panic("impl me")
	// return ret, total, nil
}

func (i *impl) Create(ctx contextx.Contextx, title string) (task *todo.Task, err error) {
	if len(title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error())
		return nil, er.ErrEmptyTitle
	}

	newTask := &todo.Task{
		Title:     title,
		Completed: false,
	}
	ret, err := i.repo.Create(ctx, newTask)
	if err != nil {
		i.logger.Error(er.ErrCreateTask.Error(), zap.Error(err), zap.String("title", title))
		return nil, er.ErrCreateTask
	}

	return ret, nil
}

func (i *impl) UpdateStatus(ctx contextx.Contextx, id primitive.ObjectID, status bool) (task *pb.Task, err error) {
	exists, err := i.repo.GetByID(ctx, primitive.ObjectID{})
	if err != nil {
		i.logger.Error(er.ErrGetTask.Error(), zap.Error(err), zap.String("id", id.Hex()))
		return nil, er.ErrGetTask
	}
	if exists == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.String("id", id.Hex()))
		return nil, er.ErrTaskNotExists
	}

	exists.Completed = status
	ret, err := i.repo.Update(ctx, exists)
	if err != nil {
		i.logger.Error(er.ErrUpdateStatusTask.Error(), zap.Error(err), zap.String("id", id.Hex()), zap.Bool("status", status))
		return nil, er.ErrUpdateStatusTask
	}

	fmt.Println(ret)
	// todo: 2021-12-19|00:50|Sean|impl me
	panic("impl me")
	// return ret, nil
}

func (i *impl) ChangeTitle(ctx contextx.Contextx, id primitive.ObjectID, title string) (task *pb.Task, err error) {
	if len(title) == 0 {
		i.logger.Error(er.ErrEmptyTitle.Error())
		return nil, er.ErrEmptyTitle
	}

	exists, err := i.repo.GetByID(ctx, primitive.ObjectID{})
	if err != nil {
		i.logger.Error(er.ErrGetTask.Error(), zap.Error(err), zap.String("id", id.Hex()))
		return nil, er.ErrGetTask
	}
	if exists == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.String("id", id.Hex()))
		return nil, er.ErrTaskNotExists
	}

	exists.Title = title
	ret, err := i.repo.Update(ctx, exists)
	if err != nil {
		i.logger.Error(er.ErrChangeTitleTask.Error(), zap.Error(err), zap.String("id", id.Hex()), zap.String("title", title))
		return nil, er.ErrChangeTitleTask
	}

	fmt.Println(ret)
	// todo: 2021-12-19|00:50|Sean|impl me
	panic("impl me")
	// return ret, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id primitive.ObjectID) error {
	err := i.repo.Remove(ctx, primitive.ObjectID{})
	if err != nil {
		i.logger.Error(er.ErrDeleteTask.Error(), zap.Error(err), zap.String("id", id.Hex()))
		return er.ErrDeleteTask
	}

	return nil
}
