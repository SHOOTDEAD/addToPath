package logger

import (
	"github.com/sirupsen/logrus"
)

var Logger = logger()

func logger() *logrus.Logger {

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	// 2006-01-02 15:04:05 this is a reference date through which go understands date format
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05"})

	return logger
}
