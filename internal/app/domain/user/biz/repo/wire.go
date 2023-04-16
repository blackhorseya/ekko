//go:generate wire
//go:build wireinject

package repo

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var testMariadbSet = wire.NewSet(NewMariadb)

func NewRepoByMariadb(rw *sqlx.DB) IRepo {
	panic(wire.Build(testMariadbSet))
}
