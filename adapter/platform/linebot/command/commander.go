package command

import (
	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// TextCommander is the interface for text command.
type TextCommander interface {
	Execute(ctx contextx.Contextx, who *idM.User, text string) ([]messaging_api.MessageInterface, error)
}

// NewCommands is the function to create new commands.
func NewCommands() []TextCommander {
	return []TextCommander{
		&PingCommand{},
		&WhoAmICommand{},
	}
}

// PingCommand is the struct for ping command.
type PingCommand struct {
}

func (cmd *PingCommand) Execute(
	ctx contextx.Contextx,
	who *idM.User,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if text == "/ping" {
		return []messaging_api.MessageInterface{
			&messaging_api.TextMessage{
				Text: "pong",
			},
		}, nil
	}

	return nil, nil
}

type WhoAmICommand struct {
}

func (cmd *WhoAmICommand) Execute(
	ctx contextx.Contextx,
	who *idM.User,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if text == "/whoami" {
		return []messaging_api.MessageInterface{
			&messaging_api.TextMessage{
				Text: who.ID,
			},
		}, nil
	}

	return nil, nil
}
