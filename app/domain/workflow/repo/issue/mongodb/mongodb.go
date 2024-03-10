package mongodb

import (
	"time"

	"github.com/blackhorseya/ekko/entity/domain/workflow/agg"
	"github.com/blackhorseya/ekko/entity/domain/workflow/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (i *impl) List(ctx contextx.Contextx, options repo.ListIssueOptions) (items []*agg.Issue, total int, err error) {
	// todo: 2024/3/10|sean|implement me
	panic("implement me")
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (item *agg.Issue, err error) {
	// todo: 2024/3/10|sean|implement me
	panic("implement me")
}

func (i *impl) Create(ctx contextx.Contextx, item *agg.Issue) (err error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, timeoutDuration)
	defer cancelFunc()

	ownerID, err := primitive.ObjectIDFromHex(item.OwnerID)
	if err != nil {
		return err
	}

	now := time.Now()
	created := &issue{
		ID:        primitive.NewObjectIDFromTimestamp(now),
		Title:     item.Title,
		Completed: item.Completed,
		OwnerID:   ownerID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	coll := i.rw.Database(dbName).Collection(collName)
	_, err = coll.InsertOne(timeout, created)
	if err != nil {
		return err
	}

	item.ID = created.ID.Hex()
	return nil
}

func (i *impl) Update(ctx contextx.Contextx, item *agg.Issue) (err error) {
	// todo: 2024/3/10|sean|implement me
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id string) (err error) {
	// todo: 2024/3/10|sean|implement me
	panic("implement me")
}
