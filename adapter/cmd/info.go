package cmd

import (
	"encoding/json"

	"github.com/blackhorseya/ekko/app/infra/configx"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "print info",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := json.Marshal(configx.C)
		cobra.CheckErr(err)

		cmd.Println(string(data))
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
