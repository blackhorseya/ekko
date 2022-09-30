package cmd

import (
	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

var getCmd = &cobra.Command{
	Use:   "get [tasks]",
	Short: "Get something",
	Long:  "Get something",
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewDevelopment()
		id, err := primitive.ObjectIDFromHex("6336564795dbab5cbe18f6f4")
		cobra.CheckErr(err)

		task, err := todoBiz.GetByID(contextx.BackgroundWithLogger(logger), id)
		cobra.CheckErr(err)

		logger.Info("get something", zap.Any("task", task))
	},
}
