---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Runtime — Build Instructions

> Last verified: 2026-06-04

## Building the Go Module

```bash
cd foundation/runtime
go build ./...
```

## Loading eBPF Probes (requires root privileges)

The eBPF module compiles kernel probes using Clang/LLVM. Ensure kernel headers are installed:

```bash
# Build eBPF tools
cd foundation/runtime/kernel
make
```
