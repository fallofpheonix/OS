#!/bin/bash
set -e
echo "Building Incident Physics..."
go build -o artifacts/physics ./src/*.go
echo "Build complete."
