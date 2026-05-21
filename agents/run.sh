#!/bin/bash
set -e
echo "Running agents..."
go run ./src/... --sim
