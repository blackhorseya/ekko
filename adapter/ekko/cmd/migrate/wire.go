//go:build wireinject

//go:generate wire

package migrate

import (
	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/blackhorseya/ekko/internal/pkg/storage/mariadb"
	"github.com/golang-migrate/migrate/v4"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var providerSet = wire.NewSet(
	mariadb.NewMariadb,
	mariadb.NewMigration,
)

func NewMigration(config *config.Config, logger *zap.Logger) (*migrate.Migrate, error) {
	panic(wire.Build(providerSet))
}
