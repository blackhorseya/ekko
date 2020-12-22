package main

import (
	"flag"
	"fmt"

	"github.com/blackhorseya/todo-app/internal/app/commands"
	"github.com/blackhorseya/todo-app/internal/pkg/utils/exit"
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
		exit.Er(err)
	}
}

func initConfig() {
	fmt.Println(*cfgPath)
}
