package linebot

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/ekko/adapter/platform/wirex"
	"github.com/blackhorseya/ekko/pkg/adapterx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/responsex"
	"github.com/blackhorseya/ekko/pkg/transports/httpx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	injector *wirex.Injector
	server   *httpx.Server
}

func newRest(injector *wirex.Injector, server *httpx.Server) adapterx.Restful {
	return &impl{
		injector: injector,
		server:   server,
	}
}

func newService(injector *wirex.Injector, server *httpx.Server) adapterx.Servicer {
	return newRest(injector, server)
}

func (i *impl) Start() error {
	ctx := contextx.Background()

	err := i.InitRouting()
	if err != nil {
		return err
	}

	err = i.server.Start(ctx)
	if err != nil {
		return err
	}

	ctx.Info("start server", zap.String("address", i.injector.A.HTTP.GetAddr()))

	return nil
}

func (i *impl) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		ctx := contextx.Background()
		ctx.Info("receive signal", zap.String("signal", sig.String()))

		err := i.server.Stop(ctx)
		if err != nil {
			ctx.Error("shutdown restful server error", zap.Error(err))
			return err
		}
	}

	return nil
}

func (i *impl) InitRouting() error {
	router := i.server.Router

	// api
	api := router.Group("/api")
	{
		api.POST("/callback", i.callback)
	}

	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.server.Router
}

func (i *impl) callback(c *gin.Context) {
	// todo: 2024/6/4|sean|implement line bot callback
	responsex.OK(c, nil)
}
