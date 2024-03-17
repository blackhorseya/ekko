package restful

import (
	"errors"
	"strings"

	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/entity/domain/workflow/agg"
	"github.com/blackhorseya/ekko/entity/domain/workflow/biz"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// TextCommander is the interface for text command.
type TextCommander interface {
	Execute(ctx contextx.Contextx, who *idM.User, text string) ([]messaging_api.MessageInterface, error)
}

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

// ListCommand is the command for list.
type ListCommand struct {
	workflow biz.IWorkflowBiz
}

func (cmd *ListCommand) Execute(
	ctx contextx.Contextx,
	who *idM.User,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if text == "list" {
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

// CreateCommand is the command for create.
type CreateCommand struct {
	workflow biz.IWorkflowBiz
}

func (cmd *CreateCommand) Execute(
	ctx contextx.Contextx,
	who *idM.User,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if strings.HasPrefix(text, "create.") {
		title := strings.TrimPrefix(text, "create.")
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
