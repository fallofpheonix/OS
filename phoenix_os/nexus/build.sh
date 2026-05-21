#!/bin/bash
set -e
echo "Building Phoenix Nexus..."
mkdir -p artifacts
go build -o artifacts/phoenix_nexus ./src/*.go
echo "Build complete."
