#!/bin/bash
set -e
echo "Testing Phoenix Sentinel..."
go test -v ./src/*.go
echo "Tests complete."
