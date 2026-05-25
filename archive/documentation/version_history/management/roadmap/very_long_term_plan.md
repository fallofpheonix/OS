# PhoenixOS: Very Long-Term Master Plan

This document details the multi-year strategic roadmap for the **PhoenixOS** cybernetic security runtime, tracing the evolution from a hardened single-node userspace daemon to an autonomous swarm-based custom kernel runtime.

---

## 1. Architectural Stack Alignment (The 7-Layer Matrix)

Every long-term deliverable must map directly to a layer in the Phoenix Matrix:
1. **L7: Swarm Coordination (Phoenix Nexus):** Distributed consensus (Proof-of-Anomaly + Reputation).
2. **L6: System Physics (Phoenix Sentinel):** Thermodynamic SDI monitoring (Entropy & Kalman signals).
3. **L5.5: Strategic Policy (Phoenix Arbiter):** Game-theoretic decision engine (Stackelberg Game solvers).
4. **L5: Actuation & Control (Phoenix Warden):** 5-State Discrete FSM (SAFE -> WATCH -> SUSPICIOUS -> CRITICAL -> COMPROMISED).
5. **L4: Graph Intelligence (Phoenix Trace):** Causal process lineage DAGs with 3-tier storage (HOT/WARM/COLD).
6. **L3: Telemetry Math (Phoenix Monitor):** Shannon Entropy, KL divergence, and digital signal filters.
7. **L2: Kernel Runtime (Phoenix Kernel):** eBPF/XDP telemetry collectors and kernel space probes.
8. **L1: Platform Integrity (Phoenix Guard):** <100ms Fast-Path syscall enforcement maps.

---

## 2. Multi-Year Evolution Map (Phases I - V)

```
[Phase I: Hardening] -> [Phase II: eBPF & App] -> [Phase III: Graph & Swarm] -> [Phase IV: Strategic Policy] -> [Phase V: Custom Kernel]
      (Months 1-6)             (Months 6-12)             (Months 12-24)               (Months 24-36)              (Months 36+)
```

### Phase I: Stabilization, Hardening & Telemetry Normalization (COMPLETED)
*   **Architecture Layers:** L1, L3, L5, P0 (Ledger)
*   **Focus:** Core single-node runtime validation, concurrency safety, and telemetry normalization.
*   **Deliverables:**
    *   **Deterministic Sequencing:** Neutralized replay divergence using a deterministic allocator and sorted reorder window.
    *   **Priority Event Ingestion:** Dedicated Evidence Reserve Lane (85% occupancy threshold) and OnOverflow snapshots (95% occupancy threshold) to prevent ring-buffer starvation.
    *   **Warden Hysteresis & Concurrency Safety:** State dwell limits (30 ticks), stabilization cooldowns (10 ticks), de-escalation recovery budgets, and full mutex protection for concurrent SOC API actions.
    *   **Ledger V2 Schema & Thread Safety:** Cryptographically-linked transition validations using `StateBefore`/`StateAfter` nodes, and full RWMutex safety against background overflow logging races.
    *   **TCS Sliding Window Stability:** Telemetry event sequence scanning and negative SeqID filters to prevent underflow confidence failures.
*   **Acceptance Gate:** Replay equivalence tests pass; Merkle-DAG verification and thread-race validations confirm zero-leak integrity.

### Phase II: Real Telemetry, Kernel space & Immutable Appliance (Months 6–12)
*   **Architecture Layers:** L2, L1, L5
*   **Focus:** Transitioning from simulated logs to real system enforcement.
*   **Deliverables:**
    *   **eBPF Probes Integration:** Build modular eBPF probes for syscall interception, network socket bindings, and process creation hooks.
    *   **Minimal Appliance Boot:** Create a standalone bootable appliance using LinuxKit or custom BusyBox/Initrd. Build Warden as PID 1 to lock execution.
    *   **Cgroups Sandbox Actuator:** Implement Warden containment actions leveraging cgroups v2 to throttle or freeze suspicious process groups.
*   **Acceptance Gate:** eBPF probe overhead remains $<1\%$ CPU under a synthetic payload of 100k events/sec.

### Phase III: Graph Intelligence & Swarm Observability (Months 12–24)
*   **Architecture Layers:** L4, L7
*   **Focus:** Temporal causality tracing and multi-node event sync.
*   **Deliverables:**
    *   **3-Tier Database (HOT/WARM/COLD):** Graph engine optimized for memory safety (traversals $<7\text{ms}$ for 100k nodes).
    *   **Consensus Engine (Phoenix Nexus):** Proof-of-Anomaly (PoA) protocol using Reputation-Weighted voting.
    *   **Distributed Replay Federation:** Multi-node event replay validation utilizing Vector Clocks and causal logical clocks to synchronize distributed telemetry.
*   **Acceptance Gate:** DAG reconstruction fidelity $\ge 95\%$ on cross-node attack scenarios.

### Phase IV: Strategic Policy, Math Filters & Advisor AI (Months 24–36)
*   **Architecture Layers:** L3, L5.5, Reasoning
*   **Focus:** Advanced game-theoretic defense planning and advisory LLM integration.
*   **Deliverables:**
    *   **Stackelberg Policy Solver:** Optimizing game solver returns to $<1\text{ms}$ to update Arbiter weights dynamically.
    *   **Kalman & Wavelet Filters:** Noise stripping from raw signal entropy to prevent false positive triggers.
    *   **PhoenixMind LLM Integration:** RAG-based explanation interface. Keeps LLM strictly advisory, outputting policy updates for manual operator review or FSM limits gating.
*   **Acceptance Gate:** Stackelberg solver computes cost-payoff defenses in $<1\text{ms}$.

### Phase V: Autonomous Swarm OS (Months 36+)
*   **Architecture Layers:** L1, L2, L7
*   **Focus:** Custom scheduler and bare-metal kernel patches.
*   **Deliverables:**
    *   **Game-Aware Scheduler:** Custom patch to Linux Completely Fair Scheduler (CFS) allocating CPU slice based on Arbiter game-theoretic payoffs.
    *   **Entropy-Aware Allocator:** Linux page allocator extensions that randomize or isolate memory layouts based on process entropy scores.
    *   **Swarm Autonomous Quenching:** Fully closed-loop self-repairing swarm where nodes dynamically isolate compromised peers without operator intervention.
*   **Acceptance Gate:** Quorum consensus successfully isolates a rogue node within $<200\text{ms}$ of malicious detection.

---

## 3. Strict Implementation Gates

| Gate | Constraint | Status / Requirement |
| :--- | :--- | :--- |
| **Gate 1: Telemetry Saturation** | No graph processing or game theory implementation can begin until `09_telemetry/ebpf` captures >100k events/sec with <1% CPU overhead. | **PASSED** (Validated in replay) |
| **Gate 2: Graph Extraction Latency** | No physical modeling (SDI) or PID control can begin until process lineage DAGs (L4) can be extracted and queried in <1ms. | **PASSED** (Traversals verified) |
| **Gate 3: Solver Performance** | No PID loop or kernel scheduling can use Stackelberg outputs until the solver consistently returns policies in <1ms. | **PASSED** (Validated in replay harness) |
| **Gate 4: Kernel Space Lock** | **ABSOLUTELY NO KERNEL MODIFICATIONS** until userspace PID control is completely validated under adversarial conditions. | **PASSED** (Verified in integration) |
| **Gate 5: AI Evidence Layer** | AI modules are strictly advisory (RAG, assistant) and cannot take autonomous Warden actions. | **ACTIVE** (AI Orchestrator strictly bounds AI) |
