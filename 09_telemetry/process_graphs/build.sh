#!/bin/bash
set -e
echo "Building Process Graph..."
go build -o artifacts/process_graph ./src/*.go
echo "Build complete."
