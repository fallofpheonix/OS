#!/bin/bash
set -e
echo "Building Event Bus..."
go build -o artifacts/event_bus ./src/*.go
echo "Build complete."
