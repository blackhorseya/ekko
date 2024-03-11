package agg

import (
	"bytes"
	"embed"
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
	var containers []messaging_api.FlexBubble
	for _, issue := range x {
		container, err := issue.FlexContainer()
		if err != nil {
			return nil, err
		}

		containers = append(containers, container.(messaging_api.FlexBubble))
	}

	return &messaging_api.FlexCarousel{
		Contents: containers,
	}, nil
}
