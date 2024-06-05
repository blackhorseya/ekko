package restful

import (
	"encoding/gob"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/ekko/adapter/restful/templates"
	v1 "github.com/blackhorseya/ekko/adapter/restful/v1"
	"github.com/blackhorseya/ekko/app/infra/configx"
	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/entity/domain/workflow/biz"
	"github.com/blackhorseya/ekko/pkg/adapterx"
	"github.com/blackhorseya/ekko/pkg/authx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/transports/httpx"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type impl struct {
	server        *httpx.Server
	bot           *messaging_api.MessagingApiAPI
	authenticator *authx.Authenticator
	workflow      biz.IWorkflowBiz
}

func newRestful(
	server *httpx.Server,
	bot *messaging_api.MessagingApiAPI,
	authenticator *authx.Authenticator,
	workflow biz.IWorkflowBiz,
) adapterx.Restful {
	return &impl{
		server:        server,
		bot:           bot,
		authenticator: authenticator,
		workflow:      workflow,
	}
}

func newService(
	server *httpx.Server,
	bot *messaging_api.MessagingApiAPI,
	authenticator *authx.Authenticator,
	workflow biz.IWorkflowBiz,
) adapterx.Servicer {
	return newRestful(
		server,
		bot,
		authenticator,
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

	ctx.Info("start restful server", zap.String("addr", configx.A.HTTP.GetAddr()))

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
	gob.Register(map[string]any{})
	store := cookie.NewStore([]byte("secret"))
	i.server.Router.Use(sessions.Sessions("auth-session", store))

	templates.SetHTMLTemplate(i.server.Router)

	i.server.Router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]any{
			"title": "Welcome to Ekko",
		})
	})

	api := i.server.Router.Group("/api")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.POST("/callback", i.callback)

		v1.Handler(api.Group("/v1"), i.authenticator)
	}

	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.server.Router
}

func (i *impl) callback(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	cb, err := webhook.ParseRequest(configx.A.LineBot.Secret, c.Request)
	if err != nil {
		if errors.Is(err, linebot.ErrInvalidSignature) {
			ctx.Error("invalid line bot signature", zap.Error(err))
			_ = c.Error(err)
		} else {
			ctx.Error("parse line bot request error", zap.Error(err))
			_ = c.Error(err)
		}

		return
	}

	var replyMessages []messaging_api.MessageInterface
	for _, event := range cb.Events {
		switch e := event.(type) {
		case webhook.MessageEvent:
			switch message := e.Message.(type) {
			case webhook.TextMessageContent:
				replyMessages, err = i.handleTextMessage(ctx, e, message)
				if err != nil {
					ctx.Error("handle text message error", zap.Error(err))
					_ = c.Error(err)
					return
				}

				_, err = i.bot.ReplyMessage(&messaging_api.ReplyMessageRequest{
					ReplyToken: e.ReplyToken,
					Messages:   replyMessages,
				})
				if err != nil {
					ctx.Error("reply message error", zap.Error(err))
					_ = c.Error(err)
					return
				}
			default:
				ctx.Debug("message type not support", zap.String("type", e.GetType()))
			}
		default:
			ctx.Debug("event type not support", zap.String("type", e.GetType()))
		}
	}

	c.Status(http.StatusOK)
}

func (i *impl) handleTextMessage(
	ctx contextx.Contextx,
	event webhook.MessageEvent,
	message webhook.TextMessageContent,
) ([]messaging_api.MessageInterface, error) {
	who := &idM.User{}

	switch source := event.Source.(type) {
	case webhook.UserSource:
		who.ID = source.UserId
	case webhook.GroupSource:
		who.ID = source.GroupId
	case webhook.RoomSource:
		who.ID = source.RoomId
	default:
		return handleError(errors.New("source type not support"))
	}

	return nil, errors.New("not support")
}

func handleError(err error) ([]messaging_api.MessageInterface, error) {
	return []messaging_api.MessageInterface{
		&messaging_api.TextMessage{
			Text: err.Error(),
		},
	}, nil
}
