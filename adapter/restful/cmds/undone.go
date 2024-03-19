package cmds

import (
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
	// todo: 2024/3/19|sean|implement me
	return nil, nil
}
