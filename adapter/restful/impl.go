package restful

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/ekko/entity/domain/workflow/biz"
	"github.com/blackhorseya/ekko/pkg/adapterx"
	"github.com/blackhorseya/ekko/pkg/configx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/response"
	"github.com/blackhorseya/ekko/pkg/transports/httpx"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"go.uber.org/zap"
)

type impl struct {
	server   *httpx.Server
	bot      *messaging_api.MessagingApiAPI
	workflow biz.IWorkflowBiz
}

func newRestful(server *httpx.Server, bot *messaging_api.MessagingApiAPI, workflow biz.IWorkflowBiz) adapterx.Restful {
	return &impl{
		server:   server,
		bot:      bot,
		workflow: workflow,
	}
}

func newService(server *httpx.Server, bot *messaging_api.MessagingApiAPI, workflow biz.IWorkflowBiz) adapterx.Servicer {
	return newRestful(
		server,
		bot,
		workflow,
	)
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

	ctx.Info("start restful server", zap.String("addr", configx.C.HTTP.GetAddr()))

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
	api := i.server.Router.Group("/api")
	{
		api.POST("/callback", i.callback)
	}

	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.server.Router
}

func (i *impl) callback(c *gin.Context) {
	// todo: 2024/3/10|sean|implement callback logic
	c.JSON(http.StatusOK, response.OK)
}
