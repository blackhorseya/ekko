package cmd

import (
	"fmt"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/todo"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.AddCommand(getTasksCmd)
}

var getCmd = &cobra.Command{
	Use:   "get [tasks]",
	Short: "Get something",
	Long:  "Get something",
}

var getTasksCmd = &cobra.Command{
	Use:   "tasks [id]",
	Short: "Display one or many tasks",
	Long:  "Display one or many tasks",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("arguments len must be less than 1")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var ret []*todo.Task
		var total int
		var err error

		if len(args) == 0 {
			ret, total, err = todoBiz.List(contextx.Background(), 0, 10)
			if err != nil {
				return err
			}
		}

		if len(args) == 1 {
			id := args[0]
			oid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return err
			}

			task, err := todoBiz.GetByID(contextx.Background(), oid)
			if err != nil {
				return err
			}

			ret = append(ret, task)
		}

		output, err := json.Marshal(map[string]interface{}{
			"total": total,
			"list":  ret,
		})
		if err != nil {
			return err
		}

		fmt.Println(string(output))

		return nil
	},
}
