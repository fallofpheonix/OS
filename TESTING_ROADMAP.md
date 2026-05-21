# SentinelOS/PhoenixOS Stage-Gated Testing Roadmap

## Overview
This document defines the mandatory testing stages for SentinelOS/PhoenixOS. No module or feature may proceed to the next stage until the current gate is passed.

**Rule:** `No stage passes ↓ next stage blocked`

---

## Stage 0: Research Completion (0 to 5%)
**Goal:** Validate ideas before coding.
**Tests:**
- Theory consistency
- Dependency order
- Missing experiments
- RFC links
- Math validity
- Physics mapping
- Game mapping

**Artifacts:** `research_review.md`, `dependency_graph.md`, `risk_report.md`
**Gate:**
- **PASS:** Theory connected.
- **FAIL:** Isolated idea.

---

## Stage 1: Prototype Ready (5 to 15%)
**Modules:** Telemetry, Graph, Entropy, Event bus.
**Tests:** Unit tests.
**Workflow:** `mock input` -> `module` -> `output`
**Validation:** Schema, state transitions, serialization, memory leaks.
**Requirement:** 3 runs minimum.

---

## Stage 2: Collector Phase (15 to 25%)
**Goal:** Test telemetry.
**Targets:** Process capture, rename events, network, filesystem, lineage.
**Tests:** Event loss, ordering, latency, duplicates, overflow.
**Metrics:**
- CPU < 5%
- Loss < 2%
- Latency < 20ms
**Stop Condition:** Loss > 5%.

---

## Stage 3: Replay Phase (25 to 35%)
**Goal:** Build datasets and ensure determinism.
**Targets:** Normal, backup, rename storm, crypto writes, SMB spread.
**Tests:** Determinism, reproducibility, artifact hashes.
**Requirement:** 3 identical runs.
**Outputs:** `replay.json`, `diff.json`, `hashes.sha256`

---

## Stage 4: Detection Layer (35 to 50%)
**Modules:** Entropy, Graph, Fast path, Importance score.
**Tests:** Precision (TP, FP, TN, FN).
**Metrics:**
- Precision > 90%
- Recall > 85%
**Constraint:** Rule-based only (No AI yet).

---

## Stage 5: Evidence Layer (50 to 60%)
**Goal:** Prove the reasoning behind actions.
**Tests:** Ledger, trace hash, timeline, proof chain.
**Workflow:** `event` -> `trace` -> `decision` -> `action`
**Validation:** Reason exists for every action.
**Rejection:** Action without evidence.

---

## Stage 6: Control Layer (60 to 70%)
**Goal:** Validate response mechanisms.
**Tests:** Observe, limit, freeze, isolate.
**Metrics:** Reaction time, overshoot, stability.
**Requirement:** Cooldown, rollback, recovery.
**Avoidance:** Immediate "kill" without analysis.

---

## Stage 7: Physics Layer (70 to 78%)
**Goal:** Connect system state to physical models.
**Tests:** Entropy, energy, disorder, temperature.
**Validation:** State evolution.
**Experiment:** Compare normal vs attack trajectories.
**Rejection:** Unused physics (every metric must affect runtime).

---

## Stage 8: Game Layer (78 to 85%)
**Goal:** Strategic adaptation.
**Tests:** Payoffs, equilibrium, adaptation.
**Workflow:** Attacker move -> Defender move.
**Validation:** Stable strategy.
**Rejection:** Game logic without telemetry.

---

## Stage 9: AI Layer (85 to 92%)
**Goal:** Intelligent augmentation.
**Tests:** Explainability, drift, confidence, replay accuracy.
**Requirement:** Evidence ledger must be accessible by AI.
**Rejection:** Black box action.

---

## Stage 10: Kernel Layer (92 to 97%)
**Goal:** Move logic to kernel space safely.
**Tests:** Hooks, scheduler, memory, runtime.
**Metrics:** Panic count, CPU, latency, deadlock, starvation.
**Environment:** QEMU, VM, Sandbox (Never on host).

---

## Stage 11: Pilot Stage (97 to 100%)
**Goal:** Real-world validation.
**Scenarios:**
- Hospital: Rename burst, entropy, SMB.
- Kubernetes: Escape, secret leak.
- OT: Modbus, PLC.
**Metrics:** MTTD, MTTR, Precision, Recall, Loss.

---

## Current Status & Priorities
- [x] **Lineage memory tests** (Stage 2/3): Completed. Verified 100k nodes in 16.4MB, <7ms traversal.
- [x] **Replay harness** (Stage 3): Completed. Verified determinism with 200k events (Hash: `d992fd79...`).
- [x] **Evidence ledger validation** (Stage 5): Completed. Verified cryptographic chain integrity for reasoning.

---

## Master Testing Matrix Status
| Stage | Description | Status |
|---|---|---|
| 0-5% | Research | PASS |
| 5-15% | Unit | PASS |
| 15-25% | Telemetry | PASS |
| 25-35% | Replay | PASS |
| 35-50% | Detection | IN PROGRESS |
| 50-60% | Evidence | PASS |
| 60-70% | Control | BUILDING |
| 70-78% | Physics | PENDING |
| 78-85% | Game | PENDING |
| 85-92% | AI | PENDING |
| 92-97% | Kernel | PENDING |
| 97-100% | Pilot | PENDING |
