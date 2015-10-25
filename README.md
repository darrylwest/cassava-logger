# Cassava Logger

A log4j-like logger with levels, multiple handlers and middleware implementation.  This implementation is optimized for rolling files suitable for services, middleware, etc.


<em>This package is based on [siddontang's](https://github.com/siddontang/go-log) go-log project.</em>

## Install

go get github.com/darrylwest/cassava-logger/logger


## Simple Use

There are instance and type methods (Trace, Debug, Info, Warn, Error, Fatal) that support simple stream logging to the console.  An example of type use would be:

	logger.Info("this is a test...") 
	// -> hh:mm:ss.SSS INFO filename:line this is a test...
	
	logger.Warn("this is a warning") 
	// -> hh:mm:ss.SSS WARN filename:line this is a warning

There are the same methods implemented on an instance, for example:

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

There are multiple rolling file handlers to support rolling by size or time.  Times include second, minute, hour and day.  The rotating day handler is used in the logger middleware project.

You will find an example similar to this in "examples".  The log's filename is changed to "day-logger.YYYY-MM-DD.log" and rotates daily.

	handler,_ := logger.NewRotatingDayHandler( "./day-logger" )
	log := logger.NewLogger( handler )

	log.Debug("my debug message (suppressed)")
	log.Info("my message")
	log.Warn("my warning")
	log.Error("my error")
	
	log.SetLevel( logger.DebugLevel )
	log.Debug("this should show")	
	
## To Do

* add categories to optionally replace file/line in log statements
* create config for reading and re-reading configurations
* complete unit tests for 100% coverage
* add travis tests


## License: MIT

Use as you wish.  Fork and help out if you can.

- - -
<em><small>Version 0.90.102 | darryl.west@raincitysoftware.com</small></em>