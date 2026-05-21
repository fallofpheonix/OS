# Knowledge Graph: PhoenixOS

This graph represents the core architectural concepts and subsystems of PhoenixOS.

## Core Pillars
*   **Mathematical-Physical-Game Architecture:** Treating security as a thermodynamic state.
*   **The Phoenix Matrix (7-Layer Stack):**
    *   **L7: Swarm Coordination (Phoenix Nexus):** PoA + Reputation consensus.
    *   **L6: System Physics (Phoenix Sentinel):** Thermodynamic SDI monitoring.
    *   **L5.5: Strategic Policy (Phoenix Arbiter):** Stackelberg Security Games.
    *   **L5: Actuation & Control (Phoenix Warden):** PID/FSM feedback loops.
    *   **L4: Graph Intelligence (Phoenix Trace):** Causal process lineage DAGs.
    *   **L3: Telemetry Math (Phoenix Monitor):** Signal processing (Entropy + Kalman).
    *   **L2: Kernel Runtime (Phoenix Kernel):** eBPF probes and ring buffers.
    *   **L1: Platform Integrity (Phoenix Guard):** <100ms Fast-Path enforcement.

## Telemetry and Evidence Systems
*   **Phoenix Ledger:** Content-addressable evidence chain (SHA-256).
*   **Security Disorder Index (SDI):** Thermodynamic measure of system threats.
*   **Shannon Entropy / Kalman Filters:** Noise reduction in telemetry.

## System Relationships
*   `Monitor (L3)` provides telemetry to `Sentinel (L6)` and `Trace (L4)`.
*   `Arbiter (L5.5)` utilizes `Sentinel (L6)` data for policy decisions.
*   `Warden (L5)` executes decisions from `Arbiter (L5.5)`.
*   `Guard (L1)` provides low-latency blocking for critical threats detected by `Monitor (L3)` or others.
