package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

const applicationName = "eleveo-performance-export"

var (
	help      bool
	Version   string
	Revision  string
	Branch    string
	BuildUser string
	BuildDate string
)

/**
Main program
*/
func main() {
	timeStart := time.Now()
	exitCode := run()
	help = false
	logToFile = false
	logLevel = "DEBUG"

	initLog()

	timeEnd := time.Now()
	log.Printf("hello word %q", timeEnd.Sub(timeStart))
	//_ = level.Info(logger).Log("msg", "program end", "duration", timeEnd.Sub(timeStart))
	time.Sleep(time.Second)
	os.Exit(exitCode)
}

func run() int {
	return 0
}
