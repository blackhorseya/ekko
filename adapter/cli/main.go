package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
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

	spew.Dump(config)

	fmt.Println("I'm a CLI")
}
