#!/bin/bash
set -e
mkdir -p artifacts
echo "Running Tests..."
go test -v -race -coverprofile=artifacts/coverage.out ./...
go tool cover -func=artifacts/coverage.out | tee artifacts/coverage.txt
echo "Tests complete."
