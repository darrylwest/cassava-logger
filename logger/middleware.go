package logger

import (
	"net/http"
	"time"
)

type MiddlewareLogger struct {
	log   *Logger
	count int
}

func NewMiddlewareLogger(log *Logger) *MiddlewareLogger {
	return &MiddlewareLogger{log, 0}
}

func (m *MiddlewareLogger) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	start := time.Now()
	m.count++
	m.log.Info(">> %d %s %s %s %s", m.count, req.Method, req.Host, req.URL.Path, req.UserAgent())

	next(rw, req)

	m.log.Info("<< %d %v", m.count, time.Since(start))
}
