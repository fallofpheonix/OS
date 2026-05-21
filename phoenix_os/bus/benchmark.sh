#!/bin/bash
echo "Benchmarking Phoenix Bus..."
go test -bench=. ./src/*.go > artifacts/benchmark.log
echo "Benchmarks complete."
