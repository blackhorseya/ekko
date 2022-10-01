//go:build wireinject

package repo

import (
	"github.com/blackhorseya/gocommon/pkg/restclient"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateIRepo(client *mongo.Client) (IRepo, error) {
	panic(wire.Build(testProviderSet))
}

var httpProviderSet = wire.NewSet(NewHTTP)

func CreateHTTP(opts *Options, client restclient.RestClient) (IRepo, error) {
	panic(wire.Build(httpProviderSet))
}
