package app

import (
	"context"
	"os"

	"github.com/blackhorseya/ekko/adapter/ekko/cmd/cli/app/config"
	"github.com/blackhorseya/ekko/adapter/ekko/cmd/cli/app/version"
	configx "github.com/blackhorseya/ekko/internal/pkg/config"
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

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if path == "" {
			path = os.Getenv("HOME") + "/.ekko/config.yaml"
		}

		cfg, err := configx.NewWithPath(path)
		cobra.CheckErr(err)

		cmd.SetContext(context.WithValue(cmd.Context(), "config", cfg))
	}

	return &cmd{
		rootCmd: rootCmd,
	}
}

func (c *cmd) Execute() error {
	return c.rootCmd.Execute()
}
