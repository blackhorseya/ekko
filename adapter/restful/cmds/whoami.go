package cmds

import (
	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// WhoAmICommand is the command for who am i.
type WhoAmICommand struct {
}

func (cmd *WhoAmICommand) Execute(
	ctx contextx.Contextx,
	who *idM.User,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if text == "whoami" {
		return []messaging_api.MessageInterface{
			&messaging_api.TextMessage{
				Text: who.ID,
			},
		}, nil
	}

	return nil, nil
}
