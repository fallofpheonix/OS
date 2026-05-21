#!/bin/bash
echo "Running Benchmarks..."
go test -bench=. -benchmem ./src/... > artifacts/benchmark.log
echo "Benchmarks complete."
