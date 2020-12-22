package tasks

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "tasks",
		Short: "Tasks command",
	}
)

func init() {
	Cmd.AddCommand(listCmd)
	Cmd.AddCommand(deleteCmd)
}
