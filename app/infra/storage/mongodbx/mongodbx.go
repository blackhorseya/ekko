package mongodbx

import (
	"github.com/blackhorseya/ekko/app/infra/configx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewClientWithDSN returns a new mongo client with dsn.
func NewClientWithDSN(dsn string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(contextx.Background(), opts)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewClient returns a new mongo client.
func NewClient() (*mongo.Client, error) {
	return NewClientWithDSN(configx.A.Storage.Mongodb.DSN)
}

// Container is used to represent a mongodb container.
type Container struct {
	*mongodb.MongoDBContainer
}

// NewContainer returns a new mongodb container.
func NewContainer(ctx contextx.Contextx) (*Container, error) {
	container, err := mongodb.RunContainer(ctx, testcontainers.WithImage("mongo:6"))
	if err != nil {
		return nil, errors.Wrap(err, "run mongodb container")
	}

	return &Container{
		MongoDBContainer: container,
	}, nil
}
