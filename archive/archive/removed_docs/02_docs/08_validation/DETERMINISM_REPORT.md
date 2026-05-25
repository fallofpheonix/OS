# PhoenixOS: Determinism Registry Report

This registry documents validation executions, hash matching percentages, drift counts, and pass status across all core runtime subsystems.

---

## 1. Determinism Execution Log

| Subsystem | Iterations | HashMatch% | DriftCount | ReplayMatch | Status | Verification Target |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **Rollback** | 1000 | 100% | 0 | YES | **PASS** | `TestRollbackDeterminism` |
| **Replay** | 1000 | 100% | 0 | YES | **PASS** | `TestCrossRunReplayHash` |
| **Recovery** | 100 | 100% | 0 | YES | **PASS** | `TestRecoveryRepeatability` |
| **State** | 100 | 100% | 0 | YES | **PASS** | `TestOutputDeterminism` |
| **FSM Rules** | 100 | 100% | 0 | YES | **PASS** | `TestTransitionInvariant` |
| **Fuzz/Parse** | 50 | 100% | 0 | YES | **PASS** | `TestFuzzPayloadParser` |

---

## 2. Validation Execution Parameters
- **Memory Race Checks:** Verified using `-race` flag during Go test compilation.
- **Order Randomization:** Verified via Go's `-shuffle=on` flag.
- **Trace Parity:** Asserts that replaying process, network, and file streams yields identical final global hashes.
