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

// Log logs something for the LoggingThing.
func (lt *LoggingThing) Log(msg string) {
	lt.Logger.Print(msg)
}

// OK, that's all well and good, but what, you ask, if I'm too lazy to set
// up a structure and whatnot?
//
// In that case you should be able to use a package variable; just be sure
// to expose it by making it begin with an Uppercase letter.
var logger_prefix = "global "
var logger_flags = log.Ldate | log.Ltime | log.Lmicroseconds
var Logger = log.New(os.Stdout, logger_prefix, logger_flags)

// Log logs something out of the blue.
func Log(msg string) {
	Logger.Print(msg)
}
