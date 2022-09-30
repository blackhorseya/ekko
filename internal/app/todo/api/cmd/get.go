package cmd

import (
	"fmt"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/goccy/go-json"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.AddCommand(getTaskByIdCmd)
}

var getCmd = &cobra.Command{
	Use:   "get [tasks]",
	Short: "Get something",
	Long:  "Get something",
}

var getTaskByIdCmd = &cobra.Command{
	Use:   "tasks [id]",
	Short: "Get a task by id",
	Long:  "Get a task by id",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("arguments len must be 1")
		}

		id := args[0]
		_, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		oid, _ := primitive.ObjectIDFromHex(id)

		ret, err := todoBiz.GetByID(contextx.Background(), oid)
		if err != nil {
			return err
		}

		output, err := json.Marshal(ret)
		if err != nil {
			return err
		}

		fmt.Println(string(output))

		return nil
	},
}
