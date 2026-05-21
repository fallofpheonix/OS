#!/bin/bash
set -e
mkdir -p debug
echo "Running __MODULE_NAME__ in debug mode..."
go test -v -race -run=. ./... 2>&1 | tee debug/trace.log
echo "Debug trace saved to debug/trace.log"
