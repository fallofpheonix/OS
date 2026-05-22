#!/bin/bash
set -e
echo "Building Phoenix Ledger..."
mkdir -p artifacts
go build -o artifacts/phoenix_ledger src/*.go
echo "Build complete."
