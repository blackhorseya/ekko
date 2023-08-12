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
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ekko version: v0.0.1")
		},
	}
}
