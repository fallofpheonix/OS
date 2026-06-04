---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Contracts — Build Instructions

> Last verified: 2026-06-04

## Building the Module

To build the contracts package independently:

```bash
cd foundation/contracts
go build ./...
```

Since the contracts package does not import other modules, this build relies entirely on the Go standard library.
