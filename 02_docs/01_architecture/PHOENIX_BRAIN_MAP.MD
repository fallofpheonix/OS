# Phoenix Brain Architectural Map

This document serves as the single source of truth for the Phoenix Brain components and their locations within the PhoenixOS codebase.

| Brain Component | Role | Codebase Path / Repository |
| :--- | :--- | :--- |
| **Phoenix Arbiter** | Strategic (L5.5) | `external/integrated/phoenix-control/arbiter` |
| **Phoenix Trace** | Causal (L4) | `phoenix_os/telemetry/process_lineage` |
| **Cognition Engine** | Advisory (F7) | `phoenix_os/ai/phoenix-mind.Modelfile` |
| **Phoenix Sentinel** | Sensory (L6) | `phoenix_os/security/physics/disorder` |
| **Phoenix Nexus** | Swarm (L7) | `phoenix_os/nexus_coordination` |

## Dependencies
- **Arbiter** must consume **Trace** lineage graphs and **Sentinel** SDI scores to make cost-aware decisions.
- **Warden** consumes **Arbiter** policy output.
- **Nexus** provides distributed context to the **Arbiter** in swarm deployments.
- **Cognition Engine (PhoenixMind)** provides human-readable advisory insights based on the state of all above components.
