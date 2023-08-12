package main

import (
	"flag"
	"log"
	"os"
)

var path = flag.String("c", "", "path to config file (default: $HOME/.ekko/config.yaml)")

func init() {
	flag.Parse()

	if *path == "" {
		*path = os.Getenv("HOME") + "/.ekko/config.yaml"
	}
}

func main() {
	config, err := NewConfig(*path)
	if err != nil {
		log.Fatalln(err)
	}

	cmd, err := NewCmd(*config)
	if err != nil {
		log.Fatalln(err)
	}

	err = cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
