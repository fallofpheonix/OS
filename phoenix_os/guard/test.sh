#!/bin/bash
echo "Testing Phoenix Guard..."
# Guard doesn't have a test file yet, so we just check if it compiles/runs
go run src/guard.go
echo "Tests complete."
