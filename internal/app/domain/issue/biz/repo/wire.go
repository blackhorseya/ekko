//go:generate wire
//go:build wireinject

package repo

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var testProviderMariadbSet = wire.NewSet(NewMariadb)

func CreateMariadb(rw *sqlx.DB) IRepo {
	panic(wire.Build(testProviderMariadbSet))
}

var testProviderHTTPClientSet = wire.NewSet(NewHTTPClient)

func CreateHTTPClient(opts *HTTPClientOptions) IRepo {
	panic(wire.Build(testProviderHTTPClientSet))
}