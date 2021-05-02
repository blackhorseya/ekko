package repo

import (
	"time"

	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type impl struct {
	client *mongo.Client
}

// NewImpl serve caller to create an IRepo
func NewImpl(client *mongo.Client) IRepo {
	return &impl{client: client}
}

func (i *impl) GetByID(ctx contextx.Contextx, id string) (task *todo.Task, err error) {
	// todo: 2021-05-02|10:15|doggy|implement me
	panic("implement me")
}

func (i *impl) List(ctx contextx.Contextx, limit, offset int) (tasks []*todo.Task, err error) {
	// todo: 2021-05-02|10:15|doggy|implement me
	panic("implement me")
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

func (i *impl) Create(ctx contextx.Contextx, newTask *todo.Task) (task *todo.Task, err error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.client.Database("todo-db").Collection("tasks")
	_, err = coll.InsertOne(timeout, newTask)
	if err != nil {
		return nil, err
	}

	return newTask, nil
}

func (i *impl) Update(ctx contextx.Contextx, updated *todo.Task) (task *todo.Task, err error) {
	// todo: 2021-05-02|10:15|doggy|implement me
	panic("implement me")
}

func (i *impl) Remove(ctx contextx.Contextx, id string) error {
	// todo: 2021-05-02|10:15|doggy|implement me
	panic("implement me")
}
