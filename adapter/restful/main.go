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
		log.Fatal(err)
	}

	logger, err := NewLogger(config)
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("I'm a restful adapter")
}