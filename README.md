# cassava-logger

A log4j-like logger with levels, multiple handlers and middleware implementation.  This implementation is optimized for rolling files suitable for services, middleware, etc.


<em>This package is based on [siddontang's](https://github.com/siddontang/go-log) go-log project.</em>

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
    
    log.SetLevel( logger.DebugLevel )
    log.Debug("this debug message shows up...")


## Rolling File Logger

You will find an example similar to this in "examples".  The log's filename is changed to "day-logger.YYYY-MM-DD.log" and rotates daily.

	handler,_ := logger.NewRotatingDayHandler( "./day-logger" )
	log := logger.NewLogger( handler )

	log.Debug("my debug message (suppressed)")
	log.Info("my message")
	log.Warn("my warning")
	log.Error("my error")
	
	log.SetLevel( logger.DebugLevel )
	log.Debug("this should show")	
	

- - -
<em><small>Version 0.90.102 | darryl.west@raincitysoftware.com</small></em>