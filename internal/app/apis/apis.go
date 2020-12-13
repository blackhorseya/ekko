package apis

import (
	"github.com/blackhorseya/todo-app/internal/app/apis/task"
	"github.com/google/wire"
)

// ProviderSet is an apis provider set
var ProviderSet = wire.NewSet(
	HealthSet,
	task.ProviderSet,
)
