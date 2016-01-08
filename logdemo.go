// Package logdemo is used to demonstrate unit-testing for the "log" package.
package logdemo

import (
	"log"
	"os"
)

type LoggingThing struct {
	Logger *log.Logger
}

// New returns a fresh new LoggingThing reference, which will log to STDOUT.
func New(name string) *LoggingThing {
	prefix := "thing-" + name + " "
	flags := log.Ldate | log.Ltime | log.Lmicroseconds
	return &LoggingThing{Logger: log.New(os.Stdout, prefix, flags)}
}

// Log writes logs something for the LoggingThing.
func (lt *LoggingThing) Log(msg string) {
	lt.Logger.Print(msg)
}
