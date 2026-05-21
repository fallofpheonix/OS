#!/bin/bash
echo "Running Benchmarks..."
go test -bench=. ./src/*.go > artifacts/benchmark.log
echo "Benchmarks complete."
