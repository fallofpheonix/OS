# Scope Singularity Mitigation: The Core Executable Strategy

To prevent project stall and ensure architectural stability, PhoenixOS will prioritize the development of a **Telemetry-First Executable Core** before expanding into higher-order control and reasoning layers.

## The Core Executable Subset
The project will focus exclusively on the following subsystems until full stabilization:

1.  **Telemetry Runtime:** High-throughput, reliable event capture (eBPF).
2.  **Evidence Ledger:** Verifiable, immutable hash-chaining of system events.
3.  **Replay & Simulation Runtime:** Ability to reconstruct causal state from the ledger.
4.  **Graph Lineage Runtime:** Real-time construction of process/network dependency DAGs.

## Phased Expansion Roadmap
Only upon the stabilization of the above core will development shift to the next layer:

*   **Phase 1 (Stabilization):** Telemetry + Evidence + Replay + Graph Lineage.
*   **Phase 2 (Control):** Safe Actuation + FSM Constraints + Budgeting.
*   **Phase 3 (Orchestration):** Distributed Consensus + Cloud Scheduler.
*   **Phase 4 (Reasoning):** AI Reasoning + Planning + Simulation.
*   **Phase 5 (Physics/Games):** Physics-inspired modeling + Game-theory optimization.

This strict layering prevents scope creep and ensures that each subsystem is built on a proven, observable foundation.
