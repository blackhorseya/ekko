package main

import (
	"flag"
	"log"
)

var path = flag.String("c", "./configs/cli/ekko/local.yaml", "set config file path")

func init() {
	flag.Parse()
}

func main() {
	service, err := InitializeService(*path)
	if err != nil {
		log.Fatalln(err)
	}

	err = service.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
