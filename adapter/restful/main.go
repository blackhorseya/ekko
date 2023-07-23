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

// @title ekko
// @version 0.1.0
// @description ekko is a simple issue management system
//
// @contact.name sean.zheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space
//
// @license.name GPL-3.0
// @license.url https://spdx.org/licenses/GPL-3.0-only.html
//
// @BasePath /api
func main() {
	config, err := NewConfig(*path)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := NewLogger(config)
	if err != nil {
		log.Fatal(err)
	}

	service, err := NewService(config, logger)
	if err != nil {
		logger.Fatal("main: create service", zap.Error(err))
	}

	err = service.Start()
	if err != nil {
		logger.Fatal("main: failed to start service", zap.Error(err))
	}

	err = service.AwaitSignal()
	if err != nil {
		logger.Fatal("main: failed to await signal", zap.Error(err))
	}
}
