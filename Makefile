
format:
	( cd logger ; gofmt -w *.go )
	( cd socket-logger ; gofmt -w *.go )

test:
	( cd logger ; go test -cover )

watch:
	./watcher.js

.PHONY: format
.PHONY: test
.PHONY: watch
.PHONY: run
