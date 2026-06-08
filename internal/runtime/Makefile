# PHOENIX MATRIX SOVEREIGN ARCHITECTURE
# [STATUS]: 18-Repository Substrate Consolidated
# [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
# [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
# [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
.PHONY: all build test clean proto

all: build test

build:
	go build ./...

test:
	go test -v -race ./...

proto:
	PATH=$(PATH):$(HOME)/go/bin protoc --go_out=. --go_opt=paths=source_relative proto/v1/fsm/fsm.proto
	PATH=$(PATH):$(HOME)/go/bin protoc --go_out=. --go_opt=paths=source_relative proto/v1/event/event.proto

clean:
	go clean
	rm -f bin/*
