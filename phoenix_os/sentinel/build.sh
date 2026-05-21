#!/bin/bash
set -e
echo "Building Phoenix Sentinel..."
mkdir -p artifacts
go build -o artifacts/phoenix_sentinel ./src/*.go
echo "Build complete."
