#!/bin/bash
set -e
echo "Building Incident Physics..."
mkdir -p artifacts
go build -o artifacts/physics .
echo "Build complete."
