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
