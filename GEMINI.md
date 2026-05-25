# PhoenixOS: Core Instructions & Architecture

This file contains foundational mandates for the PhoenixOS project.

## 1. Project Identity
**Project Name:** PhoenixOS
**Definition:** Security-Control Operating Substrate. A deterministic security runtime and telemetry platform for autonomous system protection.
**Core Philosophy:** Security as a thermodynamic state. Autonomous "quenching" of disorder via the Phoenix Matrix.

## 2. Six Immutable Axioms
1. **Determinism is sacred.** No reliance on non-deterministic primitives (unordered maps, race-dependent ordering).
2. **Replay is authoritative.** Replay governs semantics; logs, metrics, and AI-outputs are secondary.
3. **AI is advisory.** AI informs, but never directly controls kernel or actuation FSM.
4. **Control must remain bounded.** Actuation is rate-limited, isolated, and reversible.
5. **Telemetry correctness > AI sophistication.** Precise, replayable telemetry is the foundation.
6. **Never scale instability.** Single-node stability must be achieved before distributed scaling.

## 3. The 7-Layer Stack (Phoenix Matrix)
1. **L7: Swarm Coordination (Phoenix Nexus):** Distributed consensus (PoA + Reputation).
2. **L6: System Physics (Phoenix Sentinel):** Thermodynamic SDI monitoring.
3. **L5.5: Strategic Policy (Phoenix Arbiter):** Game-theoretic policy (Stackelberg).
4. **L5: Actuation & Control (Phoenix Warden):** Finite-State Controller (SAFE -> COMPROMISED).
5. **L4: Graph Intelligence (Phoenix Trace):** Causal lineage DAGs with 3-tier storage.
6. **L3: Telemetry Math (Phoenix Monitor):** Signal processing (Entropy + Kalman).
7. **L2: Kernel Runtime (Phoenix Kernel):** eBPF probes.
8. **L1: Platform Integrity (Phoenix Guard):** <100ms Fast-Path enforcement.

## 3. Mandatory Engineering Rules (Expert Refined)
- **Fast-Path Priority:** High-confidence heuristics (Entropy > 7.9) MUST bypass strategic layers via Phoenix Guard.
- **Evidence-First Actuation:** Every action MUST be recorded in the Phoenix Ledger with a SHA-256 hash-chain.
- **State-Aware Control:** Do NOT map SDI directly to PID gains. Use the Finite-State Controller (SAFE/WATCH/SUSPICIOUS/CRITICAL/COMPROMISED).
- **Lineage Retention:** Use 3-tier storage (HOT/WARM/COLD). Never prune `init`, `auth`, `kernel`, or `systemd` nodes.
- **Archive Integrity:** No code restored from `15_archive/` without a entry in `02_docs/02_integration/ARCHIVE_IMPORT.md`.

## 4. Maturity & Roadmap (Stages A-D)
Current Maturity: **Runtime Security Research Platform**.

### Stage A: Hardening (CURRENT FOCUS)
- [X] **Race Detection:** Implement and run Go race detector across all core services (VERIFIED).
- [ ] **Fuzz Testing:** Targeted fuzzing of Ledger ingestion and Warden FSM triggers.
- [ ] **Chaos Engineering:** Inject telemetry jitter and out-of-order events to verify TCS window stability.
- [ ] **Replay Stress:** Verify 100% hash-matching across 1000+ varying execution traces.

### Stage B: Formal Invariants
- [ ] **Event Bus Guarantees:** Transactional event delivery and ordering proofs.
- [ ] **Ledger Invariants:** Formal validation of hash-chain integrity and evidence non-repudiation.
- [ ] **FSM Proof:** TLA+ model of the Warden FSM to prevent oscillation loops.

### Stage C: OS Primitives (Future)
- [ ] **Syscall Monitor:** Move from generic eBPF probes to a structured syscall boundary.
- [ ] **Sandbox Runtime:** Isolate actuation effects within restricted containers.
- [ ] **Process Graph Engine:** Real-time causal lineage DAG construction.

### Stage D: Distributed Coordination (L6/L7)
- [ ] **PoA Consensus:** Distributed trust and reputation.
- [ ] **Replicated Ledger:** Multi-node evidence synchronization.

## 5. Completed P0 Foundations & Hardening
- [X] **Phoenix Ledger:** Verifiable Evidence Chain, Ledger V2 state transition validations.
- [X] **Phoenix Guard:** Kernel-level Fast Path (<100ms) with Priority Ingestion lanes.
- [X] **Phoenix Trace Storage:** 3-tier lifecycle management.
- [X] **Phoenix Warden:** Stable state controller with dwell limits and cooldowns.
- [X] **Phoenix Arbiter:** Cost-aware decision engine with Evidence Weighting and Counterfactual Analysis (Track C).
- [X] **TCS Telemetry Window:** Bounded sliding window with dynamic evaluation.
- [X] **Logical Clock Standardization:** Monotonic logical tick counter standardized.

## 6. Documentation Architecture (12-Layer Mandate)
PhoenixOS adheres to a 12-layer documentation structure in `02_docs/` to ensure traceability and auditability.
- **L0: Governance** (`00_governance/`): Vision, Roadmap, Status, Decisions.
- **L1: Architecture** (`01_architecture/`): System maps, Component specs, RFCs.
- **L2: Integration** (`02_integration/`): External repo tracking and ARCHIVE_IMPORT.md.
- **L3: Agents** (`03_agents/`): AI Agent registries and behavior protocols.
...
- **L11: Emergency** (`11_emergency/`): Disaster recovery and Safe modes.

**Mandatory Rule:** No agent edits code without updating the related documentation. Every PR must update architecture, dependency, threat, and test impact documents.

## 7. Active Phase: Phase F1 (Runtime Development)
**Current Goal:** Implement the causal replay and evidence chain in the live runtime.
**Status:** F0 CLOSED, F1 OPEN.
**Verified Invariants:** Determinism, Replay, Truth (Merkle-DAG), Race Safety.


