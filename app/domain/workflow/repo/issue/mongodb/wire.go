//go:build wireinject

//go:generate wire

package mongodb

import (
	"github.com/blackhorseya/ekko/entity/domain/workflow/repo"
	"github.com/blackhorseya/ekko/pkg/storage/mongodbx"
	"github.com/google/wire"
)

// IssueMongodbSet is the set of IssueMongodbSet.
var IssueMongodbSet = wire.NewSet(
	NewIssueRepo,
	mongodbx.NewClient,
)

func BuildIssueRepo() (repo.IIssueRepo, error) {
	panic(wire.Build(IssueMongodbSet))
}
