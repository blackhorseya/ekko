//go:build wireinject

//go:generate wire

package main

import (
	"github.com/blackhorseya/ekko/adapter/cli/app"
	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/google/wire"
)

func NewCmd() (adapters.CLI, error) {
	panic(wire.Build(app.NewCmd))
}
