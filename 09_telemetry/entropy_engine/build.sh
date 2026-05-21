#!/bin/bash
set -e
echo "Building Entropy Engine..."
go build -o artifacts/entropy_engine ./src/*.go
sha256sum artifacts/entropy_engine > hashes.sha256
echo "Build complete."
