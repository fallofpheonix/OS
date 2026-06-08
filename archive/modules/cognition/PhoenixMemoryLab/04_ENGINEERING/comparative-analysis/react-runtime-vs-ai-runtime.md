# React Runtime vs AI Runtime

Compare `react-three-fiber` runtime control-plane semantics against agent/AI runtime control-plane patterns.

## Common Themes
- Declarative surface + imperative engine: both R3F and agent systems expose a high-level intent API that maps to imperative runtime objects (Three.js objects vs agent processes/tools).
- Scheduling: `useFrame` subscriptions ≈ agent task loops; both can create backpressure when many subscribers/agents want CPU/GPU or network I/O.
- Resource pooling: three.js materials/geometries pooling ≈ model or cache pooling in agent networks.

## Coordination Gravity Parallels
- Global state stores in R3F mirror global memory or provider pools in agent runtimes; both reduce locality and increase update fan-out.
- Frequent synchronous updates to shared stores induce central scheduling pressure (render loop stalls ↔ control-plane bottlenecks).

## Locality-Preserving Patterns to Transfer
- Selective subscription: avoid global per-frame subscriptions; use conditional or region-limited subscriptions.
- Ownership clarity: define single owner per mutable object (mesh/material) — analogous to single authority for a memory shard or model cache.
- Backpressure-aware batching: group small updates into larger, scheduled batches to reduce global coordination events.

## Audit Checklist (for cross-repo comparison)
- Count active subscriptions (useFrame or task loops) and measure update rate.
- Identify global stores used for per-object/state data and evaluate their read/write fan-out.
- Find pooling patterns and evaluate whether they centralize resource access.
