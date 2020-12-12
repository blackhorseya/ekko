// +build wireinject

package repository

import (
	"github.com/blackhorseya/todo-app/internal/pkg/config"
	"github.com/blackhorseya/todo-app/internal/pkg/databases"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	databases.NewMongo,
	config.ProviderSet,
	NewImpl,
)

// CreateTaskRepo serve user to create task repo
func CreateTaskRepo(cfg string) (TaskRepo, error) {
	panic(wire.Build(testProviderSet))
}
