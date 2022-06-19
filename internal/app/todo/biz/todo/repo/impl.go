package repo

import (
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	dbName = "todo-db"

	collName = "tasks"
)

type impl struct {
	logger *zap.Logger
	client *mongo.Client
}

// NewImpl serve caller to create an IRepo
func NewImpl(logger *zap.Logger, client *mongo.Client) IRepo {
	return &impl{
		logger: logger.With(zap.String("type", "TodoRepo")),
		client: client,
	}
}

func (i *impl) GetByID(ctx contextx.Contextx, id primitive.ObjectID) (task *todo.Task, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	var ret todo.Task
	coll := i.client.Database(dbName).Collection(collName)
	err = coll.FindOne(timeout, filter).Decode(&ret)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &ret, nil
}

func (i *impl) List(ctx contextx.Contextx, limit, offset int) (tasks []*todo.Task, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.client.Database(dbName).Collection(collName)
	cur, err := coll.Find(timeout, bson.D{}, options.Find().SetLimit(int64(limit)).SetSkip(int64(offset)))
	if err != nil {
		return nil, err
	}
	defer cur.Close(timeout)

	var ret []*todo.Task
	err = cur.All(timeout, &ret)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return ret, nil
}

func (i *impl) Count(ctx contextx.Contextx) (total int, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.client.Database(dbName).Collection(collName)
	ret, err := coll.CountDocuments(timeout, bson.M{})
	if err != nil {
		return 0, err
	}

	return int(ret), nil
}

func (i *impl) Create(ctx contextx.Contextx, newTask *todo.Task) (task *todo.Task, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	now := time.Now()
	newTask.ID = primitive.NewObjectIDFromTimestamp(now)
	newTask.CreatedAt = primitive.NewDateTimeFromTime(now)
	newTask.UpdatedAt = primitive.NewDateTimeFromTime(now)

	coll := i.client.Database(dbName).Collection(collName)
	res, err := coll.InsertOne(timeout, newTask)
	if err != nil {
		return nil, err
	}

	newTask.ID = res.InsertedID.(primitive.ObjectID)

	return newTask, nil
}

func (i *impl) Update(ctx contextx.Contextx, updated *todo.Task) (task *todo.Task, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	updated.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	filter := bson.M{"_id": updated.ID}
	opt := options.FindOneAndReplace().SetUpsert(false).SetReturnDocument(options.After)
	coll := i.client.Database(dbName).Collection(collName)
	var ret *todo.Task
	err = coll.FindOneAndReplace(timeout, filter, updated, opt).Decode(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *impl) Remove(ctx contextx.Contextx, id primitive.ObjectID) error {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.client.Database(dbName).Collection(collName)
	_, err := coll.DeleteOne(timeout, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
