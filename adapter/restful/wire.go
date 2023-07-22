//go:build wireinject

//go:generate wire

package main

import (
	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/google/wire"
)

func NewConfig(path string) (*config.Config, error) {
	panic(wire.Build(config.NewWithPath))
}
