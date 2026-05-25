#!/bin/bash
set -e
echo "Building Phoenix Trace..."
mkdir -p artifacts
go build -o artifacts/phoenix_trace ./src/*.go
echo "Build complete."
