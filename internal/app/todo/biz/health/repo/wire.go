//go:build wireinject
// +build wireinject

package repo

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateIRepo(logger *zap.Logger, client *mongo.Client) (IRepo, error) {
	panic(wire.Build(testProviderSet))
}
