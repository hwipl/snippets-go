package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("Hello, Trace")
	logrus.Debug("Hello, Debug")
	logrus.Info("Hello, Info")
	logrus.Warn("Hello, Warn")
	logrus.Error("Hello, Error")
	// logrus.Fatal("Bye, Fatal")
	// logrus.Panic("Bye, Panic")
	logrus.WithFields(logrus.Fields{
		"from": "Logrus",
		"to":   "World",
	}).Info("Hello")
	logrus.WithFields(logrus.Fields{
		"from": "Logrus",
		"to":   "World",
	}).Fatal("Bye")
}
