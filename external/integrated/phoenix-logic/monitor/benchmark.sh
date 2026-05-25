#!/bin/bash
set -e
mkdir -p artifacts
echo "Running Benchmarks..."
go test -bench=. -benchmem -count=3 ./... | tee artifacts/benchmark.log
echo "Benchmarks complete."
