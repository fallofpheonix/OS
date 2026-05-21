# Dependency Graph (high-level)

Nodes (modules)
- `phoenix_os/monitor` — telemetry ingestion
- `phoenix_os/bus` — event bus
- `phoenix_os/trace` — lineage DAG & storage
- `phoenix_os/ledger` — evidence chain
- `07_security/control` — warden/actuation
- `phoenix_os/warden` — prototype controller
- `10_kernel` — eBPF hooks and kernel scheduler patches
- `07_security/nexus` — swarm
- `agents/surface/orchestrator` — control plane API

Edges (direction)
- monitor → bus → trace → ledger → control → kernel
- control → agents/surface/orchestrator (cloud control)
- nexus ↔ agents/surface/orchestrator (gossip/control plane)

Notes
- Keep dependencies acyclic where possible; prefer pub/sub and RPC contracts between layers.
- Add an explicit schema-and-contract artifact: `02_docs/specifications/telemetry_agent_spec.md` is authoritative.
