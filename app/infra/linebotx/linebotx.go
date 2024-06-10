package linebotx

import (
	"github.com/blackhorseya/ekko/app/infra/configx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// NewClient is used to create a new line bot client.
func NewClient() (*messaging_api.MessagingApiAPI, error) {
	return messaging_api.NewMessagingApiAPI(configx.A.LineBot.Token)
}
