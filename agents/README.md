# PhoenixOS: Internal Agents (Layer 3)

## Purpose
The `agents` module implements the core autonomous security response loop for PhoenixOS. It transforms raw telemetry into physics-based risk models, game-theoretic classifications, and PID-controlled actuations.

- **Validation Gates**: All core agents (Telemetry, Graph, Physics, Game, Control, Forensics, Kernel) have been implemented with robust interfaces and verified via a comprehensive end-to-end integration test.

## Core Agents
- **TelemetryAgent**: Ingests and normalizes process-level events. Interface: `Start()`, `Stop()`, `RecordEvent()`, `GetLineage()`.
- **GameAgent**: Uses recursive Bayesian updates. Interface: `UpdateBeliefs()`, `SolveBestStrategy()`, `GetBeliefs()`.
- **GraphAgent**: Maintains a causal lineage DAG with Multi-Factor Importance Scoring (Centrality, Criticality, Spread, Depth).
- **PhysicsAgent**: Calculates SDI and Threat Temperature ($\theta_T$) based on disorder and node importance.
- **ControlAgent**: Implements PID control with discrete containment states (LevelObserve to LevelKill).
- **ForensicsAgent**: Captures cryptographic snapshots for post-incident verification.
- **KernelAgent**: Safety-locked interface for kernel-level mitigations.

## Status
- [x] Phase 1: Core Communications & Monitoring (Completed)
- [x] Layer 3 (Internal Agents) Foundations (Completed)
- [x] End-to-End Pipeline Validation (Completed)
- [ ] Layer 4 (Sentinel Orchestration) Transition (In Progress)

## Structure
- `internal/`: Package-level implementations of individual agents.
- `src/`: Main simulation entry point and event loop.
- `tests/`: Integration tests for the full Telemetry-to-Kernel pipeline.
- `bench/`: Performance benchmarks for agent throughput and latency.
- `replay/`: Deterministic telemetry replay for validation.
- `artifacts/`: Build outputs and forensic snapshots.

## Performance Budget
- **Event Processing Latency**: < 5ms per agent hop.
- **Total Loop Latency**: < 50ms from Telemetry to Control Actuation.
- **Memory Overhead**: < 100MB for the Graph DAG (hot tier).

## Validation Gates
- [x] Build success
- [x] Unit test pass (all packages)
- [x] Architectural Boundary Verification (no illegal imports)
- [x] Integration Pipeline Validation (Ransomware simulation)
- [x] Performance within budget (Measured: ~0.024ms per loop)
- [x] Replay deterministic output (Continuous Validation)
