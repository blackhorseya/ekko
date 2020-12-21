package main

import (
	"github.com/blackhorseya/todo-app/internal/app/commands"
	"github.com/blackhorseya/todo-app/internal/pkg/utils/exit"
)

func main() {
	err := commands.Execute()
	if err != nil {
		exit.Er(err)
	}
}
