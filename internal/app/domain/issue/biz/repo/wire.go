//go:generate wire
//go:build wireinject

package repo

import (
	"github.com/blackhorseya/ekko/pkg/httpx"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var testProviderMariadbSet = wire.NewSet(NewMariadb)

func CreateMariadb(rw *sqlx.DB) IRepo {
	panic(wire.Build(testProviderMariadbSet))
}

var testProviderHTTPClientSet = wire.NewSet(NewHTTPClient)

func CreateHTTPClient(v *viper.Viper, client httpx.Client) (IRepo, error) {
	panic(wire.Build(testProviderHTTPClientSet))
}
