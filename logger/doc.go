// log package supplies more advanced features than go orign log package.
//
// It supports log different level: trace, debug, info, warn, error, fatal.
//
// It also supports different log handlers which you can log to stdout, file, socket, etc...
//
// Use
//
//  import "github.com/darrylwest/cassava-logger/logger"
//
//  //log with different level
//  log.Info("hello world")
//  log.Error("hello world")
//
//  // create a logger with specified handler
//  hander := NewStreamHandler(os.Stdout)
//  log := logger.NewDefault( handler )
//  log.Info("hello world")
//
package logger
