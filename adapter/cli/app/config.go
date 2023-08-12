package app

import (
	"fmt"

	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/spf13/cobra"
)

func newConfigCmd(config *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "config is used to manage ekko config",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(config)
		},
	}
}
