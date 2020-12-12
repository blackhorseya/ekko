package repository

import (
	"context"

	"github.com/blackhorseya/todo-app/internal/app/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type impl struct {
	MongoClient *mongo.Client
}

// NewImpl is a constructor task of implement repository
func NewImpl(mongoClient *mongo.Client) TaskRepo {
	return &impl{MongoClient: mongoClient}
}

// QueryTaskList handle query task list by limit and offset
func (i *impl) QueryTaskList(limit, offset int32) (tasks []*entities.Task, err error) {
	// todo: 2020-12-12|20:05|doggy|implement me
	panic("implement me")
}

// CreateTask handle create a task
func (i *impl) CreateTask(newTask *entities.Task) (task *entities.Task, err error) {
	coll := i.MongoClient.Database("todo-db").Collection("tasks")
	_, err = coll.InsertOne(context.TODO(), newTask)
	if err != nil {
		return newTask, err
	}

	return newTask, nil
}
