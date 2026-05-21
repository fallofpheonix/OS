#!/bin/bash
set -e
echo "Running Tests..."
go test -v -coverprofile=artifacts/coverage.out ./src/*.go
go tool cover -func=artifacts/coverage.out > artifacts/coverage.txt
echo "Tests complete."
