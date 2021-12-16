// +build wireinject

package main

import (
	"github.com/blackhorseya/todo-app/internal/app/todo"
	"github.com/blackhorseya/todo-app/internal/app/todo/apis"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz"
	"github.com/blackhorseya/todo-app/internal/pkg/app"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/config"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/database"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/generator"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/log"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	generator.ProviderSet,
	http.ProviderSet,
	database.ProviderSet,
	todo.ProviderSet,
	apis.ProviderSet,
	biz.ProviderSet,
)

// CreateApp serve caller to create an injector
func CreateApp(path string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
