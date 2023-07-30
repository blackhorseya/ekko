//go:build wireinject

//go:generate wire

package main

import (
	"github.com/blackhorseya/ekko/adapter/restful/app"
	"github.com/blackhorseya/ekko/internal/app/domain/issue/biz"
	"github.com/blackhorseya/ekko/internal/app/domain/issue/biz/repo"
	taskB "github.com/blackhorseya/ekko/internal/app/domain/task/biz"
	taskR "github.com/blackhorseya/ekko/internal/app/domain/task/biz/repo"
	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/blackhorseya/ekko/internal/pkg/genx"
	"github.com/blackhorseya/ekko/internal/pkg/log"
	"github.com/blackhorseya/ekko/internal/pkg/storage/mariadb"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func NewConfig(path string) (*config.Config, error) {
	panic(wire.Build(config.NewWithPath))
}

func NewLogger(config *config.Config) (*zap.Logger, error) {
	panic(wire.Build(log.NewLogger))
}

var providerSet = wire.NewSet(
	app.ProviderSet,

	mariadb.NewMariadb,
	genx.NewGenerator,

	// biz
	biz.IssueSet,
	taskB.NewImpl,

	// repo
	repo.MariadbSet,
	taskR.NewMariadb,
)

func NewService(config *config.Config, logger *zap.Logger, id int64) (*app.Service, error) {
	panic(wire.Build(providerSet))
}
