# AGI — Coordination Analysis

File: quick analysis of `fallofpheonix/agi` (Hyperspace fork). Focus: coordination, memory, tooling, and centralization pressure — treat this repo as a semantic coordination system, not an "AI feature" project.

## Coordination Model
- Topology: peer-to-peer mesh (libp2p/GossipSub) with voluntary capability advertising (Inference, Training, Memory, Orchestration, etc.).
- Two-tier convergence: fast GossipSub (seconds) for inspiration + CRDT leaderboards (minutes) for convergent state + durable GitHub archive (human-readable durable store).
- Agents are independent peers identified by libp2p peer IDs; snapshots and GitHub branches provide durable global view.

## Failure Semantics
- Fail-stop and transient failures tolerated via gossip and CRDT convergence; training deltas are compressed and retried.
- Distributed training relies on local progress + occasional aggregation; failure modes include stragglers, delta loss, snapshot inconsistency.
- Pulse verification (VRF + commit-reveal) creates a dependency on timely participation; partial participation reduces reward but must not break progress.

## Trace Semantics
- Traces are decentralized: agent journals, run-*.json, and CRDT leaderboards form the canonical traces of experiments.
- Hourly network snapshots produce a global trace artifact (snapshots/latest.json) suitable for offline analysis; these are coarse-grained but frequent enough to observe trends.

## Hidden State Risks
- Hidden local state: per-agent caches, model artifacts, and sidecars (pod VM) are opaque until pushed; this creates delayed observability.
- GitHub-as-archive risk: pushing only 'best' results to durable archive hides intermediate state and decisions.
- Patchy replication: nodes with stronger connectivity or resources implicitly gain influence unless measures exist to bound their weight.

## Centralization Pressure
- Incentives: points economy (uptime, capability bonuses, work points) can drive resource concentration (nodes buying more capacity to earn more), producing coordination gravity.
- Shared services (Pod VM, pooled providers) concentrate control-plane requirements and can create hot paths for failure and central control.
- Leader-election or pulse verification centralizes attention during rounds — watch for single-node leadership dependencies.

## Coordination Gravity Signals
- Rapid adoption of a single configuration across many agents (gossip → mass migration) indicates low locality preservation.
- Increased frequency of global snapshots, pushes to GitHub, or leaderboard-driven workflows correlates with rising coordination mass.
- Emergence of shared provider pools (pooled API keys, shared model caches) is a strong signal that local autonomy is being reduced.

## Runtime Isolation
- Pods and AVM (Agent Virtual Machine) provide isolation primitives; however, Pod capsules and sidecars expose surface area for shared-state leaks.
- Resource heterogeneity implies varied isolation guarantees; design for weak isolation and validate boundaries (I/O, filesystem, network).

## Architectural Weaknesses
- Incentive coupling to global metrics risks centralization (points → resource concentration → more points).
- GitHub archive as coordination primitive hides intermediate experiment semantics and can induce global coordination to satisfy archival requirements.
- Complexity of distributed training and the pulse protocol increases cognitive load; this complexity can mask coordination mass growth.

## Ideas Worth Preserving
- CRDT leaderboards for convergent, eventually-consistent shared metrics — they preserve local contribution while allowing global ranking.
- Gossip-first, archive-later flow (fast local gossip → convergent CRDT → durable push) supports locality-first operation.
- Compact delta exchange (SparseLoCo + Parcae) as an efficient, bandwidth-friendly collaborative training primitive.

## Ideas Worth Avoiding
- Reward rules that directly scale with resource size without bounding locality (unlimited earning incentives).
- Treating GitHub archive pushes as part of the critical path for coordination — archival should be asynchronous and non-blocking.
- Allowing pooled provider credentials to become the normative path for capability provisioning without strict budget/partitioning controls.

---

Next actions (recommended):
- Clone the repo locally and inspect `agents/` + `docs/` + `network-snapshots` for concrete trace structure.
- Run a lightweight snapshot analysis on `snapshots/latest.json` to quantify coordination signals (frequency of leaderboard churn, concentration metrics).
