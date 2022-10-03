//go:build wireinject

package repo

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var testProviderSet = wire.NewSet(NewMariadb)

func CreateMariadb(rw *sqlx.DB) (IHealthRepo, error) {
	panic(wire.Build(testProviderSet))
}
