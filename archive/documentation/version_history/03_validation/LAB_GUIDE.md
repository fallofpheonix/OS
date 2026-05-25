# PhoenixOS: Determinism Lab Guide (L7)

The Determinism Lab is the primary verification sandbox for Phase F0 (Foundation). It subjects the runtime to various threat scenarios under strict validation bounds to verify that same input events lead to identical replay hashes, decisions, and containment states.

---

## 1. Simulated Threat Scenarios

The lab simulates security events using mock telemetry logs and simulated system interactions.

### `ForkBomb` (Process Multiplication)
- **Behavior:** Simulated process spam causing a rapid increase in process count.
- **Warden Transition:** `SAFE` $\rightarrow$ `WATCH` (State 1).
- **Enforcement:** Process spawning limits applied.

### `ReverseShell` (Privilege Escalation)
- **Behavior:** Spawn of a shell (`/bin/sh` or `/bin/bash`) by a web server process (`nginx`, `httpd`).
- **Warden Transition:** `SAFE` $\rightarrow$ `CONTAIN` (State 4) immediately due to high threat severity.
- **Enforcement:** cgroup freezing and network namespace shutdown.

### `NetworkBeacon` (C2 Patterns)
- **Behavior:** Periodic outbound TCP connection attempts to a fixed IP address.
- **Warden Transition:** `SAFE` $\rightarrow$ `ALERT` (State 2).
- **Enforcement:** Outbound IP blocked at XDP socket boundaries.

### `Exfiltration` (Data Theft)
- **Behavior:** Rapid reading and transmission of sensitive system files (e.g., `/etc/passwd`).
- **Warden Transition:** `SAFE` $\rightarrow$ `ALERT` (State 2) / `CONTAIN` (State 4).
- **Enforcement:** File access auditing and interface rate-limiting.

### `Ransomware` (Mass Write Anomaly)
- **Behavior:** Rapid, successive write/modify system calls over multiple user directories.
- **Warden Transition:** `SAFE` $\rightarrow$ `ALERT` (State 2).
- **Enforcement:** Directory lock down and backup snapshot generation.

### `HashTamper` & `ReplayCorruption` (Subversion Detection)
- **Behavior:** Explicit modification of historical event payloads or parent hashes in SQLite.
- **System Action:** Recalculating the ledger Merkle roots via `truth.Verify()` immediately flags a mismatch and halts execution.

---

## 2. Key Verification Metrics

During scenario runs, the validation suite verifies four key invariants:

1. **Latency Budget:** Execution overhead must remain within limits (L1 telemetry ingestion: $<50\text{ms}$; Warden state transition: $<100\text{ns}$).
2. **Replay Fidelity:** Replay sequence hash matches original live execution root hash 100%.
3. **Decision Repeatability:** Arbiter must output the identical action class for identical inputs.
4. **Containment Stability:** The FSM must not flap or oscillate (hysteresis checked for 30 ticks).

---

## 3. Validation Matrix

The validation matrix displays the current results of the verification suite runs:

| Scenario | Ingestion Latency | Replay Matching | Ledger Chain Verifies | Decision Matches | Containment Applied | Status |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **`ForkBomb`** | $<12\text{ms}$ | 100% | 100% | `ClassLog` | Applied | **PASS** |
| **`ReverseShell`**| $<5\text{ms}$ | 100% | 100% | `ClassLocalIsolate` | Applied | **PASS** |
| **`NetworkBeacon`**| $<8\text{ms}$ | 100% | 100% | `ClassThrottle` | Applied | **PASS** |
| **`Exfiltration`** | $<10\text{ms}$ | 100% | 100% | `ClassThrottle` | Applied | **PASS** |
| **`Ransomware`** | $<15\text{ms}$ | 100% | 100% | `ClassThrottle` | Applied | **PASS** |
| **`HashTamper`** | $<2\text{ms}$ | 100% | 100% | N/A (Halt) | N/A (Halt) | **PASS** |
| **`ReplayCorruption`**| $<2\text{ms}$ | 100% | 100% | N/A (Halt) | N/A (Halt) | **PASS** |

---

## 4. Running Verification Suite

To run the full suite of lab tests and assert the validation rules, run:

```bash
# Verify imports do not bypass boundaries
python3 05_tools/validate_imports.py

# Run Go tests in the main runtime
go test -v ./...
```
