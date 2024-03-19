package cmds

import (
	idM "github.com/blackhorseya/ekko/entity/domain/identity/model"
	"github.com/blackhorseya/ekko/entity/domain/workflow/biz"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// TextCommander is the interface for text command.
type TextCommander interface {
	Execute(ctx contextx.Contextx, who *idM.User, text string) ([]messaging_api.MessageInterface, error)
}

// NewCommands is the function to create new commands.
func NewCommands(workflow biz.IWorkflowBiz) []TextCommander {
	return []TextCommander{
		&PingCommand{},
		&WhoAmICommand{},
		&ListCommand{workflow: workflow},
		&CreateCommand{workflow: workflow},
		&DoneCommand{workflow: workflow},
		&UndoneCommand{workflow: workflow},
	}
}
