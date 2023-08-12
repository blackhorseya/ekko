package app

import (
	"encoding/json"
	"fmt"

	"github.com/blackhorseya/ekko/internal/pkg/config"
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
			bytes, err := json.MarshalIndent(config, "", "  ")
			cobra.CheckErr(err)

			fmt.Println(string(bytes))
		},
	})

	return versionCmd
}
