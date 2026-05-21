# PhoenixOS: Task List

## Phase 1: Core Communications & Monitoring
1.  **[Phoenix Bus]** System-wide message router (L2/L3 backbone). **[IN PROGRESS]**
2.  **[Phoenix Monitor]** Real-time entropy and signal analysis (L3).
3.  **[Phoenix Trace]** Process causality and lineage mapping (L4).

## Phase 2: Integrity & Decisions
4.  **[Phoenix Sentinel]** Thermodynamic system integrity (L6).
5.  **[Phoenix Arbiter]** Game-theoretic strategic policy engine (L5.5).

## Phase 3: Control & Actuation
6.  **[Phoenix Warden]** Closed-loop resource control and containment (L5).
7.  **[Phoenix Kernel]** In-kernel security hooks and scheduling (L1/L2).

---
## Task 1: Phoenix Bus (IPC Backbone)
- **Purpose:** Centralized event distribution.
- **Subsystem:** IPC / Message Bus.
- **Interface:** Pub/Sub via Go channels.
- **Budget:** < 50ns per fan-out.
