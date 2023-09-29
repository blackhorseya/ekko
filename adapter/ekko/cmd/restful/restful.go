package restful

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/blackhorseya/ekko/adapter/ekko/api/docs" // import swagger docs
	v1 "github.com/blackhorseya/ekko/adapter/ekko/cmd/restful/v1"
	taskB "github.com/blackhorseya/ekko/entity/domain/task/biz"
	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/cors"
	"github.com/blackhorseya/ekko/pkg/er"
	"github.com/blackhorseya/ekko/pkg/httpx"
	"github.com/blackhorseya/ekko/pkg/response"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger

	server httpx.Server
	router *gin.Engine

	task taskB.IBiz
}

func newRestful(logger *zap.Logger, server httpx.Server, router *gin.Engine, task taskB.IBiz) adapters.Restful {
	return &impl{
		logger: logger.With(zap.String("type", "restful")),
		server: server,
		router: router,
		task:   task,
	}
}

func (i *impl) InitRouting() {
	logger := i.logger
	router := i.router

	router.Use(cors.AddAllowAll())
	router.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/api/healthz"},
		TraceID:    false,
		Context:    nil,
	}))
	router.Use(ginzap.CustomRecoveryWithZap(logger, true, func(c *gin.Context, err any) {
		msg := fmt.Sprintf("%v", err)
		resp := er.New(http.StatusInternalServerError, msg)
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
	}))
	router.Use(contextx.WithContextx(logger))
	router.Use(er.HandleError())

	apiG := i.router.Group("/api")
	{
		apiG.GET("/healthz", i.healthz)
		apiG.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		v1.Handle(apiG.Group("/v1"), i.task)
	}
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

// healthz is a handler for health check
// @Summary Health check
// @Description Health check
// @Tags Health
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /healthz [get]
func (i *impl) healthz(c *gin.Context) {
	c.JSON(http.StatusOK, response.OK)
}
