package biz

import (
	health2 "github.com/blackhorseya/todo-app/internal/app/todo/biz/health"
	"github.com/google/wire"
)

// ProviderSet is a business provider set
var ProviderSet = wire.NewSet(
	health2.ProviderSet,
)
