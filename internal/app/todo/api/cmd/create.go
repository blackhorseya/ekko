package cmd

import (
	"fmt"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.AddCommand(createTaskCmd)
}

var createCmd = &cobra.Command{
	Use:   "create [tasks]",
	Short: "Create one resource",
	Long:  "Create one resource",
}

var createTaskCmd = &cobra.Command{
	Use:   "tasks [title]",
	Short: "Create a task",
	Long:  "Create a task",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("arguments len must be 1")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		title := args[0]

		logger, _ := zap.NewDevelopment()
		task, err := todoBiz.Create(contextx.BackgroundWithLogger(logger), title)
		if err != nil {
			return err
		}

		output, err := json.Marshal(task)
		if err != nil {
			return err
		}

		fmt.Println(string(output))

		return nil
	},
}
