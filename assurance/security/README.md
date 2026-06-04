# PhoenixGuard — Security Enforcement Layer

> Fast-path enforcement, bounded actuation, and the Warden FSM for the PhoenixOS ecosystem.

## Overview

PhoenixGuard is the security enforcement layer of PhoenixOS. It implements the Warden FSM (finite state machine) that governs system security states, provides bounded execution harnesses for actuation, and manages the kill switch emergency mechanism.

**All enforcement actions MUST go through PhoenixGuard. No direct kernel access is allowed.**

## Repository Structure

```
PhoenixGuard/
├── engine/              # Warden FSM enforcement engine
│   ├── warden.go        # Core FSM: SAFE→WATCH→SUSPICIOUS→CRITICAL→COMPROMISED
│   └── warden_test.go   # FSM transition tests
├── emergency/           # Kill switch mechanism
│   ├── killswitch.go    # Thermodynamic failsafe
│   └── killswitch_test.go
├── actuation/           # Bounded execution harness
│   ├── executor.go      # Timeout + rollback execution
│   ├── executor_test.go
│   └── sandbox.go       # Process isolation actuator
├── policies/            # Trust matrix
│   ├── trust_matrix.go  # Cross-domain access control
│   └── trust_matrix_test.go
├── audit/               # Violation logging
│   ├── violation_log.go
│   ├── violation_log_test.go
│   └── jsonl_writer.go
├── validation/          # Security validators (stubs)
│   └── security_validator.go
├── runtime/             # Python runtime (filesystem, orchestration, shell)
│   ├── filesystem/      # Sandboxed filesystem operations
│   ├── orchestration/   # Control-plane execution
│   └── shell/           # Shell command execution
├── core/                # Orchestrator
│   └── orchestrator.py
├── interfaces/          # CLI interface
│   └── cli/
├── infrastructure/      # Logging
│   └── logging/
├── tests/               # Integration and runtime tests
│   ├── integration/
│   └── runtime/
├── docs/                # Architecture and specification documents
│   ├── ARCHITECTURE.md
│   ├── INVARIANTS.md
│   └── ...
├── control.go           # Control FSM bridge
├── go.mod
└── go.sum
```

## Core Principles

1. **Warden FSM**: Strict state ladder — no skipping states
2. **Bounded Execution**: All actions have timeout and mandatory rollback
3. **Kill Switch**: One-way emergency halt — no disengagement without restart
4. **AI Disconnected**: Warden operates independently of AI systems
5. **Audit Trail**: Every state transition is recorded with evidence

## Build & Test

```bash
# Run all tests
go test ./...

# Run with race detector
go test -race ./...
```

## Dependencies

- **Depends on**: PhoenixCore (contracts, bus), PhoenixKernel (runtime)
- **Depended by**: PhoenixMind (orchestrator), PhoenixOS (integration)

## Invariants

- FSM transitions must follow the strict state ladder
- All actuations must have a rollback plan
- Kill switch is irreversible within a process lifetime
- No AI system can directly trigger state transitions

## License

PhoenixGuard is part of the PhoenixOS ecosystem.
