package command

import (
	"errors"
	"strings"

	"github.com/blackhorseya/ekko/adapter/platform/wirex"
	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

var _ TextCommander = (*DoneCommand)(nil)

// DoneCommand is the struct for done command.
type DoneCommand struct {
	injector *wirex.Injector
}

func (cmd *DoneCommand) Execute(
	ctx contextx.Contextx,
	who *idM.User,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if strings.HasPrefix(text, "/done") {
		id := strings.TrimSpace(strings.TrimPrefix(text, "/done"))
		if id == "" {
			return nil, errors.New("missing id")
		}

		todo, err := cmd.injector.Todo.CompleteTodo(ctx, id)
		if err != nil {
			return nil, err
		}

		container, err := todo.FlexContainer()
		if err != nil {
			return nil, err
		}

		return []messaging_api.MessageInterface{
			&messaging_api.FlexMessage{
				AltText:  "Done",
				Contents: container,
			},
		}, nil
	}

	return nil, nil
}
