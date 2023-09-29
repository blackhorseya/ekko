package restful

import (
	"github.com/blackhorseya/ekko/pkg/adapters"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
}

func newRestful(logger *zap.Logger) adapters.Restful {
	return &impl{
		logger: logger.With(zap.String("type", "restful")),
	}
}

func (i *impl) InitRouting() {
	// todo: 2023/9/30|sean|impl me
	panic("implement me")
}

func (i *impl) Start() error {
	// todo: 2023/9/30|sean|impl me
	panic("implement me")
}

func (i *impl) AwaitSignal() error {
	// todo: 2023/9/30|sean|impl me
	panic("implement me")
}
