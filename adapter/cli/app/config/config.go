package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewConfigCmd is used to create config command
func NewConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "config is used to manage ekko config",
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "show",
		Short: "show is used to show ekko config",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("show config")
		},
	})

	return cmd
}
