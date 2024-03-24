package cmds

import (
	"errors"
	"strings"

	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/entity/domain/workflow/biz"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// CreateCommand is the command for create.
type CreateCommand struct {
	workflow biz.IWorkflowBiz
}

func (cmd *CreateCommand) Execute(
	ctx contextx.Contextx,
	who *idM.User,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if strings.HasPrefix(text, "/create ") {
		title := strings.TrimPrefix(text, "/create ")
		if len(title) == 0 {
			return nil, errors.New("title is required")
		}

		item, err := cmd.workflow.CreateTodo(ctx, who, title)
		if err != nil {
			return nil, err
		}

		container, err := item.FlexContainer()
		if err != nil {
			return nil, err
		}

		return []messaging_api.MessageInterface{
			messaging_api.FlexMessage{
				AltText:  "Issue Information",
				Contents: container,
			},
		}, nil
	}

	return nil, nil
}
