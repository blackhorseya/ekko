package health

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type impl struct {
	MongoClient *mongo.Client
}

// NewImpl is a constructor of implement business with parameters
func NewImpl(mongoClient *mongo.Client) Biz {
	return &impl{MongoClient: mongoClient}
}

// Readiness to handle application has been ready
func (i *impl) Readiness() (ok bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = i.MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return false, err
	}

	return true, nil
}

// Liveness to handle application was alive
func (i *impl) Liveness() (ok bool, err error) {
	return true, nil
}
