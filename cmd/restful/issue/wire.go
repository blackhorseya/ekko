//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/ekko/internal/adapter/issue/restful"
	"github.com/blackhorseya/ekko/internal/app/domain/issue/biz"
	"github.com/blackhorseya/ekko/internal/app/domain/issue/biz/repo"
	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/blackhorseya/ekko/internal/pkg/genx"
	"github.com/blackhorseya/ekko/internal/pkg/httpx"
	"github.com/blackhorseya/ekko/internal/pkg/log"
	"github.com/blackhorseya/ekko/internal/pkg/storage/mariadb"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// infra
	config.NewWithPath,
	log.NewLogger,
	genx.ProviderSet,

	// storage
	mariadb.ProviderSet,

	// transports
	httpx.ServerSet,

	// implementation
	restful.IssueSet,
	biz.IssueSet,
	repo.ProvideMariadb,

	// main
	NewService,
)

func CreateService(path string, id int64) (*Service, error) {
	panic(wire.Build(providerSet))
}
