package main

import (
	"flag"
	"fmt"
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
	_ = config

	fmt.Println("I'm a restful adapter")
}
