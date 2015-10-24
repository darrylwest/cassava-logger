package main

import (
    "../logger"
    "time"
    "fmt"
    "os"
)

func main() {
    home := os.Getenv("HOME")
    logdir := home + "/logs/"
    filename := logdir + "day-logger"

    fmt.Printf("Version: %s\n", logger.Version)

    handler,_ := logger.NewRotatingDayHandler( filename )
    log := logger.NewLogger( handler )

    fmt.Printf("Log filename: %s\n", handler.CreateFilename())

    log.Debug("my debug message (suppressed)")
    log.Info("my message")
    log.Warn("my warning")
    log.Error("my error")

    log.SetLevel( logger.DebugLevel )
    log.Debug("this should show")

    time.Sleep(250 * time.Millisecond)
}
