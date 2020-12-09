package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var cfgPath = flag.String("c", "configs/app.yaml", "set config file path")

// @title Todo list API
// @version 0.0.1
// @description Todo list API

// @contact.name Sean Cheng
// @contact.email blackhorseya@gmail.com
// @contact.url https://blog.seancheng.space

// @license.name GPL-3.0
// @license.url https://spdx.org/licenses/GPL-3.0-only.html

// @BasePath /api
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
