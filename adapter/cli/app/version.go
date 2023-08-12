package app

import (
	"fmt"

	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/spf13/cobra"
)

func newVersionCmd(config *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of ekko",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s version %s\n", config.App.Name, config.App.Version)
		},
	}
}
