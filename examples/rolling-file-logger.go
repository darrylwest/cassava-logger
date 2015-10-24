package main

import (
    "../logger"
    "time"
    "fmt"
)

func main() {
    fmt.Printf("Version: %s\n", logger.Version)

    handler,_ := logger.NewTimeRotatingFileHandler( "./day-logger.out", logger.WhenDay, 1 )
    log := logger.NewLogger( handler )

    log.Debug("my debug message (suppressed)")
    log.Info("my message")
    log.Warn("my warning")
    log.Error("my error")

    log.SetLevel( logger.DebugLevel )
    log.Debug("this should show")

    time.Sleep(250 * time.Millisecond)
}
