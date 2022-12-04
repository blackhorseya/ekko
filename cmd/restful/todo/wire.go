//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/todo-app/internal/app/domain/todo/biz"
	"github.com/blackhorseya/todo-app/internal/pkg/config"
	"github.com/blackhorseya/todo-app/internal/pkg/genx"
	"github.com/blackhorseya/todo-app/internal/pkg/httpx"
	"github.com/blackhorseya/todo-app/internal/pkg/log"
	"github.com/blackhorseya/todo-app/internal/pkg/storage/mariadb"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// infra
	config.ProviderSet,
	log.ProviderSet,
	genx.ProviderSet,

	// storage
	mariadb.ProviderSet,

	// transports
	httpx.ProviderClientSet,
	httpx.ProviderServerSet,

	// implementation
	biz.ProviderStorageSet,

	// main
	NewRestful,
	NewService,
)

func CreateService(path string, id int64) (*Service, error) {
	panic(wire.Build(providerSet))
}
