package app

import (
	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

func newConfigCmd(config *config.Config) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "config",
		Short: "config is used to manage ekko config",
	}

	versionCmd.AddCommand(&cobra.Command{
		Use:   "show",
		Short: "show is used to show ekko config",
		Run: func(cmd *cobra.Command, args []string) {
			spew.Dump(config)
		},
	})

	return versionCmd
}
