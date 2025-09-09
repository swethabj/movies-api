package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger(level string) { // Takes a parameter level string → tells the logger what level to use (info, debug, warn, error).
	// set output to stdout (platform friendly)
	Logger.SetOutput(os.Stdout)
	// Development - os.Stdout (terminal)
	// Production - File (app.log) or centralized logging system
	// Cloud/Containerized - stdout/stderr, collected by logging agents

	// parse level (default Info)
	lvl, err := logrus.ParseLevel(level)
	//logrus.ParseLevel(level) converts a string like "debug" or "warn" to a Logrus log level type.
	//If parsing fails (invalid string), it defaults to InfoLevel.
	if err != nil {
		lvl = logrus.InfoLevel
	}
	logrus.SetLevel(lvl)
	//Logger.SetLevel(lvl) → now the logger will only output messages at or above this level.
	//Example: if level = warn, " info messages won’t appear " level debug, info , warn , error , fatal -- in asc order

	// use text formatter for dev; in prod use JSONFormatter
	Logger.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	//Logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
}
