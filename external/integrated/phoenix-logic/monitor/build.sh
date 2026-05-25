#!/bin/bash
set -e
echo "Building Phoenix Monitor..."
mkdir -p artifacts
go build -o artifacts/phoenix_monitor ./src/*.go
echo "Build complete."
