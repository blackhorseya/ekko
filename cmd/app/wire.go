// +build wireinject

package main

import (
	"github.com/blackhorseya/todo-app/internal/app"
	"github.com/blackhorseya/todo-app/internal/app/apis"
	"github.com/blackhorseya/todo-app/internal/app/router"
	"github.com/google/wire"
)

// BuildInjector build a injector
func BuildInjector() (*app.Injector, func(), error) {
	wire.Build(
		app.NewGinEngine,
		apis.ProviderSet,
		router.Set,
		app.InjectorSet,
	)

	return new(app.Injector), nil, nil
}
