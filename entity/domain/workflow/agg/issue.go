package agg

import (
	"errors"

	"github.com/blackhorseya/ekko/entity/domain/workflow/model"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

// Issue is an aggregate root that represents an issue.
type Issue struct {
	*model.Ticket
}

// FlexMessage returns the issue as a flex message.
func (x *Issue) FlexMessage() (messaging_api.MessageInterface, error) {
	// todo: 2024/3/10|sean|implement this method
	return nil, errors.New("issue flex message not implemented")
}

// Issues is a collection of Issue.
type Issues []*Issue

// FlexMessage returns the issues as a flex message.
func (x Issues) FlexMessage() (messaging_api.MessageInterface, error) {
	// todo: 2024/3/10|sean|implement this method
	return nil, errors.New("issues flex message not implemented")
}
