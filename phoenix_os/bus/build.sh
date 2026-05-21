#!/bin/bash
set -e
echo "Building Phoenix Bus..."
mkdir -p artifacts
go build -o artifacts/phoenix_bus ./src/bus.go
echo "Build complete."
