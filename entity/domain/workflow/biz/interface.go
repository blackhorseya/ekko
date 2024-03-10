//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/entity/domain/workflow/agg"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

// ListTodosOptions is the options for list todos.
type ListTodosOptions struct {
	Page int
	Size int
}

// IWorkflowBiz is the interface for workflow business logic.
type IWorkflowBiz interface {
	// CreateTodo is to create a todo item.
	CreateTodo(ctx contextx.Contextx, who *idM.User, title string) (item *agg.Issue, err error)

	// ListTodos is to list todo items.
	ListTodos(ctx contextx.Contextx, who *idM.User, opts ListTodosOptions) (items []*agg.Issue, total int, err error)
}
