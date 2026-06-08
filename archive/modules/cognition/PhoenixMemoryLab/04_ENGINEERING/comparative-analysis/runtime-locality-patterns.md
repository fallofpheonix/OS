# Runtime Locality Patterns

Compare runtime locality behaviors across repos and identify patterns that preserve or break locality.

Sections:

- **Pattern Name**: short name
- **Observed In**: repos/paths
- **Description**: what happens and why
- **Locality Outcome**: preserved / eroded
- **Containment Techniques**: how to preserve locality
- **Notes**: counterexamples or caveats

Example pattern:

- **Pattern Name**: Selective Frame Subscription
- **Observed In**: `react-three-fiber` (`useFrame`), some examples in forks
- **Description**: components subscribe conditionally to per-frame updates rather than always-on subscriptions
- **Locality Outcome**: preserved (reduces global update fan-out)
- **Containment Techniques**: region-limited subscriptions, event-triggered updates, batched updates
- **Notes**: translate to agent runtimes as conditional task polling or event-driven activities

---

- **Pattern Name**: Local CRDT / Leaderboard Convergence
- **Observed In**: `agi` (CRDT leaderboards, gossip flow)
- **Description**: local experiments converge via CRDTs into a global leaderboard; nodes adopt top results via gossip
- **Locality Outcome**: erosion when adoption is unconditional (agents replace local variants rapidly)
- **Containment Techniques**: allow multi-version coexistence, slow adoption windows, conditional adoption rules
- **Notes**: maintain local experiment diversity to avoid premature global convergence

- **Pattern Name**: Adapter Layer Isolation
- **Observed In**: `liquid-glass-js`, `shadergradient`
- **Description**: adapter boundary (DOM↔GL or CPU↔GPU) can contain heavy operations locally if designed as per-component adapters
- **Locality Outcome**: preserved when adapters own their resources; eroded when adapters share global uniforms/caches
- **Containment Techniques**: small per-component adapters, explicit ownership, avoid global uniform state
- **Notes**: document uniform naming and ownership conventions

- **Pattern Name**: Data Resource Separation
- **Observed In**: `the_well`
- **Description**: large datasets streamed independently of runtime reduce coordination needs by isolating data transfer from control-plane
- **Locality Outcome**: preserved (data is remote resource, not a coordination primitive)
- **Containment Techniques**: treat datasets as read-only resources; avoid embedding dataset state into runtime coordination logic
- **Notes**: useful pattern for separating heavy IO from control-plane

*** End Patch