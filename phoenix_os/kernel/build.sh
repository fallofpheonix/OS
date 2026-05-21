#!/bin/bash
set -e
echo "Building Phoenix Kernel Service..."
mkdir -p artifacts
# In a real environment: clang -O2 -target bpf -c ebpf/phoenix_trace.c -o artifacts/phoenix_trace.o
go build -o artifacts/phoenix_kernel ./src/loader.go
echo "Build complete."
