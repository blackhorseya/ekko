//go:build wireinject

//go:generate wire

package restful

import (
	"github.com/blackhorseya/ekko/pkg/adapterx"
	"github.com/blackhorseya/ekko/pkg/transports/httpx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(
		httpx.NewServer,
		newService,
	))
}

func NewRestful(v *viper.Viper) (adapterx.Restful, error) {
	panic(wire.Build(
		httpx.NewServer,
		newRestful,
	))
}
