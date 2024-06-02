package biz

import (
	"github.com/blackhorseya/ekko/entity/domain/todo/biz"
	"github.com/blackhorseya/ekko/entity/domain/todo/model"
	"github.com/blackhorseya/ekko/entity/domain/todo/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

type impl struct {
	todos repo.ITodoRepo
}

// NewTodoBiz creates a new todo biz instance.
func NewTodoBiz(todos repo.ITodoRepo) biz.ITodoBiz {
	return &impl{
		todos: todos,
	}
}

func (i *impl) ListTodo(ctx contextx.Contextx, opts biz.ListTodoOptions) (items []*model.Todo, total int, err error) {
	// todo: 2024/6/3|sean|implement me
	panic("implement me")
}

func (i *impl) CreateTodo(ctx contextx.Contextx, title string) (item *model.Todo, err error) {
	// todo: 2024/6/3|sean|implement me
	panic("implement me")
}

func (i *impl) CompleteTodo(ctx contextx.Contextx, id string) (item *model.Todo, err error) {
	// todo: 2024/6/3|sean|implement me
	panic("implement me")
}
