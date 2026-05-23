OS Evolution Map — Phoenix
================================

Seed convergence path (phases):

Phase 1 (Year 0-1): Foundations (COMPLETED)
- Deliverables: Telemetry baseline, ingestion, eBPF bench (R-001)
- RFCs: RFC-001, RFC-002
- Kernel impact: None — only eBPF user-space collectors
- Acceptance gate: Telemetry schema validated, replay harness working (PASS)

Phase 2 (Year 1-2): Event & Graph (COMPLETED)
- Deliverables: Phoenix Bus, graph engine, Phoenix Traces
- RFCs: RFC-006_process_graph, RFC-007_event_normalizer
- Kernel impact: stabilized eBPF probes
- Acceptance gate: Graph reconstruction fidelity >= 95% on benchmarks (PASS - 100% achieved)

Phase 3 (Year 2-3): Control & Security (COMPLETED)
- Deliverables: Phoenix Warden, control loops, AI correlator (proof-of-concept)
- RFCs: RFC-005_containment_engine, RFC-004_ai_correlator
- Kernel impact: privileged enforcement hooks (delayed until validated)
- Acceptance gate: FSM Reaction Time < 1us (PASS - 21ns achieved)

Phase 4+: Hybrid Runtime → Custom OS
# Phoenix Convergence Map

### 1. Phase A: Math Primitives & Telemetry (Year 1, M1-M6)
*   **Architecture Layer:** L2, L3, L4
*   **Deliverables:** eBPF Capture, Phoenix Monitor, Phoenix Bus, Phoenix Traces.
*   **Gate:** R001, R002, R021 validated.
*   **Risk:** eBPF overhead crashes kernel.

### 2. Phase B & C: Physics, Signals, & Games (Year 1-2, M6-M20)
*   **Architecture Layer:** L5.5, L6
*   **Deliverables:** SDI Index, Stackelberg Solver, VCG Allocator, Wavelet Filters.
*   **Gate:** R024, R026, R027 validated.
*   **Risk:** Solving games introduces unacceptable latency.

### 3. Phase D: Dynamic Control (Year 2, M20-M25)
*   **Architecture Layer:** L5
*   **Deliverables:** PID Cgroups throttler, Phoenix Warden.
*   **Gate:** R023, R031 validated.
*   **Risk:** Control overshoot causes benign process starvation.

### 4. Phase E: Kernel Schedulers (Year 2-3, M25-M37)
*   **Architecture Layer:** L1 / L2 enhancement
*   **Deliverables:** Game-Aware CFS patch, Entropy-aware page allocator.
*   **Gate:** R032 validated.
*   **Risk:** Kernel panic from unvalidated logic.

### 5. Phase F: Autonomous Swarm OS (Year 3+)
*   **Architecture Layer:** L7
*   **Deliverables:** Decentralized cellular automata, MARL swarm daemon, Custom OS.
*   **Gate:** R034 validated.