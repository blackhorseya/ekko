//go:build wireinject

//go:generate wire

package biz

import (
	"github.com/blackhorseya/ekko/app/domain/workflow/repo/issue/mongodb"
	"github.com/blackhorseya/ekko/entity/domain/workflow/biz"
	"github.com/google/wire"
)

// DefaultWorkflowSet is the set of DefaultWorkflowSet.
var DefaultWorkflowSet = wire.NewSet(
	NewWorkflowBiz,
	mongodb.IssueMongodbSet,
)

func BuildWorkflowBiz() (biz.IWorkflowBiz, error) {
	panic(wire.Build(DefaultWorkflowSet))
}
