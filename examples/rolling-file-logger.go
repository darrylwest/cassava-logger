package main

import (
    "../logger"
)

func main() {
    handler,_ := logger.NewTimeRotatingFileHandler( "./day-logger.out", logger.WhenDay, 1 )
    log := logger.NewLogger( handler )

    log.Info("my message")
}
