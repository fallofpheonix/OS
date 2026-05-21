#!/bin/bash
set -e
echo "Running Tests..."
go test -v ./src/*.go
echo "Tests complete."
