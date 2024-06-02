package wirex

import (
	"github.com/blackhorseya/ekko/app/infra/configx"
)

// Injector is used to inject market data.
type Injector struct {
	A *configx.Application
}
