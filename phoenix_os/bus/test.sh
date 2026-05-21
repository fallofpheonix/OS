#!/bin/bash
set -e
echo "Testing Phoenix Bus..."
go test -v ./src/*.go
echo "Tests complete."
