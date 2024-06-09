package biz

import (
	"fmt"

	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/entity/domain/todo/biz"
	"github.com/blackhorseya/ekko/entity/domain/todo/model"
	"github.com/blackhorseya/ekko/entity/domain/todo/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/otelx"
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
	ctx, span := otelx.StartSpan(ctx, "biz")
	defer span.End()

	who, err := idM.FromContext(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("get identity from context: %w", err)
	}

	return i.todos.List(ctx, repo.ListCondition{
		CreatedBy: who.ID,
		Limit:     opts.Size,
		Skip:      (opts.Page - 1) * opts.Size,
	})
}

func (i *impl) CreateTodo(ctx contextx.Contextx, title string) (item *model.Todo, err error) {
	ctx, span := otelx.StartSpan(ctx, "biz")
	defer span.End()

	who, err := idM.FromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("get identity from context: %w", err)
	}

	todo, err := model.NewTodo(title)
	if err != nil {
		return nil, fmt.Errorf("new todo model: %w", err)
	}
	todo.CreatedBy = who.ID

	err = i.todos.Create(ctx, todo)
	if err != nil {
		return nil, fmt.Errorf("create todo: %w", err)
	}

	return todo, nil
}

func (i *impl) CompleteTodo(ctx contextx.Contextx, id string) (item *model.Todo, err error) {
	ctx, span := otelx.StartSpan(ctx, "biz")
	defer span.End()

	who, err := idM.FromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("get identity from context: %w", err)
	}

	item, err = i.todos.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get todo by id: %w", err)
	}
	if item.CreatedBy != who.ID {
		return nil, fmt.Errorf("todo not belong to you")
	}

	item.Done = true
	err = i.todos.Update(ctx, item)
	if err != nil {
		return nil, fmt.Errorf("update todo: %w", err)
	}

	return item, nil
}
