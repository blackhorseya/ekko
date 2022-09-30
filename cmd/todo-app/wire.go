//go:build wireinject

package main

import (
	"github.com/blackhorseya/gocommon/pkg/config"
	"github.com/blackhorseya/gocommon/pkg/log"
	"github.com/blackhorseya/todo-app/internal/app/todo"
	"github.com/blackhorseya/todo-app/internal/app/todo/api/restful"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz"
	"github.com/blackhorseya/todo-app/internal/pkg/app"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/database"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	http.ProviderSet,
	database.ProviderSet,
	todo.ProviderSet,
	restful.ProviderSet,
	biz.ProviderSet,
)

// CreateApp serve caller to create an injector
func CreateApp(path string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
