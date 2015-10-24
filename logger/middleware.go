package logger

import (
    "time"
    "net/http"
)
type MiddlewareLogger struct {
    log *Logger
}

func NewMiddlewareLogger(log *Logger) *MiddlewareLogger {
    return &MiddlewareLogger{ log }
}

func (m *MiddlewareLogger) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
    start := time.Now()
	m.log.Info("Started %s %s", req.Method, req.URL.Path)

    next(rw, req)

    m.log.Info("Completed in %v", time.Since(start))
}
