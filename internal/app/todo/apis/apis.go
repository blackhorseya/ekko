package apis

import (
	health2 "github.com/blackhorseya/todo-app/internal/app/todo/apis/health"
	"github.com/google/wire"
)

// ProviderSet is an apis provider set
var ProviderSet = wire.NewSet(
	health2.ProviderSet,
)
