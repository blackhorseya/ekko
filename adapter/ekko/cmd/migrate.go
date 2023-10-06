package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
}

func init() {
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)

	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "migrate up",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate up")
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "migrate down",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate down")
	},
}
