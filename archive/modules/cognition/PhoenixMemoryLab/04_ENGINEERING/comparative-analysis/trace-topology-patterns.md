# Trace Topology Patterns

Document common trace topology shapes, mutation points, and aggregation pressures.

Sections:

- **Topology Shape**: flat / hierarchical / DAG / mesh
- **Mutation Points**: where traces are appended, merged, or compacted
- **Aggregation Pressure**: how often traces get consolidated into global artifacts
- **Risk**: does this shape encourage centralization of observability?
- **Containment**: recommendations to limit aggregation-induced coordination

Example entry:

- **Topology Shape**: Mesh + CRDT leaderboards
- **Mutation Points**: agent run-*.json pushes, CRDT merges, hourly snapshot export
- **Aggregation Pressure**: high when snapshots are used for decisions or leaderboards are authoritative
- **Risk**: medium-high — leaderboard-driven incentives encourage convergence
- **Containment**: use read-only aggregation for human review; do not make leaderboard state a critical runtime signal

---

- **Topology Shape**: Hierarchical scene graph + per-frame event fan
- **Observed In**: `react-three-fiber`
- **Mutation Points**: per-frame updates (`useFrame`), scene mount/unmount, shared material updates
- **Aggregation Pressure**: medium — many updates funnel into the render loop for per-frame composition
- **Risk**: medium — render-loop stalls act as a synchronous bottleneck
- **Containment**: region-scoped updates, off-main-thread composition, selective subscriptions

- **Topology Shape**: Local DAG of effect composition
- **Observed In**: `liquid-glass-js`, `shadergradient`
- **Mutation Points**: effect parameter updates, uniform changes, compositing passes
- **Aggregation Pressure**: low-medium depending on whether effects share global resources
- **Risk**: low if effects remain per-component; increases if shared caches/uniforms are used
- **Containment**: per-effect ownership, explicit composition APIs, avoid global uniform mutation

- **Topology Shape**: Flat dataset + episodic benchmark traces
- **Observed In**: `the_well`
- **Mutation Points**: benchmark run outputs, dataset checkpointing
- **Aggregation Pressure**: low for runtime control-plane; high for archival/benchmark aggregation
- **Risk**: low for runtime centralization; treat aggregated benchmark as research artifact, not control plane
- **Containment**: separate archival/benchmarking flows from runtime control loops

*** End Patch