//go:build wireinject

//go:generate wire

package rest

import (
	"github.com/blackhorseya/ekko/adapter/platform/wirex"
	"github.com/blackhorseya/ekko/app/infra/configx"
	"github.com/blackhorseya/ekko/pkg/logging"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func initApplication() (*configx.Application, error) {
	app, err := configx.LoadApplication(&configx.C.PlatformRest)
	if err != nil {
		return nil, err
	}

	err = logging.InitWithConfig(app.Log)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func New(v *viper.Viper) (wirex.Injector, error) {
	panic(wire.Build(
		wire.Struct(new(wirex.Injector), "*"),
		initApplication,
	))
}
