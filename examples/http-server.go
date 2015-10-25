package main

import (
	"../logger"
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
	"os"
	"time"
)

type Context struct {
	port   int
	static string
	lg     *logger.Logger
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	var m = map[string]interface{}{
		"status":  "ok",
		"ts":      time.Now().UnixNano() / 1000000,
		"version": "1.0",
		"webStatus": map[string]interface{}{
			"version":    "2015-10-21",
			"pid":        os.Getpid(),
			"host":       r.Host,
			"path":       r.URL.Path,
			"agent":      r.UserAgent(),
		},
	}

	json, err := json.Marshal(m)

	if err != nil {
		fmt.Fprintf(w, "json error")
	} else {
		headers := w.Header()
		headers.Set("Content-Type", "application/json")

		w.Write(json)
	}
}

func startServer(context Context) {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", StatusHandler)

	server := negroni.New()
	server.Use(negroni.NewRecovery())
	server.Use(logger.NewMiddlewareLogger(context.lg))

	server.Use(negroni.NewStatic(http.Dir("staging")))

	server.UseHandler(mux)

    p := fmt.Sprintf(":%d", context.port )
	context.lg.Info("starting server at port: %s", p)

	server.Run( p )
}

func main() {
	handler, _ := logger.NewRotatingDayHandler("logger-middle")
	log := logger.NewLogger(handler)

	var port = 5001
	var static = "public"

	log.Info("start service on port %d", port)

	ctx := Context{port, static, log}
	startServer(ctx)
}
