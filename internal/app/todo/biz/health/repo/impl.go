package repo

import (
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
)

type impl struct {
	MongoClient *mongo.Client
}

// NewImpl is a constructor health of implement repo
func NewImpl() IRepo {
	return &impl{MongoClient: nil}
}

// Ping sends a ping command to verify that the client can connect to the deployment
func (i *impl) Ping(ctx contextx.Contextx, timeout time.Duration) error {
	// withTimeout, cancel := contextx.WithTimeout(ctx, timeout)
	// defer cancel()
	//
	// err := i.MongoClient.Ping(withTimeout, readpref.Primary())
	// if err != nil {
	// 	return err
	// }
	//
	// return nil

	// todo: 2022/10/4|sean|impl me
	panic("impl me")
}
