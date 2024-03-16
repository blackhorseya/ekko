package biz

import (
	"time"

	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/entity/domain/workflow/agg"
	"github.com/blackhorseya/ekko/entity/domain/workflow/biz"
	"github.com/blackhorseya/ekko/entity/domain/workflow/model"
	"github.com/blackhorseya/ekko/entity/domain/workflow/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"go.uber.org/zap"
)

type impl struct {
	issues repo.IIssueRepo
}

// NewWorkflowBiz is to create a new workflow business logic.
func NewWorkflowBiz(issues repo.IIssueRepo) biz.IWorkflowBiz {
	return &impl{issues: issues}
}

func (i *impl) GetTodoByID(ctx contextx.Contextx, who *idM.User, id string) (item *agg.Issue, err error) {
	// todo: 2024/3/16|sean|implement me
	panic("implement me")
}

func (i *impl) CreateTodo(ctx contextx.Contextx, who *idM.User, title string) (item *agg.Issue, err error) {
	issue := &agg.Issue{
		Ticket: &model.Ticket{
			ID:        "",
			Title:     title,
			Completed: false,
			OwnerID:   who.ID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}
	err = i.issues.Create(ctx, issue)
	if err != nil {
		ctx.Error("repo.IIssueRepo.Create", zap.Error(err))
		return nil, err
	}

	return issue, nil
}

func (i *impl) ListTodos(
	ctx contextx.Contextx,
	who *idM.User,
	opts biz.ListTodosOptions,
) (items []*agg.Issue, total int, err error) {
	ret, t, err := i.issues.List(ctx, repo.ListIssueOptions{
		OwnerID: who.ID,
		Limit:   opts.Size,
		Offset:  (opts.Page - 1) * opts.Size,
	})
	if err != nil {
		ctx.Error("repo.IIssueRepo.List", zap.Error(err))
		return nil, 0, err
	}

	return ret, t, nil
}

func (i *impl) CompleteTodoByID(ctx contextx.Contextx, who *idM.User, id string) (err error) {
	// todo: 2024/3/16|sean|implement me
	panic("implement me")
}
