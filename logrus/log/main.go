package main

import (
	"errors"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	// Println adds spaces, other functions do not
	logrus.Println("Hello,", "Info,", "from", "Println")

	logrus.Trace("Hello, Trace")
	logrus.Debug("Hello, ", "Debug")
	logrus.Info("Hello, Info")
	logrus.Warn("Hello, Warn")
	logrus.Error("Hello, ", errors.New("Error"))
	// logrus.Fatal("Bye, Fatal")
	// logrus.Panic("Bye, Panic")
	logrus.WithFields(logrus.Fields{
		"from": "Logrus",
		"to":   "World",
	}).Info("Hello")
	logrus.WithError(errors.New("Something went wrong!")).Error("Oops")
	logrus.WithFields(logrus.Fields{
		"from": "Logrus",
		"to":   "World",
	}).Fatal("Bye")
}
