package wirex

import (
	"github.com/blackhorseya/ekko/app/infra/configx"
	"github.com/blackhorseya/ekko/entity/domain/todo/biz"
)

// Injector is used to inject market data.
type Injector struct {
	A    *configx.Application
	Todo biz.ITodoBiz
}
