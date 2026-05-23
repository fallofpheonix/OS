# Component Map: The Phoenix Matrix

## 1. Layered Subsystems

| Layer | Component | Name | Description |
| :--- | :--- | :--- | :--- |
| **L7** | Swarm Coordination | **Phoenix Nexus** | Distributed consensus (PoA + Reputation). |
| **L6** | System Physics | **Phoenix Sentinel** | Thermodynamic SDI monitoring. |
| **L5.5** | Strategic Policy | **Phoenix Arbiter** | Game-theoretic policy (Stackelberg). |
| **L5** | Actuation & Control | **Phoenix Warden** | Finite-State Controller (SAFE -> COMPROMISED). |
| **L4** | Graph Intelligence | **Phoenix Trace** | Causal lineage DAGs with 3-tier storage. |
| **L3** | Telemetry Math | **Phoenix Monitor** | Signal processing (Entropy + Kalman). |
| **L2** | Kernel Runtime | **Phoenix Kernel** | eBPF probes and event hooks. |
| **L1** | Platform Integrity | **Phoenix Guard** | <100ms Fast-Path enforcement. |

## 2. Core Infrastructure
- **Phoenix Ledger:** SHA-256 hash-chained verifiable record of all system state changes.
- **Phoenix Mind:** The collective AI/ML models providing advisory intelligence to the Arbiter.

## 3. Ownership & Dependencies
- **Core Runtime:** PhoenixOS Kernel Team.
- **Intelligence Layers:** Phoenix AI Research Team.
- **External Dependencies:** Linux Kernel (5.x+), eBPF, LLVM/Clang (for probe compilation).
