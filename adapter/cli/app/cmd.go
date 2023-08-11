package app

import (
	"github.com/blackhorseya/ekko/pkg/adapters"
)

type cmd struct {
}

// NewCmd is used to create a new cmd instance
func NewCmd() adapters.CLI {
	return &cmd{}
}

func (c *cmd) Execute() error {
	panic("implement me")
}
