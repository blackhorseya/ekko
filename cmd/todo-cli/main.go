package main

import (
	"github.com/spf13/cobra"
)

func main() {
	command, err := CreateApp()
	if err != nil {
		panic(err)
	}

	err = command.Execute()
	cobra.CheckErr(err)
}
