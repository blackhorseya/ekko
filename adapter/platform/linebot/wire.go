//go:build wireinject

//go:generate wire

package linebot

import (
	"github.com/blackhorseya/ekko/adapter/platform/wirex"
	"github.com/blackhorseya/ekko/app/domain/todo/biz"
	"github.com/blackhorseya/ekko/app/domain/todo/repo/todo"
	"github.com/blackhorseya/ekko/app/infra/authx"
	"github.com/blackhorseya/ekko/app/infra/configx"
	"github.com/blackhorseya/ekko/pkg/adapterx"
	"github.com/blackhorseya/ekko/pkg/linebotx"
	"github.com/blackhorseya/ekko/pkg/logging"
	"github.com/blackhorseya/ekko/pkg/storage/mongodbx"
	"github.com/blackhorseya/ekko/pkg/transports/httpx"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func initApplication() (*configx.Application, error) {
	app, err := configx.LoadApplication(&configx.C.PlatformLinebot)
	if err != nil {
		return nil, err
	}

	err = logging.InitWithConfig(app.Log)
	if err != nil {
		return nil, err
	}

	return app, nil
}

var providerSet = wire.NewSet(
	wire.Struct(new(wirex.Injector), "*"),
	initApplication,
	authx.NewNil,
	linebotx.NewClient,

	biz.NewTodoBiz,
	todo.NewMongodb,
	mongodbx.NewClient,

	httpx.NewServer,
)

func New(v *viper.Viper) (adapterx.Servicer, error) {
	panic(wire.Build(newService, providerSet))
}
