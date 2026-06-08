# Coordination Gravity Signals

Catalog observable signals that indicate rising coordination gravity across systems.

Signals (examples):

- Shared recovery semantics: many services rely on centralized retry/rollback coordination.
- Orchestration-global propagation: small local events require global control-plane to resolve.
- Centralized runtime awareness: single service or agent holds global knowledge needed for many actions.
- Topology-aware scheduling: scheduling decisions depend on global topology information.
- Shared trace mutation: multiple actors mutate the same global trace/leaderboard artifact.

For each signal, note: source repo(s), example locations (file/path), why it matters, and containment recommendations.

Example entry:

- **Signal**: Shared provider pool adoption
- **Repos**: `agi`
- **Files/Areas**: Pod VM, pooled providers documentation
- **Why it matters**: pooled credentials concentrate capability and create single points of failure and influence
- **Containment**: enforce per-agent quotas, budgeted provider partitions, async provisioning

---

- **Signal**: Pulse/leader round centralization
- **Repos**: `agi`
- **Files/Areas**: Pulse verification description, leader election sections
- **Why it matters**: Leaders or pulse-verified windows concentrate responsibility and can become single points of influence for decision-making or rewards
- **Containment**: limit leader authority to ephemeral coordination roles; ensure fallback/no-leader modes and distribute reward influence

- **Signal**: Per-frame subscription fan-out
- **Repos**: `react-three-fiber`
- **Files/Areas**: `useFrame` examples, large demos
- **Why it matters**: many components subscribing per-frame create systemic scheduling pressure and synchronized update cycles
- **Containment**: selective subscriptions, region-scoped updates, batched state changes

- **Signal**: Shared artifact archival as coordination milestone
- **Repos**: `agi`
- **Files/Areas**: GitHub archival, snapshots push
- **Why it matters**: making archival an implicit coordination milestone converts an otherwise asynchronous durability step into a synchronization point
- **Containment**: make archival asynchronous and optional for runtime decisions; treat snapshots as read-only observability

- **Signal**: Global resource pools (model caches, provider pools)
- **Repos**: `agi`, `react-three-fiber` (analogue: shared geometry/material cache)
- **Files/Areas**: Pod provider docs, shared caches
- **Why it matters**: pools improve efficiency but centralize access and failure domains
- **Containment**: expose local caches with best-effort replication; enforce quotas; prefer partitioned pools

*** End Patch
