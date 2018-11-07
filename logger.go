// Copyright 2015 Nevio Vesic
// Please check out LICENSE file for more information about what you CAN and what you CANNOT do!
// Basically in short this is a free software for you to do whatever you want to do BUT copyright must be included!
// I didn't write all of this code so you could say it's yours.
// MIT License

package goesl

import (
	"os"

	logging "github.com/githmiao/go-logging"
)

var (
	log = logging.MustGetLogger("goesl")

	// Example format string. Everything except the message has a custom color
	// which is dependent on the log level. Many fields have a custom output
	// formatting too, eg. the time returns the hour down to the milli second.
	format = logging.MustStringFormatter(
		"%{shortfunc}->%{level:.8s} %{message}",
	)
)

func Debug(message string, args ...interface{}) {
	log.Debugf(message, args...)
}

func Error(message string, args ...interface{}) {
	log.Errorf(message, args...)
}

func Notice(message string, args ...interface{}) {
	log.Noticef(message, args...)
}

func Info(message string, args ...interface{}) {
	log.Infof(message, args...)
}

func Warning(message string, args ...interface{}) {
	log.Warningf(message, args...)
}

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(formatter)
}
