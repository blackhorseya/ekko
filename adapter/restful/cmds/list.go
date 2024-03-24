package cmds

import (
	"errors"

	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/entity/domain/workflow/agg"
	"github.com/blackhorseya/ekko/entity/domain/workflow/biz"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// ListCommand is the command for list.
type ListCommand struct {
	workflow biz.IWorkflowBiz
}

func (cmd *ListCommand) Execute(
	ctx contextx.Contextx,
	who *idM.User,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if text == "/list" {
		var items agg.Issues
		items, _, err := cmd.workflow.ListTodos(ctx, who, biz.ListTodosOptions{Page: 1, Size: 5})
		if err != nil {
			return nil, err
		}

		if len(items) == 0 {
			return nil, errors.New("no todos")
		}

		container, err := items.FlexContainer()
		if err != nil {
			return nil, err
		}

		return []messaging_api.MessageInterface{
			&messaging_api.FlexMessage{
				AltText:  "Issue List",
				Contents: container,
			},
		}, nil
	}

	return nil, nil
}
