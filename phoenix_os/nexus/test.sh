#!/bin/bash
set -e
echo "Testing Phoenix Nexus..."
go test -v ./src/*.go
echo "Tests complete."
