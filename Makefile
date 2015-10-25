
format:
	( cd logger ; gofmt -w *.go )
	( cd socket-logger ; gofmt -w *.go )
	( cd examples ; gofmt -w *.go )

run:
	( go run examples/rolling-file-logger.go )

test:
	( cd logger ; go test -cover )

watch:
	./watcher.js

.PHONY: format
.PHONY: test
.PHONY: watch
.PHONY: run
