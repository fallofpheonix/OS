# PhoenixOS: Core Instructions & Architecture

This file contains foundational mandates for the PhoenixOS project.

## 1. Project Identity
**Project Name:** PhoenixOS
**Definition:** Deterministic Cybernetic Security Runtime (running on Linux).
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

## 4. Completed P0 Foundations
- [X] **Phoenix Ledger:** Verifiable evidence chain.
- [X] **Phoenix Guard:** Kernel-level Fast Path (<100ms).
- [X] **Phoenix Trace Storage:** 3-tier lifecycle management.

## 5. Active Phase: Phase 1 (State & Intelligence)
**Current Goal:** Implement Importance Scoring ($S_I$) in Phoenix Arbiter.
