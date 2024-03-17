package cmds

import (
	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// PingCommand is the command for ping.
type PingCommand struct {
}

func (cmd *PingCommand) Execute(
	ctx contextx.Contextx,
	who *idM.User,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if text == "ping" {
		return []messaging_api.MessageInterface{
			&messaging_api.TextMessage{
				Text: "pong",
			},
		}, nil
	}

	return nil, nil
}
