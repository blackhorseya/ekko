package main

import (
	"log"
)

func main() {
	service, err := InitializeService(1)
	if err != nil {
		log.Fatalln(err)
	}

	err = service.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
