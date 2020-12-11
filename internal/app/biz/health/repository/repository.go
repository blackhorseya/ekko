package repository

import (
	"time"

	"github.com/google/wire"
)

// HealthRepo is a repository to health Business
type HealthRepo interface {
	Ping(timeout time.Duration) error
}

// ProviderSet is a repository of health of provider set
var ProviderSet = wire.NewSet(NewImpl)
