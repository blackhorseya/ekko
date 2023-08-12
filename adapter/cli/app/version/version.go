package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewVersionCmd is used to create a new version cmd instance
func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of ekko",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("ekko version 0.0.1")

			return nil
		},
	}
}
