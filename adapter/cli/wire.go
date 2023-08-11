//go:build wireinject

//go:generate wire

package main

import (
	"github.com/blackhorseya/ekko/adapter/cli/app"
	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/blackhorseya/ekko/internal/pkg/log"
	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func NewConfig(path string) (*config.Config, error) {
	panic(wire.Build(config.NewWithPath))
}

func NewLogger(config *config.Config) (*zap.Logger, error) {
	panic(wire.Build(log.NewLogger))
}

func NewCmd(config config.Config, logger *zap.Logger) (adapters.CLI, error) {
	panic(wire.Build(app.NewCmd))
}
