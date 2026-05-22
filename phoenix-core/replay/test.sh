#!/bin/bash
set -e
echo "Testing Phoenix Trace..."
go test -v ./src/*.go
echo "Tests complete."
