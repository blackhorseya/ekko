package cmd

import (
	"fmt"
	"os"

	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/consts"
	"github.com/google/wire"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "Todo APP cli",
	Long: `
 ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄               ▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄        ▄ 
▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌             ▐░▌▐░░░░░░░░░░░▌▐░░▌      ▐░▌
 ▀▀▀▀▀█░█▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀█░▌ ▐░▌           ▐░▌ ▐░█▀▀▀▀▀▀▀█░▌▐░▌░▌     ▐░▌
      ▐░▌    ▐░▌       ▐░▌▐░▌       ▐░▌  ▐░▌         ▐░▌  ▐░▌       ▐░▌▐░▌▐░▌    ▐░▌
      ▐░▌    ▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌   ▐░▌       ▐░▌   ▐░█▄▄▄▄▄▄▄█░▌▐░▌ ▐░▌   ▐░▌
      ▐░▌    ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌    ▐░▌     ▐░▌    ▐░░░░░░░░░░░▌▐░▌  ▐░▌  ▐░▌
      ▐░▌    ▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀█░█▀▀      ▐░▌   ▐░▌     ▐░█▀▀▀▀▀▀▀█░▌▐░▌   ▐░▌ ▐░▌
      ▐░▌    ▐░▌       ▐░▌▐░▌     ▐░▌        ▐░▌ ▐░▌      ▐░▌       ▐░▌▐░▌    ▐░▌▐░▌
 ▄▄▄▄▄█░▌    ▐░▌       ▐░▌▐░▌      ▐░▌        ▐░▐░▌       ▐░▌       ▐░▌▐░▌     ▐░▐░▌
▐░░░░░░░▌    ▐░▌       ▐░▌▐░▌       ▐░▌        ▐░▌        ▐░▌       ▐░▌▐░▌      ▐░░▌
 ▀▀▀▀▀▀▀      ▀         ▀  ▀         ▀          ▀          ▀         ▀  ▀        ▀▀
`,
}

var (
	cfgFile string
	level   string
	output  string

	todoBiz todo.ITodoBiz
)

func NewRootCmd(biz todo.ITodoBiz) (*cobra.Command, error) {
	todoBiz = biz

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/.%s.yaml)", consts.AppName))
	rootCmd.PersistentFlags().StringVar(&level, "log-level", "info", "log level")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "console", "output style [console, json]")

	rootCmd.Version = consts.Version

	return rootCmd, nil
}

func NewViper() (*viper.Viper, error) {
	initConfig()

	return viper.GetViper(), nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("." + consts.AppName)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewRootCmd, NewViper)
