package todo

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo"
	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/er"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/google/uuid"
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

func (i *impl) GetByID(ctx contextx.Contextx, id string) (task *todo.Task, err error) {
	if len(id) == 0 {
		i.logger.Error(er.ErrMissingID.Error())
		return nil, er.ErrMissingID
	}

	_, err = uuid.Parse(id)
	if err != nil {
		i.logger.Error(er.ErrInvalidID.Error(), zap.Error(err), zap.String("id", id))
		return nil, er.ErrInvalidID
	}

	ret, err := i.repo.GetByID(ctx, id)
	if err != nil {
		i.logger.Error(er.ErrGetTask.Error(), zap.Error(err), zap.String("id", id))
		return nil, er.ErrGetTask
	}
	if ret == nil {
		i.logger.Error(er.ErrTaskNotExists.Error(), zap.String("id", id))
		return nil, er.ErrTaskNotExists
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, start, end int) (tasks []*todo.Task, total int, err error) {
	// todo: 2021-05-01|23:27|doggy|implement me
	panic("implement me")
}

func (i *impl) Create(ctx contextx.Contextx, title string) (task *todo.Task, err error) {
	// todo: 2021-05-01|23:27|doggy|implement me
	panic("implement me")
}

func (i *impl) UpdateStatus(ctx contextx.Contextx, id string, status bool) (task *todo.Task, err error) {
	// todo: 2021-05-01|23:27|doggy|implement me
	panic("implement me")
}

func (i *impl) ChangeTitle(ctx contextx.Contextx, id string, title string) (task *todo.Task, err error) {
	// todo: 2021-05-01|23:27|doggy|implement me
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id string) error {
	// todo: 2021-05-01|23:27|doggy|implement me
	panic("implement me")
}
