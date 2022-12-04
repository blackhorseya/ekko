package biz

import (
	"github.com/blackhorseya/todo-app/pkg/contextx"
	tb "github.com/blackhorseya/todo-app/pkg/entity/domain/todo/biz"
	"github.com/blackhorseya/todo-app/pkg/entity/domain/todo/model"
)

type impl struct {
}

func NewImpl() tb.IBiz {
	return &impl{}
}

func (i *impl) Liveness(ctx contextx.Contextx) error {
	return nil
}

func (i *impl) Readiness(ctx contextx.Contextx) error {
	return nil
}

func (i *impl) GetByID(ctx contextx.Contextx, id int64) (info *model.Task, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) List(ctx contextx.Contextx, page, size int) (info []*model.Task, total int, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) Create(ctx contextx.Contextx, title string) (info *model.Task, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) UpdateStatus(ctx contextx.Contextx, id int64, status model.TaskStatus) (info *model.Task, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) ChangeTitle(ctx contextx.Contextx, id int64, title string) (info *model.Task, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id int64) error {
	// TODO implement me
	panic("implement me")
}
