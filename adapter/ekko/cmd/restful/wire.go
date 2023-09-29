//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/ekko/internal/app/domain/task/biz"
	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/blackhorseya/ekko/internal/pkg/httpx"
	"github.com/blackhorseya/ekko/internal/pkg/storage/mariadb"
	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var providerSet = wire.NewSet(
	httpx.ServerSet,
	newRestful,

	mariadb.NewMariadb,
	biz.TaskBizSet,
)

func NewService(config *config.Config, logger *zap.Logger) (adapters.Restful, error) {
	panic(wire.Build(providerSet))
}
