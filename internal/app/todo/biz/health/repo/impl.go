package repo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type impl struct {
	MongoClient *mongo.Client
}

// NewImpl is a constructor health of implement repo
func NewImpl(mongoClient *mongo.Client) IRepo {
	return &impl{MongoClient: mongoClient}
}

// Ping sends a ping command to verify that the client can connect to the deployment
func (i *impl) Ping(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err := i.MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	return nil
}
