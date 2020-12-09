package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var cfgPath = flag.String("c", "configs/app.yaml", "set config file path")

// @title Todo list API
// @version 0.0.1
func main() {
	app, _, err := CreateApp(*cfgPath)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Panicf("create an application is panic")
	}

	logrus.WithFields(logrus.Fields{
		"config": app.C,
	}).Debugf("print config of application")

	err = app.Engine.Run(app.C.HTTP.GetAddress())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Panicf("run engine of app is panic")
	}
}
