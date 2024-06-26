package todo

import (
	"github.com/blackhorseya/ekko/entity/domain/todo/model"
	"github.com/blackhorseya/ekko/entity/domain/todo/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

type empty struct{}

// NewNil creates a new todo repo instance.
func NewNil() repo.ITodoRepo {
	return &empty{}
}

func (x empty) List(ctx contextx.Contextx, condition repo.ListCondition) (items []*model.Todo, total int, err error) {
	panic("implement me")
}

func (x empty) GetByID(ctx contextx.Contextx, id string) (item *model.Todo, err error) {
	panic("implement me")
}

func (x empty) Create(ctx contextx.Contextx, item *model.Todo) (err error) {
	panic("implement me")
}

func (x empty) Update(ctx contextx.Contextx, item *model.Todo) (err error) {
	panic("implement me")
}
