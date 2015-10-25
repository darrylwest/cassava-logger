package main

import (
	"../logger"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	home := os.Getenv("HOME")
	logdir := home + "/logs/"
	filename := logdir + "day-logger"

	fmt.Printf("Version: %s\n", logger.Version)

	handler, _ := logger.NewRotatingDayHandler(filename)
	log := logger.NewLogger(handler)

	fmt.Printf("Log filename: %s\n", handler.CreateFilename())

	log.Debug("my debug message: (%s)", "suppressed")
	log.Info("my message number: %d", 443)
	log.Warn("my warning at: %v", time.Now())
	log.Error("my error")

	log.SetLevel(logger.DebugLevel)
	log.Debug("this should show")

	// put in a short sleep
	time.Sleep(10 * time.Millisecond)
	log.Info("completed after: %v", time.Since(start))

	time.Sleep(250 * time.Millisecond)
}
