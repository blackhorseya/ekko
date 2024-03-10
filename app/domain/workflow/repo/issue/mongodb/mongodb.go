package mongodb

import (
	"time"

	"github.com/blackhorseya/ekko/entity/domain/workflow/agg"
	"github.com/blackhorseya/ekko/entity/domain/workflow/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	timeoutDuration = 5 * time.Second

	dbName   = "ekko"
	collName = "issues"
)

type impl struct {
	rw *mongo.Client
}

// NewIssueRepo is to create a new issue repository.
func NewIssueRepo(rw *mongo.Client) repo.IIssueRepo {
	return &impl{rw: rw}
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (item *agg.Issue, err error) {
	// todo: 2024/3/10|sean|implement me
	panic("implement me")
}

func (i *impl) Create(ctx contextx.Contextx, item *agg.Issue) (err error) {
	// todo: 2024/3/10|sean|implement me
	panic("implement me")
}

func (i *impl) Update(ctx contextx.Contextx, item *agg.Issue) (err error) {
	// todo: 2024/3/10|sean|implement me
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id string) (err error) {
	// todo: 2024/3/10|sean|implement me
	panic("implement me")
}
