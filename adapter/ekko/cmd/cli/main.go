package main

import (
	"log"
)

func main() {
	cmd, err := NewCmd()
	if err != nil {
		log.Fatalln(err)
	}

	err = cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
