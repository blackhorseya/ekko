package model

import (
	"bytes"
	"embed"
	"errors"
	"html/template"
	"time"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

//go:embed todo.tmpl
var f embed.FS

// Todo is the aggregate root of the todo domain
type Todo struct {
	ID        string    `json:"id,omitempty" bson:"_id"`
	Title     string    `json:"title,omitempty" bson:"title"`
	Done      bool      `json:"done,omitempty" bson:"done"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

// NewTodo is to create a new todo
func NewTodo(title string) (*Todo, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}

	return &Todo{
		Title:     title,
		UpdatedAt: time.Now(),
	}, nil
}

// FlexContainer returns the issue as a flex message.
func (x *Todo) FlexContainer() (messaging_api.FlexContainerInterface, error) {
	tmpl, err := template.New("todo.tmpl").ParseFS(f, "todo.tmpl")
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

// Todos is a collection of Todo.
type Todos []*Todo

// FlexContainer returns the issues as a flex message.
func (x Todos) FlexContainer() (messaging_api.FlexContainerInterface, error) {
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
