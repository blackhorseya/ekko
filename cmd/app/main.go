package main

import (
	"flag"

	"github.com/blackhorseya/todo-app/internal/app/config"
	"github.com/blackhorseya/todo-app/internal/pkg/logger"
	"github.com/sirupsen/logrus"
)

var cfgPath = flag.String("c", "configs/app.yaml", "set config file path")

func init() {
	flag.Parse()

	logger.SetDefault()
}

func initLogger(config config.Log) {
	// set level of log
	logger.SetLevel(config.Level)

	// set formatter of log
	logger.SetFormatter(config.Format)
}

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

	initLogger(app.C.Log)

	logrus.WithFields(logrus.Fields{
		"config": app.C,
	}).Debugf("print config of application")

	address := app.C.HTTP.GetAddress()
	logrus.WithFields(logrus.Fields{
		"host": app.C.HTTP.Host,
		"port": app.C.HTTP.Port,
	}).Infof("listening and serving HTTP")
	err = app.Engine.Run(address)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Panicf("run engine of app is panic")
	}

}
