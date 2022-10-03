//go:build wireinject

package repo

import (
	"github.com/blackhorseya/gocommon/pkg/restclient"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var mariadbProviderSet = wire.NewSet(NewMariadb)

func CreateMariadb(rw *sqlx.DB) (ITodoRepo, error) {
	panic(wire.Build(mariadbProviderSet))
}

var httpProviderSet = wire.NewSet(NewHTTP)

func CreateHTTP(opts *Options, client restclient.RestClient) (ITodoRepo, error) {
	panic(wire.Build(httpProviderSet))
}
