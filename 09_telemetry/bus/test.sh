#!/bin/bash
set -e
echo "Running Tests..."
go test -v -coverprofile=artifacts/coverage.out ./src/*.go
echo "Tests complete."
