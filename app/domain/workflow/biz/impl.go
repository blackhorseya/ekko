package biz

import (
	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/entity/domain/workflow/agg"
	"github.com/blackhorseya/ekko/entity/domain/workflow/biz"
	"github.com/blackhorseya/ekko/entity/domain/workflow/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
)

type impl struct {
	issues repo.IIssueRepo
}

func NewWorkflowBiz(issues repo.IIssueRepo) biz.IWorkflowBiz {
	return &impl{issues: issues}
}

func (i *impl) CreateTodo(ctx contextx.Contextx, who *idM.User, title string) (item *agg.Issue, err error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) ListTodos(
	ctx contextx.Contextx,
	who *idM.User,
	opts biz.ListTodosOptions,
) (items []*agg.Issue, total int, err error) {
	// TODO implement me
	panic("implement me")
}
