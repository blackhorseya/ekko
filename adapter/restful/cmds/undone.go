package cmds

import (
	"errors"
	"strings"

	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/entity/domain/workflow/biz"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// UndoneCommand is a command to mark a task as undone.
type UndoneCommand struct {
	workflow biz.IWorkflowBiz
}

func (cmd *UndoneCommand) Execute(
	ctx contextx.Contextx,
	who *idM.User,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if strings.HasPrefix(text, "/undone ") {
		id := strings.TrimPrefix(text, "/undone ")
		if len(id) == 0 {
			return nil, errors.New("id is required")
		}

		err := cmd.workflow.UndoneTodoByID(ctx, who, id)
		if err != nil {
			return nil, err
		}

		got, err := cmd.workflow.GetTodoByID(ctx, who, id)
		if err != nil {
			return nil, err
		}

		container, err := got.FlexContainer()
		if err != nil {
			return nil, err
		}

		return []messaging_api.MessageInterface{
			&messaging_api.FlexMessage{
				AltText:  got.Title,
				Contents: container,
			},
		}, nil
	}

	return nil, nil
}
