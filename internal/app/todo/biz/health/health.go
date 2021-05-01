package health

import (
	repository2 "github.com/blackhorseya/todo-app/internal/app/todo/biz/health/repository"
	"github.com/google/wire"
)

// Biz describe health business service function
type Biz interface {
	Readiness() (ok bool, err error)
	Liveness() (ok bool, err error)
}

// ProviderSet is a health provider set
var ProviderSet = wire.NewSet(NewImpl, repository2.ProviderSet)
