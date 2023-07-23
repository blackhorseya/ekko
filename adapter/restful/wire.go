//go:build wireinject

//go:generate wire

package main

import (
	"github.com/blackhorseya/ekko/adapter/restful/app"
	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/blackhorseya/ekko/internal/pkg/log"
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
)

func NewService(config *config.Config, logger *zap.Logger) (*app.Service, error) {
	panic(wire.Build(providerSet))
}
