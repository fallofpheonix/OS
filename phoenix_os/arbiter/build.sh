#!/bin/bash
set -e
echo "Building Phoenix Arbiter..."
mkdir -p artifacts
go build -o artifacts/phoenix_arbiter ./src/*.go
echo "Build complete."
