package cmd

import (
	"github.com/blackhorseya/ekko/adapter/ekko/cmd/restful"
	"github.com/blackhorseya/ekko/internal/pkg/config"
	"github.com/blackhorseya/ekko/internal/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a service",
}

func init() {
	startCmd.AddCommand(startApiCmd)

	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var startApiCmd = &cobra.Command{
	Use:   "api",
	Short: "start api service",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.NewConfigWithViper(viper.GetViper())
		cobra.CheckErr(err)

		logger, err := log.NewLogger(cfg)
		cobra.CheckErr(err)

		service, err := restful.NewService(cfg, logger)
		cobra.CheckErr(err)

		err = service.Start()
		cobra.CheckErr(err)

		err = service.AwaitSignal()
		cobra.CheckErr(err)
	},
}
