package apis

import (
	"github.com/blackhorseya/todo-app/internal/app/apis/health"
	"github.com/blackhorseya/todo-app/internal/app/apis/task"
	"github.com/google/wire"
)

// ProviderSet is an apis provider set
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	task.ProviderSet,
)
