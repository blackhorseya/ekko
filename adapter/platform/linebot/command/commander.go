package command

import (
	"errors"
	"strings"

	"github.com/blackhorseya/ekko/adapter/platform/wirex"
	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/entity/domain/todo/biz"
	"github.com/blackhorseya/ekko/entity/domain/todo/model"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"go.uber.org/zap"
)

// TextCommander is the interface for text command.
type TextCommander interface {
	Execute(ctx contextx.Contextx, who *idM.User, text string) ([]messaging_api.MessageInterface, error)
}

// NewCommands is the function to create new commands.
func NewCommands(injector *wirex.Injector) []TextCommander {
	return []TextCommander{
		&PingCommand{},
		&WhoAmICommand{},
		&ListTodoCommand{injector: injector},
		&CreateTodoCommand{injector: injector},
		&DoneCommand{injector: injector},
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

// ListTodoCommand is the struct for list todo command.
type ListTodoCommand struct {
	injector *wirex.Injector
}

func (cmd *ListTodoCommand) Execute(
	ctx contextx.Contextx,
	who *idM.User,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if text == "/list" {
		var todos model.Todos
		var err error
		todos, _, err = cmd.injector.Todo.ListTodo(ctx, biz.ListTodoOptions{
			Page: 0,
			Size: 0,
		})
		if err != nil {
			return nil, err
		}

		if len(todos) == 0 {
			return nil, errors.New("no todos")
		}

		container, err := todos.FlexContainer()
		if err != nil {
			return nil, err
		}

		return []messaging_api.MessageInterface{
			&messaging_api.FlexMessage{
				AltText:  "Todo List",
				Contents: container,
			},
		}, nil
	}

	return nil, nil
}

// CreateTodoCommand is the struct for create todo command.
type CreateTodoCommand struct {
	injector *wirex.Injector
}

func (cmd *CreateTodoCommand) Execute(
	ctx contextx.Contextx,
	who *idM.User,
	text string,
) ([]messaging_api.MessageInterface, error) {
	if strings.HasPrefix(text, "/create ") {
		title := strings.TrimPrefix(text, "/create ")
		if len(title) == 0 {
			ctx.Error("title is required")
			return nil, errors.New("title is required")
		}

		todo, err := cmd.injector.Todo.CreateTodo(ctx, title)
		if err != nil {
			ctx.Error("create todo error", zap.Error(err))
			return nil, err
		}

		container, err := todo.FlexContainer()
		if err != nil {
			ctx.Error("todo flex container error", zap.Error(err))
			return nil, err
		}

		return []messaging_api.MessageInterface{
			&messaging_api.FlexMessage{
				AltText:  "Issue Information",
				Contents: container,
			},
		}, nil
	}

	return nil, nil
}
