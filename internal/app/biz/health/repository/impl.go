package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type impl struct {
	MongoClient *mongo.Client
}

func NewImpl(mongoClient *mongo.Client) HealthRepo {
	return &impl{MongoClient: mongoClient}
}

func (i *impl) Ping(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err := i.MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	return nil
}
