// +build wireinject

package repo

import (
	"github.com/blackhorseya/todo-app/internal/pkg/entity/config"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/database"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/log"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	NewImpl,
)

// CreateIRepo serve caller to create an IRepo
func CreateIRepo(path string) (IRepo, error) {
	panic(wire.Build(testProviderSet))
}
