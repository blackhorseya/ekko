package main

import (
	"flag"
	"strings"

	"github.com/blackhorseya/todo-app/internal/app/config"
	"github.com/sirupsen/logrus"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

var cfgPath = flag.String("c", "configs/app.yaml", "set config file path")

func init() {
	flag.Parse()

	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: timeFormat,
		DisableQuote:    true,
	})
}

func initLogger(config config.Log) {
	// set level of log
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Panicf("parse log.level is panic")
	}
	logrus.SetLevel(level)

	// set formatter of log
	if strings.ToLower(config.Format) == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: timeFormat,
		})
	}
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

	err = app.Engine.Run(app.C.HTTP.GetAddress())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Panicf("run engine of app is panic")
	}
}
