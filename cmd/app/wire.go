// +build wireinject

package main

import (
	"github.com/blackhorseya/todo-app/internal/app"
	"github.com/blackhorseya/todo-app/internal/app/apis"
	"github.com/blackhorseya/todo-app/internal/app/biz"
	"github.com/blackhorseya/todo-app/internal/app/router"
	"github.com/blackhorseya/todo-app/internal/pkg/config"
	"github.com/blackhorseya/todo-app/internal/pkg/databases"
	"github.com/google/wire"
)

// CreateApp create an application
func CreateApp(cfg string) (*app.Injector, func(), error) {
	wire.Build(
		app.NewGinEngine,
		databases.NewMongo,
		config.ProviderSet,
		biz.ProviderSet,
		apis.ProviderSet,
		router.ProviderSet,
		app.InjectorSet,
	)

	return new(app.Injector), nil, nil
}
