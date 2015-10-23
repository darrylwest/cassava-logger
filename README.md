# cassava-logger

A log4j-like loggger with levels, multiple handlers and middleware implementation.

## Simple Use

    import "github.com/darrylwest/cassava-logger/logger"

    // create a logger with specified handler
    handler := logger.NewStreamHandler(os.Stdout)
    log := logger.NewLogger( handler )

    log.Debug("debug level log statement")
    log.Info("info level log statement")
    log.Warn("warn level log statement")
    log.Error("errorlevel log statement")


