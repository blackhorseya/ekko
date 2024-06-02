//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/ekko/entity/domain/todo/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

// ListTodoOptions is the options for listing todo
type ListTodoOptions struct {
	Page int
	Size int
}

// ITodoBiz is the interface that defines the methods that the todo business logic should implement
type ITodoBiz interface {
	ListTodo(ctx contextx.Contextx, opts ListTodoOptions) (items []*model.Todo, total int, err error)
	CreateTodo(ctx contextx.Contextx, title string) (item *model.Todo, err error)
	CompleteTodo(ctx contextx.Contextx, id string) (item *model.Todo, err error)
}
