package config

import (
	"encoding/json"
	"fmt"

	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/pkg/errors"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, ok := cmd.Context().Value("config").(*config.Config)
			if !ok {
				return errors.New("failed to get config")
			}

			bytes, err := json.MarshalIndent(cfg, "", "  ")
			cobra.CheckErr(err)

			fmt.Println(string(bytes))

			return nil
		},
	})

	return cmd
}
