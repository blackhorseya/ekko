package main

import (
	"flag"
	"log"

	"go.uber.org/zap"
)

var path = flag.String("c", "./configs/default.yaml", "path to config file (default: ./configs/default.yaml")

func init() {
	flag.Parse()
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

	service, err := NewService(config, logger, 0)
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
