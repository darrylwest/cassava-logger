# cassava-logger

A log4j-like logger with levels, multiple handlers and middleware implementation.  This implementation is optimized for rolling files suitable for services, middleware, etc.

## Install

go get github.com/darrylwest/cassava-logger/logger


## Simple Use

    import "github.com/darrylwest/cassava-logger/logger"

    // create a logger with specified handler
    handler,_ := logger.NewStreamHandler(os.Stdout)
    log := logger.NewLogger( handler )

    log.Debug("debug level log statement") // will be suppressed
    log.Info("info level log statement")
    log.Warn("warn level log statement")
    log.Error("errorlevel log statement")
    
    log.SetLevel( logger.LevelDebug )
    log.Debug("this debug message shows up...")


- - -
<h6><small>darryl.west@raincitysoftware.com | Version 0.90.101</small></h6>