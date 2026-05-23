# System Architecture: PhoenixOS

PhoenixOS is structured around a 7-layer cybernetic stack known as the **Phoenix Matrix**.

```
                   ┌──────────────────────────────────────┐
                   │ L7: Swarm Coordination (Nexus)       │
                   └──────────────────┬───────────────────┘
                                      ▼
                   ┌──────────────────────────────────────┐
                   │ L6: System Physics (Sentinel)        │
                   └──────────────────┬───────────────────┘
                                      ▼
                   ┌──────────────────────────────────────┐
                   │ L5.5: Strategic Policy (Arbiter)     │
                   └──────────────────┬───────────────────┘
                                      ▼
                   ┌──────────────────────────────────────┐
                   │ L5: Actuation & Control (Warden)     │
                   └──────────────────┬───────────────────┘
                                      ▼
                   ┌──────────────────────────────────────┐
                   │ L4: Graph Intelligence (Trace)       │
                   └──────────────────┬───────────────────┘
                                      ▼
                   ┌──────────────────────────────────────┐
                   │ L3: Telemetry Math (Monitor/TCS)     │
                   └──────────────────┬───────────────────┘
                                      ▼
                   ┌──────────────────────────────────────┐
                   │ L2: Kernel Runtime (Probes)          │
                   └──────────────────┬───────────────────┘
                                      ▼
                   ┌──────────────────────────────────────┐
                   │ L1: Platform Integrity (Guard)       │
                   └──────────────────────────────────────┘
```

## Layer Breakdown

1. **L7: Swarm Coordination (Phoenix Nexus):** Proof-of-Anomaly consensus and multi-node replication.
2. **L6: System Physics (Phoenix Sentinel):** Continuous SDI signal physics modeling.
3. **L5.5: Strategic Policy (Phoenix Arbiter):** Game-theoretic decision engine solving Stackelberg policy bounds.
4. **L5: Actuation & Control (Phoenix Warden):** 5-State Discrete FSM (SAFE -> WATCH -> SUSPICIOUS -> CRITICAL -> COMPROMISED) enforcing hysteresis constraints.
5. **L4: Graph Intelligence (Phoenix Trace):** Causal process lineage DAGs stored across HOT/WARM/COLD tiers.
6. **L3: Telemetry Math (Phoenix Monitor):** Signal processing engine utilizing Kalman filters and sliding window sequence verification.
7. **L2: Kernel Runtime (Phoenix Kernel):** eBPF telemetry collectors hook process, socket, and file boundaries.
8. **L1: Platform Integrity (Phoenix Guard):** Microsecond-latency fast-path syscall enforcement maps.

## Inter-Process Communication (IPC) & Flow Paths
- **Raw Telemetry Ingestion:** Syscalls -> L1 (Guard) -> L2 (eBPF Probes) -> L3 (Event Bus).
- **Processing Loop:** L3 Event Bus -> L4 (Trace Lineage Graph) -> L3 Monitor (Entropy & Kalman calculations) -> L5.5 Arbiter (Payoff & Threshold gating) -> L5 Warden (FSM State changes).
- **Advisory AI Loop:** Chaotic signals (Z > 3.0) queue requests to PhoenixMind LLM (running as an asynchronous advisory thread) which provides offline audit suggestions.
