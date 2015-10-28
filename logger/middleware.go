package logger

import (
	"net/http"
	"time"
    "strings"
)

type MiddlewareLogger struct {
	log   *Logger
	count int
}

func NewMiddlewareLogger(log *Logger) *MiddlewareLogger {
	return &MiddlewareLogger{log, 0}
}

func (m *MiddlewareLogger) Skip(path string, agent string) bool {
    if path == "/ping" && strings.Contains( strings.ToLower( agent ), "healthcheck" ) {
        return true
    }

    return false
}

func (m *MiddlewareLogger) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
    // count every request
    m.count++

    // skip the load balancer health-check pings
    if m.Skip( req.URL.Path, req.UserAgent() ) {
        next(rw, req)
    } else {
        start := time.Now()
        m.log.Info(">> %d %s %s %s %s", m.count, req.Method, req.Host, req.URL.Path, req.UserAgent())

        next(rw, req)

        m.log.Info("<< %d %v", m.count, time.Since(start))
    }
}
