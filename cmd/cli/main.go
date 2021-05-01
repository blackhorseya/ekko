package main

import (
	"flag"

	"github.com/blackhorseya/todo-app/internal/app/commands"
	exit2 "github.com/blackhorseya/todo-app/internal/pkg/base/exit"
	"github.com/spf13/cobra"
)

var cfgPath = flag.String("c", "configs/app.yaml", "set config file path")

func init() {
	flag.Parse()

	cobra.OnInitialize(initConfig)
}

func main() {
	err := commands.Execute()
	if err != nil {
		exit2.Er(err)
	}
}

func initConfig() {
	// todo: 2021-05-01|21:38|doggy|fix me
	// _, err := config.NewConfig(*cfgPath)
	// if err != nil {
	// 	exit2.Er(err)
	// }
}
