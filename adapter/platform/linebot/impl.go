package linebot

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/ekko/adapter/platform/linebot/command"
	"github.com/blackhorseya/ekko/adapter/platform/wirex"
	"github.com/blackhorseya/ekko/app/infra/configx"
	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/pkg/adapterx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/blackhorseya/ekko/pkg/responsex"
	"github.com/blackhorseya/ekko/pkg/transports/httpx"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"go.uber.org/zap"
)

type impl struct {
	injector *wirex.Injector
	server   *httpx.Server

	bot      *messaging_api.MessagingApiAPI
	commands []command.TextCommander
}

func newRest(injector *wirex.Injector, server *httpx.Server, bot *messaging_api.MessagingApiAPI) adapterx.Restful {
	return &impl{
		injector: injector,
		server:   server,
		bot:      bot,
		commands: command.NewCommands(injector),
	}
}

func newService(injector *wirex.Injector, server *httpx.Server, bot *messaging_api.MessagingApiAPI) adapterx.Servicer {
	return newRest(injector, server, bot)
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
				replyMessages, err = i.generateReplyMessage(ctx, e, message)
				if err != nil {
					ctx.Error("handle text message error", zap.Error(err))
					_ = c.Error(err)
					return
				}

				// if no reply message, skip
				if len(replyMessages) == 0 {
					continue
				}

				// reply message
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

	responsex.OK(c, nil)
}

func (i *impl) generateReplyMessage(
	ctx contextx.Contextx,
	event webhook.MessageEvent,
	message webhook.TextMessageContent,
) ([]messaging_api.MessageInterface, error) {
	text := message.Text
	who := &idM.User{}

	switch source := event.Source.(type) {
	case webhook.UserSource:
		who.ID = source.UserId
	case webhook.GroupSource:
		who.ID = source.GroupId
	case webhook.RoomSource:
		who.ID = source.RoomId
	default:
		return nil, errors.New("source type not support")
	}
	ctx = contextx.WithValue(ctx, contextx.KeyWho, who)

	for _, cmd := range i.commands {
		messages, err := cmd.Execute(ctx, who, text)
		if err != nil {
			return nil, err
		}

		if len(messages) > 0 {
			return messages, nil
		}
	}

	return nil, nil
}
