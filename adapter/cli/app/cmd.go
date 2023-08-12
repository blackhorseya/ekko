package app

import (
	"github.com/blackhorseya/ekko/adapter/cli/app/config"
	"github.com/blackhorseya/ekko/adapter/cli/app/version"
	"github.com/blackhorseya/ekko/pkg/adapters"
	"github.com/spf13/cobra"
)

type cmd struct {
	rootCmd *cobra.Command
}

// NewCmd is used to create a new cmd instance
func NewCmd() adapters.CLI {
	rootCmd := &cobra.Command{
		Short:        "ekko is a tool for todo list management",
		SilenceUsage: true,
	}

	var path string
	rootCmd.PersistentFlags().StringVarP(&path, "config", "f", "", "path to config file (default: $HOME/.ekko/config.yaml)")

	rootCmd.AddCommand(version.NewVersionCmd())
	rootCmd.AddCommand(config.NewConfigCmd())

	return &cmd{
		rootCmd: rootCmd,
	}
}

func (c *cmd) Execute() error {
	return c.rootCmd.Execute()
}
