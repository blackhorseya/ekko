package main

import "github.com/sirupsen/logrus"

func main() {
	app, _, err := CreateApp()
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
