---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# PhoenixOS Troubleshooting Guide

## Build and Compilation Issues

### Duplicate Module / Package Conflicts
- **Problem**: Running `go work sync` or building returns errors about duplicate modules or overlapping packages.
- **Cause**: Some nested subdirectories or legacy directories (e.g. `PhoenixStimulation` vs `PhoenixSimulation`) might declare the same module path (like `github.com/fallofpheonix/phoenix-os`).
- **Solution**: Keep standalone simulation suites out of the global `go.work` file if they declare conflicting module paths, or adjust their `go.mod` module declarations to reflect their specific path.

### "Not Enough Arguments in Call to AddEntryV2"
- **Problem**: Tests in `PhoenixCore/ledger` fail to compile with argument mismatches on `AddEntryV2`.
- **Cause**: The ledger V2 signature was updated to support 7 arguments (adding `traceHash` as the 4th argument) but test files were not fully refactored.
- **Solution**: Insert the empty string `""` as the `traceHash` argument (4th position) in `concurrency_test.go` and `ledger_v2_test.go`.

### Contract Drift Between EventEnvelope and Event
- **Problem**: Validation and replay tests pass incompatible event types across the boundary.
- **Cause**: The event contract was split between protocol envelopes and internal event models without a stable adapter.
- **Solution**: Introduce a versioned adapter and make contract tests the gate for replay and validation changes.

---

## Test Failures

### "Boot Non-Determinism Detected" in `boot_test.go`
- **Problem**: `TestBootReproducibility` fails with mismatching genesis checksums.
- **Cause**: `boot.NewSubsystemInfo` assigns a high-resolution timestamp using `time.Now().UnixNano()`, which differs between two consecutive system initializations.
- **Solution**: Override the `Timestamp` field of the captured `SubsystemInfo` structs to a static/deterministic value inside the test function before calculating the checksum.

### Race Detection Warnings
- **Problem**: Running `go test -race` reports a data race inside `ebpf_ring_stress_test.go`.
- **Cause**: The bus `OnOverflow` callback runs asynchronously inside a new goroutine, causing concurrent read/write operations on the `overflowTriggered` boolean flag.
- **Solution**: Use `sync/atomic` methods (`atomic.StoreInt32`, `atomic.LoadInt32`) to set and check the flag, and insert a small retry loop to wait for the async execution to finish.
