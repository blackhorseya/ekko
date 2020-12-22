package tasks

import "github.com/spf13/cobra"

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a task",
	}
)
