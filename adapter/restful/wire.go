//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/ekko/app/domain/workflow/biz"
	"github.com/blackhorseya/ekko/pkg/adapterx"
	"github.com/blackhorseya/ekko/pkg/linebotx"
	"github.com/blackhorseya/ekko/pkg/transports/httpx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var providerSet = wire.NewSet(
	httpx.NewServer,
	linebotx.NewClient,
	biz.DefaultWorkflowSet,
)

func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(
		newService,
		providerSet,
	))
}

func NewRestful(v *viper.Viper) (adapterx.Restful, error) {
	panic(wire.Build(
		newRestful,
		providerSet,
	))
}
