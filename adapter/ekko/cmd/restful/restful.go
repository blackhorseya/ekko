package restful

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/blackhorseya/ekko/pkg/httpx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger

	server httpx.Server
	router *gin.Engine
}

func newRestful(logger *zap.Logger, server httpx.Server, router *gin.Engine) adapters.Restful {
	return &impl{
		logger: logger.With(zap.String("type", "restful")),
		server: server,
		router: router,
	}
}

func (i *impl) InitRouting() {
	// todo: 2023/9/30|sean|impl me
	panic("implement me")
}

func (i *impl) Start() error {
	i.logger.Info("start restful server")

	i.InitRouting()

	err := i.server.Start()
	if err != nil {
		return err
	}

	return nil
}

func (i *impl) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		i.logger.Info("receive signal", zap.String("signal", sig.String()))

		err := i.server.Stop()
		if err != nil {
			i.logger.Warn("stop restful server", zap.Error(err))
		}

		os.Exit(0)
	}

	return nil
}
