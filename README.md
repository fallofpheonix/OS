# Phoenix Matrix OS

Phoenix Matrix is a security-control operating substrate designed for deterministic system protection and verifiable autonomous actuation.

## Current System Status

**Vertical Slice: PROVEN**

The core architectural pipeline from kernel telemetry to durable storage is now functional and verified.

### Current Proven Capabilities (Vertical Slice)
- **Persist to disk:** Verified via `TestPersistorWriteAndRead`.
- **Crash recovery:** Verified via `TestPersistorCrashRecovery` (partial line detection).
- **Replay integrity:** Verified via `TestReplayerHeadHashContinuity`.
- **Cross-process survival:** Verified via `TestEndToEndPersistenceAndRecovery`.
- **Deterministic Math:** Fixed-point arithmetic used in all security and physics paths.
- **Authenticated Identity:** Ed25519-signed events and validator set management.
- **P2P Gossip:** libp2p-based propagation of authenticated state transitions.

### Not Yet Implemented (Gaps)
- **TOKEN-WARDEN-001 authentication:** Root authority token is generated in-memory but needs full CLI integration.
- **eBPF -> Bus wiring:** Connected in bootstrap but requires root/kernel for live production polling.
- **Consensus Voting:** Quorum certificates defined but full BFT voting loop (PreVote/PreCommit) is planned.
- **Slashing:** Punitive logic for Byzantine behavior is specified but not yet implemented.

## Core Mandates
1. **Determinism is sacred.** No float64 in state-transition or security paths.
2. **Replay is authoritative.** Disk-backed ledger is the single source of truth.
3. **Fail-Closed.** System restricts I/O on integrity violation or connectivity loss.

## Usage

```bash
# Start the Phoenix Matrix Daemon
./platform/os/phoenixd
```

## Documentation
- [Consensus Specification](docs/CONSENSUS.md)
- [Ledger Implementation](foundation/ledger/docs/README.md)
- [System Invariants](foundation/runtime/INVARIANTS.md)

---
*Authorized by Phoenix Sovereign Governance*
