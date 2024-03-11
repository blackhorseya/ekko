package agg

import (
	"bytes"
	"embed"
	"errors"
	"text/template"

	"github.com/blackhorseya/ekko/entity/domain/workflow/model"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

//go:embed issue.tmpl
var f embed.FS

// Issue is an aggregate root that represents an issue.
type Issue struct {
	*model.Ticket
}

// FlexContainer returns the issue as a flex message.
func (x *Issue) FlexContainer() (messaging_api.FlexContainerInterface, error) {
	tmpl, err := template.New("issue.tmpl").ParseFS(f, "issue.tmpl")
	if err != nil {
		return nil, err
	}

	var layout bytes.Buffer
	err = tmpl.Execute(&layout, x)
	if err != nil {
		return nil, err
	}

	return messaging_api.UnmarshalFlexContainer(layout.Bytes())
}

// Issues is a collection of Issue.
type Issues []*Issue

// FlexContainer returns the issues as a flex message.
func (x Issues) FlexContainer() (messaging_api.FlexContainerInterface, error) {
	// todo: 2024/3/10|sean|implement this method
	return nil, errors.New("issues flex message not implemented")
}
