package rest

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/blackhorseya/ekko/adapter/api/platform_rest" // swagger docs
	v1 "github.com/blackhorseya/ekko/adapter/platform/rest/v1"
	"github.com/blackhorseya/ekko/adapter/platform/wirex"
	"github.com/blackhorseya/ekko/app/infra/configx"
	"github.com/blackhorseya/ekko/pkg/adapterx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/otelx"
	"github.com/blackhorseya/ekko/pkg/responsex"
	"github.com/blackhorseya/ekko/pkg/transports/httpx"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title Ekko Platform Restful API
// @version 0.1.0
// @description Ekko Platform Restful API document.
//
// @contact.name Sean Zheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space
//
// @license.name GPL-3.0
// @license.url https://spdx.org/licenses/GPL-3.0-only.html
//
// @BasePath /api
//
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
type impl struct {
	injector *wirex.Injector
	server   *httpx.Server

	shutdown func(context.Context) error
}

func newRestful(injector *wirex.Injector, server *httpx.Server) adapterx.Restful {
	shutdown, err := otelx.SetupOTelSDK(contextx.Background(), injector.A.Name)
	if err != nil {
		log.Fatal(err)
	}

	return &impl{
		injector: injector,
		server:   server,
		shutdown: shutdown,
	}
}

func newService(injector *wirex.Injector, server *httpx.Server) adapterx.Servicer {
	return newRestful(injector, server)
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

	ctx.Info("start restful server", zap.String("swagger_url", fmt.Sprintf(
		"http://%s/api/docs/index.html",
		strings.ReplaceAll(configx.A.HTTP.GetAddr(), "0.0.0.0", "localhost"),
	)))

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
		}

		err = i.shutdown(ctx)
		if err != nil {
			ctx.Error("shutdown otel sdk error", zap.Error(err))
		}
	}

	return nil
}

func (i *impl) InitRouting() error {
	router := i.server.Router

	// api
	api := router.Group("/api")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.InstanceName("platform_rest"),
		))
		api.GET("/healthz", i.Healthz)

		v1.Handle(api.Group("/v1", i.injector.Authx.ParseJWT()), i.injector)
	}

	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.server.Router
}

// Healthz is used to check the health of the service.
// @Summary Check the health of the service.
// @Description Check the health of the service.
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} responsex.Response
// @Failure 500 {object} responsex.Response
// @Router /healthz [get]
func (i *impl) Healthz(c *gin.Context) {
	responsex.OK(c, nil)
}
