//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/todo-app/internal/pkg/config"
	"github.com/blackhorseya/todo-app/internal/pkg/httpx"
	"github.com/blackhorseya/todo-app/internal/pkg/log"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// infra
	config.ProviderSet,
	log.ProviderSet,
	httpx.ProviderClientSet,

	// server
	httpx.ProviderServerSet,

	// implementation

	// main
	NewRestful,
	NewService,
)

func CreateService(path string) (*Service, error) {
	panic(wire.Build(providerSet))
}
