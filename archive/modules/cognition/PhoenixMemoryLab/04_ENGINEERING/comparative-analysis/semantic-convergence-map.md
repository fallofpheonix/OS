# Semantic Convergence Map

Track where multiple repos converge on the same semantic pressures and topology problems.

Columns:

- **Repo**: repository name
- **Pressure Type**: e.g., scheduling, control-plane, trace aggregation
- **Semantic Pattern**: concise description of the recurring pattern
- **Locality Behavior**: does the pattern preserve locality or require global coordination?
- **Coordination Gravity Risk**: Low / Medium / High — how likely this pattern produces centralization
- **Trace Topology Shape**: flat / hierarchical / DAG / mesh
- **Extraction Potential**: None / Pattern / Component / Strategic

Example row:

| Repo | Pressure Type | Semantic Pattern | Locality Behavior | Coordination Gravity Risk | Trace Topology Shape | Extraction Potential |
|------|---------------|------------------|-------------------|--------------------------|----------------------|---------------------|
| `agi` | Orchestration | Gossip-led config adoption → mass migration | Weak locality (adoption favors global) | High | Mesh + CRDT leaderboard | Pattern / Strategic

Use this file to record convergences observed across `agi`, `react-three-fiber`, `liquid-*`, `shadergradient`, and `the_well`.

## Initial Observations (qualitative)

| Repo | Pressure Type | Semantic Pattern | Locality Behavior | Coordination Gravity Risk | Trace Topology Shape | Extraction Potential |
|------|---------------|------------------|-------------------|--------------------------|----------------------|---------------------|
| `agi` | Orchestration / Incentives | Gossip-driven rapid config adoption; reward-driven resource concentration | Erodes locality when incentives favor scale | High | Mesh + CRDT leaderboards | Pattern / Strategic
| `react-three-fiber` | Scheduling / Render-loop | Many per-frame subscribers (`useFrame`) create backpressure; selective subscription preserves locality | Preserved when subscriptions are conditional | Medium | Hierarchical scene graph with per-frame event fans | Pattern / Reimplementation
| `liquid-glass-js` | Rendering composition | Layered effects pipeline with adapter layer (DOM↔GL) | Local if effect chains are per-component; global if shared caches used | Low-Medium | Local DAG per-element | Component (visual primitives)
| `shadergradient` | GPU control-plane | Offload interpolation and timing to GPU; uniform conventions centralize parameterization | Local if materials owned by component; centralization if global uniforms used | Low | Per-material hierarchical | Pattern / Snippet
| `the_well` | Data / Benchmarking | Large-scale dataset streaming + hydra experiment control-plane | Not directly control-plane-heavy; useful as resource only | Low | Flat dataset catalog + benchmark runs | Resource / Knowledge Only

Use this map to mark repeated patterns across repos — when the same row appears for multiple repos, that signals semantic convergence.
