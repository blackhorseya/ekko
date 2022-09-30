package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [tasks]",
	Short: "Get something",
	Long:  "Get something",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get somethings")
	},
}
