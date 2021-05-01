package health

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health/repo"
	"github.com/google/wire"
)

// IBiz describe health business service function
type IBiz interface {
	Readiness() (ok bool, err error)
	Liveness() (ok bool, err error)
}

// ProviderSet is a health provider set
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)
