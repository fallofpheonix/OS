#!/bin/bash
set -e
echo "Testing Phoenix Warden..."
go test -v ./src/*.go
echo "Tests complete."
