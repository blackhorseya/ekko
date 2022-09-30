package repo

import (
	"time"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/google/wire"
)

// IRepo is a repo to health Business
//
//go:generate mockery --all --inpackage
type IRepo interface {
	Ping(ctx contextx.Contextx, timeout time.Duration) error
}

// ProviderSet is a repo of health of provider set
var ProviderSet = wire.NewSet(NewImpl)
