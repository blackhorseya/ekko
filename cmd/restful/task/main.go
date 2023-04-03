package main

import (
	"flag"
	"log"
)

var path = flag.String("c", "./configs/restful/task/local.yaml", "set config file path")

func init() {
	flag.Parse()
}

// @title Ekko
// @version 0.0.2
// @description ekko is a simple task management system
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
	svc, err := CreateService(*path, 1)
	if err != nil {
		log.Fatalln(err)
	}

	err = svc.Start()
	if err != nil {
		log.Fatalln(err)
	}

	err = svc.AwaitSignal()
	if err != nil {
		log.Fatalln(err)
	}
}
