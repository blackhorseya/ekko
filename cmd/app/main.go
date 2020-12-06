package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var cfgPath = flag.String("c", "app.yaml", "set config file path")

func main() {
	app, _, err := CreateApp(*cfgPath)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Panicf("create an application is panic")
	}

	err = app.Engine.Run(":8080")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Panicf("run engine of app is panic")
	}
}
