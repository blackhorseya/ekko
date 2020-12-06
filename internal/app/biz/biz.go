package biz

import (
	"github.com/blackhorseya/todo-app/internal/app/biz/health"
	"github.com/blackhorseya/todo-app/internal/app/biz/task"
	"github.com/google/wire"
)

// ProviderSet is a business provider set
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	task.ProviderSet,
)
