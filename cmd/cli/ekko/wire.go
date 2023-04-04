//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/ekko/internal/adapter/cli"
	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/blackhorseya/ekko/internal/pkg/log"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// infrastructure
	config.ProviderSet,
	log.ProviderSet,

	// adapters
	cli.ProviderSet,

	// main
	NewService,
)

// InitializeService serve caller to initialize service
func InitializeService(path string) (*Service, error) {
	panic(wire.Build(providerSet))
}
