#!/bin/bash
set -e
echo "Building Phoenix Warden..."
mkdir -p artifacts
go build -o artifacts/phoenix_warden ./src/*.go
echo "Build complete."
