package main

import (
	"flag"
	"log"
	"os"

	"go.uber.org/zap"
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

	logger, err := NewLogger(config)
	if err != nil {
		log.Fatalln(err)
	}

	cmd, err := NewCmd(*config, logger)
	if err != nil {
		logger.Fatal("failed to create cmd", zap.Error(err))
	}

	err = cmd.Execute()
	if err != nil {
		logger.Error("failed to execute cmd", zap.Error(err))
		os.Exit(1)
	}
}
