package todo

import (
	"time"

	"github.com/blackhorseya/ekko/app/infra/otelx"
	"github.com/blackhorseya/ekko/entity/domain/todo/model"
	"github.com/blackhorseya/ekko/entity/domain/todo/repo"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	defaultTimeout = 5 * time.Second
	dbName         = "ekko"
	collName       = "todos"
)

type mongodb struct {
	rw *mongo.Client
}

// NewMongodb is to create a new mongodb
func NewMongodb(rw *mongo.Client) repo.ITodoRepo {
	return &mongodb{rw: rw}
}

func (i *mongodb) List(
	ctx contextx.Contextx,
	condition repo.ListCondition,
) (items []*model.Todo, total int, err error) {
	ctx, span := otelx.StartSpan(ctx, "repo")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	filter := bson.M{}
	if condition.CreatedBy != "" {
		filter["created_by"] = condition.CreatedBy
	}

	opts := options.Find()
	if condition.Limit > 0 {
		opts.SetLimit(int64(condition.Limit))
	}
	if condition.Skip > 0 {
		opts.SetSkip(int64(condition.Skip))
	}
	opts.SetSort(bson.M{"updated_at": -1})

	coll := i.rw.Database(dbName).Collection(collName)
	cursor, err := coll.Find(timeout, filter, opts)
	if err != nil {
		ctx.Error("mongodb list todo error", zap.Error(err))
		return nil, 0, err
	}
	defer cursor.Close(timeout)

	err = cursor.All(timeout, &items)
	if err != nil {
		ctx.Error("mongodb list todo error", zap.Error(err))
		return nil, 0, err
	}

	count, err := coll.CountDocuments(timeout, filter)
	if err != nil {
		ctx.Error("mongodb list todo error", zap.Error(err))
		return nil, 0, err
	}

	return items, int(count), nil
}

func (i *mongodb) GetByID(ctx contextx.Contextx, id string) (item *model.Todo, err error) {
	ctx, span := otelx.StartSpan(ctx, "repo")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	coll := i.rw.Database(dbName).Collection(collName)
	err = coll.FindOne(timeout, bson.M{"_id": id}).Decode(&item)
	if err != nil {
		ctx.Error("mongodb get todo by id error", zap.Error(err), zap.String("id", id))
		return nil, err
	}

	return item, nil
}

func (i *mongodb) Create(ctx contextx.Contextx, item *model.Todo) (err error) {
	ctx, span := otelx.StartSpan(ctx, "repo")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	if item.ID == "" {
		item.ID = uuid.New().String()
	}
	item.UpdatedAt = time.Now()

	coll := i.rw.Database(dbName).Collection(collName)
	_, err = coll.InsertOne(timeout, item)
	if err != nil {
		ctx.Error("mongodb create todo error", zap.Error(err), zap.Any("item", &item))
		return err
	}

	return nil
}

func (i *mongodb) Update(ctx contextx.Contextx, item *model.Todo) (err error) {
	ctx, span := otelx.StartSpan(ctx, "repo")
	defer span.End()

	timeout, cancelFunc := contextx.WithTimeout(ctx, defaultTimeout)
	defer cancelFunc()

	item.UpdatedAt = time.Now()

	filter := bson.M{"_id": item.ID}
	update := bson.M{"$set": item}

	coll := i.rw.Database(dbName).Collection(collName)
	_, err = coll.UpdateOne(timeout, filter, update)
	if err != nil {
		ctx.Error("mongodb update todo error", zap.Error(err), zap.Any("item", &item))
		return err
	}

	return nil
}
