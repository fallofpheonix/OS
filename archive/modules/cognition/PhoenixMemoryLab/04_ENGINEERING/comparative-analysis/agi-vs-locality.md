# AGI vs Locality

Summary: map `fallofpheonix/agi` coordination architecture to your locality/containment model and surface coordination-gravity risks.

## Key Mappings
- Topology: P2P mesh (GossipSub + CRDT) → localized contributions that converge globally. This maps to "locality-first" only if gossip remains the dominant flow and pushes are asynchronous.
- Archive path: frequent pushes to a global GitHub archive create a synchronous global obligation that can amplify coordination gravity.
- Incentives: points/earning mechanics that reward scale → risk of resource concentration and positive feedback.

## Where Coordination Gravity Appears
- Shared provider pools and Pod VM capsule usage (shared credentials, pooled caches) centralize capability provisioning.
- Leader/pulse rounds create brief centralization windows — if any critical path depends on elected leaders, centralization mass increases.

## Locality Failure Modes
- Rapid mass adoption of a single configuration (gossip-driven but unbounded adoption) reduces per-node decision autonomy.
- Heavy reliance on GitHub archival as a coordination milestone turns archival into a global synchronization point.

## Containment Recommendations
- Keep archival and durable pushes asynchronous and non-blocking to avoid turning archive into a coordination primitive.
- Introduce explicit locality-preservation invariants: measure per-round dependency concentration (e.g., % of peers relying on top-10 nodes).
- Limit pooled-provider influence with budget/partitioning controls and per-agent quotas.

## Observability Hooks
- Use snapshots (`snapshots/latest.json`) to compute concentration metrics: Gini coefficient on contributions, top-k scrape frequency, leader influence per pulse.

## Short Actionable Items
- Instrument `snapshots/latest.json` with a small script to compute concentration metrics (top-10 share, Gini). Save results under `04_ENGINEERING/control-plane-analysis/metrics/`.
- Add a note to `agi-analysis.md` describing which incentives to monitor (uptime bonus, capability bonus, pooled-provider adoption rate).
