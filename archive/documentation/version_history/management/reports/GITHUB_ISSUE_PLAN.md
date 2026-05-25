# Unified Repository & Feature Integration: Roadmap to F1 Maturity

## 1. Vision & Strategic Rationale (The "Why")
PhoenixOS has transitioned from a fragmented collection of research modules and integrated experiments into a stabilized deterministic runtime. To achieve **F1 Maturity (Runtime Hardening)**, we must unify the codebase.

**Rationale:**
- **Architectural Clarity:** Eliminate the boundary between "integrated" repositories and the core runtime.
- **Noise Reduction:** Move remaining speculative research (Quantum, generic AI) into deep isolation to prevent architectural drift.
- **Single Source of Truth:** Establish a single package hierarchy that aligns with the 7-layer stack (L1 Guard to L7 Nexus).
- **Verified Determinism:** Ensure every module, including external logic, adheres to the sacred axioms of replayability and immutability.

---

## 2. Current State Assessment (The "What")
Following the Phase 0 Master Audit, the following milestones are verified:
- [x] **Blockers Removed:** Security compile failures and tooling import gaps resolved.
- [x] **Deterministic Substrate:** 1,000+ rollback stress runs passing with 0 divergence.
- [x] **Userspace Replay:** 100% precision achieved on 200k event traces.
- [x] **Topology Integrity:** Zero illegal dependency paths detected in the core flow.

**Post-Audit Reality:**
- **F0 Foundation:** `CONDITIONALLY COMPLETE` (Substrate stable, pending formal F2 proofs).
- **F1 Runtime:** `PREPARED` (Scaffolding exists, pending eBPF/XDP integration).

---

## 3. Merges & Unification Map (The "How")

### 3.1 Integrated Module Migration
We will move code from `03_repositories/integrated/` directly into the `phoenix_os/` hierarchy to simplify imports and build cycles.

| Original Component | New Core Target | Purpose |
|---|---|---|
| `phoenix-control/warden` | `phoenix_os/warden` | L5 Finite-State Actuation |
| `phoenix-control/arbiter` | `phoenix_os/arbiter` | L5.5 Game-Theoretic Policy |
| `phoenix-logic/monitor` | `phoenix_os/monitor` | L3 Signal Processing |
| `phoenix-logic/truth` | `phoenix_os/truth` | L2 Immutable Ledger |
| `phoenix-runtime/guard` | `phoenix_os/guard` | L1 Fast-Path Enforcement |

### 3.2 Feature Integration Detail
1. **Warden FSM Hardening:** Integrate the `DwellTicks` hysteresis barrier globally to prevent state oscillation.
2. **Arbiter Cost-Policy:** Implement the GS (Severity) and FA (Frequency) importance formulas as the primary gate for actuation.
3. **Kernel Probes (F1 Stage):** Transition from simulated `05_tools/telemetry/replay` inputs to real-time `10_kernel/sandbox` and eBPF ring buffers.

---

## 4. Implementation Steps & Gates

### Step 1: Semantic Package Alignment (Week 1-2)
- [ ] Move `integrated/` source files to `phoenix_os/`.
- [ ] Rename packages to follow the flat `phoenix/layer` convention (e.g., `package warden`).
- [ ] Update `go.work` and `go.mod` to remove local replacement hacks.
- [ ] **Gate:** `go test ./...` must pass with zero modifications to existing test logic.

### Step 2: Enforcement of Forbidden Paths (Week 3)
- [ ] Integrate the `check_illegal_deps.py` logic into the CI/CD pipeline.
- [ ] Explicitly block any `telemetry` -> `warden` imports.
- [ ] **Gate:** Build failure if any illegal dependency is introduced.

### Step 3: Formal Model Synchronization (Week 4)
- [ ] Align the Go implementation of Warden with the TLA+ models in `F2_PREP/`.
- [ ] Implement `rollback_proof.md` requirements: state parity checks on every restore.
- [ ] **Gate:** Successful execution of `tests/proofs/` battery.

---

## 5. Risk Matrix & Mitigations

| Risk | Mitigation |
|---|---|
| **Circular Dependencies** | Strictly enforce the 7-layer hierarchy via automated import scanning. |
| **Logic Drift** | Use the verified `REPLAY_IDENTITY_REPORT.md` hashes as the baseline for every merge. |
| **Research Pollution** | Keep `06_research/` as a "sandbox only" area; no research code allowed in `phoenix_os/`. |

---

## 6. Detailed Engineering Comments (For Implementation)

**Import Restructuring:**
All imports should be migrated from:
`"github.com/fallofpheonix/phoenix-os/phoenix_os/ai"`
To:
`"phoenix/ai"`
This relies on the established `go.work` configuration and ensures that internal core components treat each other as local modules.

**Warden Actuation Interface:**
The `Actuate` function signature in `warden.go` must be the **only** way for the Arbiter to trigger system state changes. 
```go
func (w *Warden) Actuate(target SystemState, class ActuationClass, confidence float64, ...)
```
Any attempt to modify `w.State` directly from outside the `warden` package is a P0 violation.

**Ledger Immuntability:**
The `TruthLedger` in `phoenix_os/truth` must remain the authoritative source for replay. Any feature that requires history (Arbiter's Frequency Mode, Recovery Snapshots) must pull data from the Ledger hash-chain.
