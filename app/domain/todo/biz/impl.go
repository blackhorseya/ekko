package biz

import (
	"fmt"

	"github.com/blackhorseya/ekko/entity/domain/todo/biz"
	"github.com/blackhorseya/ekko/entity/domain/todo/model"
	"github.com/blackhorseya/ekko/entity/domain/todo/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

// NewNilTodoBiz creates a new nil todo biz instance.
func NewNilTodoBiz() biz.ITodoBiz {
	return nil
}

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
	return i.todos.List(ctx, repo.ListCondition{
		Limit: opts.Size,
		Skip:  (opts.Page - 1) * opts.Size,
	})
}

func (i *impl) CreateTodo(ctx contextx.Contextx, title string) (item *model.Todo, err error) {
	todo, err := model.NewTodo(title)
	if err != nil {
		return nil, fmt.Errorf("new todo model: %w", err)
	}

	err = i.todos.Create(ctx, todo)
	if err != nil {
		return nil, fmt.Errorf("create todo: %w", err)
	}

	return todo, nil
}

func (i *impl) CompleteTodo(ctx contextx.Contextx, id string) (item *model.Todo, err error) {
	item, err = i.todos.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get todo by id: %w", err)
	}

	item.Done = true
	err = i.todos.Update(ctx, item)
	if err != nil {
		return nil, fmt.Errorf("update todo: %w", err)
	}

	return item, nil
}
