#!/bin/bash
set -e
echo "Testing Phoenix Arbiter..."
go test -v ./src/*.go
echo "Tests complete."
