# PhoenixValidation — Deterministic Testing Layer

> Replay validation, fuzzing, chaos engineering, and invariant testing for the PhoenixOS ecosystem.

## Overview

PhoenixValidation provides the testing and verification backbone for PhoenixOS. It implements deterministic replay validation, fuzzing for edge cases, chaos engineering for resilience testing, and invariant verification for critical system properties.

**All critical code paths must have corresponding validation tests.**

## Repository Structure

```
PhoenixValidation/
├── replay/                 # Deterministic replay engine
│   ├── engine.go           # State reconstruction from events
│   ├── engine_test.go      # Replay engine tests
│   └── authority.go        # Hash comparison for drift detection
├── determinism/            # Determinism verification
│   └── determinism_test.go # Replay, crossrun, ordering tests (stubs)
├── evidence/               # Evidence chain verification
│   ├── chain_integrity_test.go  # Chain integrity (stub)
│   ├── fork_reject_test.go      # Fork rejection (stub)
│   └── mutation_block_test.go   # Mutation blocking (stub)
├── formal/                 # Formal invariant tests
│   ├── hash_chain_test.go       # Hash chain integrity (stub)
│   ├── ledger_invariant_test.go # Ledger invariant (stub)
│   ├── replay_consistency_test.go # Replay consistency (stub)
│   └── rollback_consistency_test.go # Rollback consistency (stub)
├── proofs/                 # Formal proofs (stubs)
│   ├── ledger_proof.go
│   ├── ordering_proof.go
│   ├── transition_proof.go
│   ├── rollback_proof.go
│   └── replay_identity_proof.go
├── kernel/                 # Kernel-level tests
│   ├── runtime_ring_test.go
│   └── ebpf_ring_stress_test.go
├── security/               # Security tests
│   └── thermo_test.go
├── unit/                   # Unit tests
│   └── registry_test.go
├── chaos/                  # Chaos engineering
│   └── fuzz_test.go
├── soak/                   # Long-running tests
│   └── drift_test.go
├── invariants/             # Invariant tests
│   └── clock_test.go
├── integration/            # Integration tests
├── runtime_graph/          # Runtime graph tests
├── truth/                  # Truth verification tests
├── proof/                  # Proof verification tests
├── distribution_test.go
├── infrastructure_test.go
├── go.mod
└── go.sum
```

## Core Principles

1. **Deterministic Replay**: Same input must produce same output
2. **Formal Proofs**: Critical properties must have formal verification
3. **Chaos Testing**: System must survive fault injection
4. **Soak Testing**: Long-running stability verification

## Build & Test

```bash
# Run all tests
go test ./...

# Run with race detector
go test -race ./...

# Run fuzzing
go test -fuzz=. ./chaos/...

# Run soak tests (manual)
go test -v ./soak/... -run TestDrift -timeout 24h
```

## Dependencies

- **Depends on**: PhoenixCore (contracts, ledger), PhoenixGuard (FSM), PhoenixKernel (eBPF)
- **Depended by**: All repositories (testing)

## Invariants

- Replay must produce identical hash chains
- No deterministic drift allowed
- FSM transitions must follow the ladder
- Ledger must be append-only

## License

PhoenixValidation is part of the PhoenixOS ecosystem.
