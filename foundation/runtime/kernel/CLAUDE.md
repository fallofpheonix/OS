# PhoenixKernel — Kernel Integration Layer

## Agent Skills
### Issue Tracker
GitHub issue tracker. See `docs/agents/issue-tracker.md`.

### Triage Labels
Default triage label vocabulary. See `docs/agents/triage-labels.md`.

### Domain Docs
Multi-context layout. See `docs/agents/domain.md`.

## Build & Test
```bash
# Build eBPF (requires Linux with clang)
cd src && make

# Run Go tests
go test ./...

# Run with race detector
go test -race ./...
```

## Architecture
PhoenixKernel provides the kernel interface for PhoenixOS. It loads eBPF programs into the Linux kernel, attaches to syscall tracepoints, and provides the enforcement hooks for process isolation.

## Key Components
- **ebpf_loader.go** — eBPF program loader and ring buffer reader
- **enforcer.go** — Ring buffer → Bus bridge
- **types.go** — TelemetryEvent and EventPublisher interfaces
- **src/phoenix_exec.c** — eBPF C programs (tracepoint + LSM hooks)
- **runtime/** — Go runtime (namespaces, affinity, probe injection)
- **sandbox/** — Kernel simulator for testing
- **hooks/** — eBPF enforcer and isolation hooks
- **live/** — Live probe management (stubs)

## Invariants
- eBPF programs must be loaded from verified paths
- Ring buffer reader must handle transient errors gracefully
- Namespace operations require CAP_SYS_ADMIN
- All kernel events must be published to the Bus
