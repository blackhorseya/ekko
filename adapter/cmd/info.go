package cmd

import (
	"github.com/blackhorseya/ekko/pkg/configx"
	"github.com/blackhorseya/ekko/pkg/contextx"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "print info",
	Run: func(cmd *cobra.Command, args []string) {
		contextx.Background().Info("info", zap.Any("config", &configx.C))
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
