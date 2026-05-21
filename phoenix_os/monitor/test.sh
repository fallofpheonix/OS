#!/bin/bash
set -e
echo "Testing Phoenix Monitor..."
go test -v ./src/*.go
echo "Tests complete."
