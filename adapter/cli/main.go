package main

import (
	"flag"
	"fmt"
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
	fmt.Println("I'm a CLI")
}
