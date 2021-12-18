package repo

import (
	"time"

	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName = "todo-db"

	collName = "tasks"
)

type impl struct {
	client *mongo.Client
}

// NewImpl serve caller to create an IRepo
func NewImpl(client *mongo.Client) IRepo {
	return &impl{client: client}
}

func (i *impl) GetByID(ctx contextx.Contextx, id int64) (task *pb.Task, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.client.Database(dbName).Collection(collName)
	res := coll.FindOne(timeout, bson.D{{"id", id}})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, res.Err()
	}

	var ret *pb.Task
	err = res.Decode(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, limit, offset int) (tasks []*pb.Task, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.client.Database(dbName).Collection(collName)
	cur, err := coll.Find(timeout, bson.D{}, options.Find().SetLimit(int64(limit)).SetSkip(int64(offset)))
	if err != nil {
		return nil, err
	}
	defer cur.Close(timeout)

	var ret []*pb.Task
	for cur.Next(timeout) {
		var task *pb.Task
		err := cur.Decode(&task)
		if err != nil {
			continue
		}

		ret = append(ret, task)
	}

	return ret, nil
}

func (i *impl) Count(ctx contextx.Contextx) (total int, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.client.Database("todo-db").Collection("tasks")
	ret, err := coll.CountDocuments(timeout, bson.M{})
	if err != nil {
		return 0, err
	}

	return int(ret), nil
}

func (i *impl) Create(ctx contextx.Contextx, newTask *pb.Task) (task *pb.Task, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.client.Database(dbName).Collection(collName)
	_, err = coll.InsertOne(timeout, newTask)
	if err != nil {
		return nil, err
	}

	return newTask, nil
}

func (i *impl) Update(ctx contextx.Contextx, updated *pb.Task) (task *pb.Task, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.client.Database(dbName).Collection(collName)
	res := coll.FindOneAndUpdate(timeout, bson.D{{"id", updated.Id}}, bson.D{{"$set", updated}})
	if res.Err() != nil {
		return nil, res.Err()
	}

	return updated, nil
}

func (i *impl) Remove(ctx contextx.Contextx, id int64) error {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.client.Database(dbName).Collection(collName)
	_, err := coll.DeleteOne(timeout, bson.D{{"id", id}})
	if err != nil {
		return err
	}

	return nil
}
