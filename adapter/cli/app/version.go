package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of ekko",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ekko v0.0.1")
		},
	}
}
