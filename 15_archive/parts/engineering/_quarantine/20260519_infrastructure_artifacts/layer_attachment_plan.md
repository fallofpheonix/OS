# Layer Attachment Plan

## Overview
This plan outlines how future scientific and expansion layers will be attached to the ecosystem's remote runtime.

## Attachment Schedule

| Layer | Repo | GitHub | Dependencies | Attach Point | Runtime |
|---|---|---|---|---|---|
| **Physics (P1)** | physics | .../physics.git | [astraeus-core] | `workspace/active/physics` | generic |
| **Mathematics (P2)** | mathematics | .../mathematics.git | [physics] | `workspace/active/math` | generic |
| **Simulation (P3)** | simulation | .../simulation.git | [physics, math] | `workspace/active/simulation` | generic |
| **RL (P4)** | reinforcement-learning | .../RL.git | [simulation] | `workspace/active/RL` | generic |
| **Memory (P5)** | memory-engine | .../memory.git | [brain] | `workspace/active/memory` | generic |
| **Neuromorphic (P6)** | neuromorphic-computing | .../neuro.git | [brain, physics] | `workspace/active/neuromorphic` | generic |
| **Swarm (P7)** | agent-swarm | .../swarm.git | [forge-agent, control-plane] | `workspace/active/swarm` | generic |

## Attachment Procedure
1. **Registry Update:** Add layer to `remote_runtime/execution_registry.yaml` with `ON_DEMAND` trigger.
2. **Policy Definition:** Define `CLONE_ON_DEMAND` or `CACHE_TEMP` in `runtime_policy.yaml`.
3. **Dependency Check:** Verify all listed dependencies are either `present` or `clonable`.
4. **Attach Point Creation:** Ensure target directory in `workspace/active/` is ready.
5. **Execution Test:** Perform a dry-run clone into the attach point (simulated for now).

## Constraints
- No implementation until Phase G9 is complete.
- Layers must follow the `cache_lifecycle.md` for non-core development.
