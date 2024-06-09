package cmd

import (
	"github.com/blackhorseya/ekko/adapter/platform/linebot"
	"github.com/blackhorseya/ekko/adapter/platform/rest"
	"github.com/blackhorseya/ekko/pkg/cmdx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a service",
}

func init() {
	startAPICmd := cmdx.NewServiceCmd("api", "start a restful server", rest.New)
	startAPICmd.Flags().Bool("auth", false, "enable auth middleware")
	_ = viper.GetViper().BindPFlag("auth0.enabled", startAPICmd.Flags().Lookup("auth"))

	startCmd.AddCommand(startAPICmd)
	startCmd.AddCommand(cmdx.NewServiceCmd("linebot", "start a linebot server", linebot.New))

	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
