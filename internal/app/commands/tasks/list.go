package tasks

import "github.com/spf13/cobra"

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Print all tasks",
	}
)
