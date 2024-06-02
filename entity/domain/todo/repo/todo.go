//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package repo

import (
	"github.com/blackhorseya/ekko/entity/domain/todo/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

// ListCondition is the condition for listing todo
type ListCondition struct {
	Limit  int
	Offset int
}

// ITodoRepo is the interface that defines the methods that the todo repository should implement
type ITodoRepo interface {
	List(ctx contextx.Contextx, condition ListCondition) (items []*model.Todo, total int, err error)
	GetByID(ctx contextx.Contextx, id string) (item *model.Todo, err error)
	Create(ctx contextx.Contextx, item *model.Todo) (err error)
	Update(ctx contextx.Contextx, item *model.Todo) (err error)
}
