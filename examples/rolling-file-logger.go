package main

import (
    "../logger"
    "time"
)

func main() {
    handler,_ := logger.NewTimeRotatingFileHandler( "./day-logger.out", logger.WhenDay, 1 )
    log := logger.NewLogger( handler )

    log.Debug("my debug message")
    log.Info("my message")
    log.Warn("my warning")
    log.Error("my error")

    time.Sleep(250 * time.Millisecond)
}
