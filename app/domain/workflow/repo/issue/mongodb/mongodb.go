package mongodb

import (
	"time"

	"github.com/blackhorseya/ekko/entity/domain/workflow/agg"
	"github.com/blackhorseya/ekko/entity/domain/workflow/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (i *impl) List(ctx contextx.Contextx, cond repo.ListIssueOptions) (items []*agg.Issue, total int, err error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, timeoutDuration)
	defer cancelFunc()

	filter := bson.D{}
	if cond.OwnerID != "" {
		filter = append(filter, bson.E{Key: "owner_id", Value: cond.OwnerID})
	}

	opts := options.Find()
	if cond.Limit > 0 {
		opts.SetLimit(int64(cond.Limit))
	}

	if cond.Offset > 0 {
		opts.SetSkip(int64(cond.Offset))
	}

	coll := i.rw.Database(dbName).Collection(collName)
	cur, err := coll.Find(timeout, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(timeout)

	var ret []*agg.Issue
	for cur.Next(timeout) {
		var got *issue
		err = cur.Decode(&got)
		if err != nil {
			return nil, 0, err
		}

		ret = append(ret, got.ToAgg())
	}

	return ret, 0, nil
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (item *agg.Issue, err error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, timeoutDuration)
	defer cancelFunc()

	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": hex}
	coll := i.rw.Database(dbName).Collection(collName)

	var got *issue
	err = coll.FindOne(timeout, filter).Decode(&got)
	if err != nil {
		return nil, err
	}

	return got.ToAgg(), nil
}

func (i *impl) Create(ctx contextx.Contextx, item *agg.Issue) (err error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, timeoutDuration)
	defer cancelFunc()

	now := time.Now()
	created := &issue{
		ID:        primitive.NewObjectIDFromTimestamp(now),
		Title:     item.Title,
		Completed: item.Completed,
		OwnerID:   item.OwnerID,
		// CreatedAt: now,
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
	timeout, cancelFunc := contextx.WithTimeout(ctx, timeoutDuration)
	defer cancelFunc()

	item.UpdatedAt = time.Now()

	update := newIssue(item)
	filter := bson.M{"_id": update.ID}
	opts := options.FindOneAndReplace().SetReturnDocument(options.After)
	coll := i.rw.Database(dbName).Collection(collName)
	err = coll.FindOneAndReplace(timeout, filter, update, opts).Err()
	if err != nil {
		return err
	}

	return nil
}

func (i *impl) Delete(ctx contextx.Contextx, id string) (err error) {
	// todo: 2024/3/10|sean|implement me
	panic("implement me")
}
