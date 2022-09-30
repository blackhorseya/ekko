//go:build wireinject

package repo

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateIRepo(client *mongo.Client) (IRepo, error) {
	panic(wire.Build(testProviderSet))
}
