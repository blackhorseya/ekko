// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/blackhorseya/ekko/adapter/cli/app"
	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/blackhorseya/ekko/pkg/adapters"
)

// Injectors from wire.go:

func NewConfig(path2 string) (*config.Config, error) {
	configConfig, err := config.NewWithPath(path2)
	if err != nil {
		return nil, err
	}
	return configConfig, nil
}

func NewCmd(config2 *config.Config) (adapters.CLI, error) {
	cli := app.NewCmd(config2)
	return cli, nil
}