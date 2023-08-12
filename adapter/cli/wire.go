//go:build wireinject

//go:generate wire

package main

import (
	"github.com/blackhorseya/ekko/adapter/cli/app"
	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/google/wire"
)

func NewConfig(path string) (*config.Config, error) {
	panic(wire.Build(config.NewWithPath))
}

func NewCmd(config config.Config) (adapters.CLI, error) {
	panic(wire.Build(app.NewCmd))
}
