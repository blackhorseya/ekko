package app

import (
	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/spf13/cobra"
)

type cmd struct {
	rootCmd *cobra.Command
}

// NewCmd is used to create a new cmd instance
func NewCmd() adapters.CLI {
	return &cmd{
		rootCmd: &cobra.Command{
			Short:        "ekko is a tool for todo list management",
			SilenceUsage: true,
		},
	}
}

func (c *cmd) Execute() error {
	return c.rootCmd.Execute()
}
