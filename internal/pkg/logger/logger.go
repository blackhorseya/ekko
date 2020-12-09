package logger

import "github.com/sirupsen/logrus"

const timeFormat = "2006-01-02 15:04:05"

// SetDefault set default logrus
func SetDefault() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: timeFormat,
		DisableQuote:    true,
	})
}

// SetLevel set level of logrus
func SetLevel(lvl string) {
	if level, err := logrus.ParseLevel(lvl); err != nil {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.Warn(err)
	} else {
		logrus.SetLevel(level)
	}
}

// SetFormatter set formatter of logrus
func SetFormatter(format string) {
	switch format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: timeFormat,
		})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: timeFormat,
			DisableQuote:    true,
		})
	}
}
