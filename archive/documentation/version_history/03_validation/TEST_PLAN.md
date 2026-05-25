# Test Plan: PhoenixOS Validation

## 1. Testing Strategy

| Level | Target | Tooling | Frequency |
| :--- | :--- | :--- | :--- |
| **Unit** | Individual L1-L7 logic | Go Test / PyTest | Every Commit |
| **Integration** | Cross-layer communication | Docker Compose | Nightly |
| **System** | Full OS runtime | QEMU / Bare Metal | Weekly |
| **Chaos** | Random component failure | Custom Chaos Agent | Monthly |
| **Stress** | High-load telemetry | Synthetic Load Gens | Release Candidate |
| **Fault Injection**| Malicious probe injection | Phoenix Guard Tests | Bi-Weekly |

## 2. Validation Targets
- **Deterministic Replay:** Replay a known attack trace and ensure the exact same state transitions occur.
- **Latency Bounds:** Verify <100ms response for Fast-Path triggers.
- **Ledger Integrity:** Tamper with a ledger entry and verify detection by the hash-chain validator.
- **State Stability:** Ensure Warden does not enter infinite oscillation between states.
- **Timeline Split Detection:** Ensure divergent process action chains generate different head hashes (`TestReplaySplit`).
- **Rollback Forking Resistance:** Ensure rollback and re-entry on the ledger creates a new timeline with a distinct hash chain (`TestHashFork`).
- **Cross-Run Identity Consistency:** Ensure snapshot hashes are byte-for-byte identical across 1000 consecutive runs (`TestCrossRunReplayHash`).

