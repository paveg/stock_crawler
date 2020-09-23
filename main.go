package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	logLevel := log.InfoLevel
	if os.Getenv("DEBUG_MODE") != "" {
		logLevel = log.DebugLevel
	}

	log.SetLevel(logLevel)
	log.Debugln("PS5 crawler has initialized.")
}

func main() {
	log.Infoln("execute main.")
}
