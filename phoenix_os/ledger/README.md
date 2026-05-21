# Phoenix Ledger (Evidence Layer)

The verifiable source of truth for all autonomous actions in PhoenixOS.

## Purpose
Binds causal process lineage (Trace) to physical system state (Sentinel) and strategic decisions (Arbiter).

## Data Structure (Evidence Tuple)
```text
(
  trace_hash,   // SHA256 of the process DAG subgraph
  sdi,          // Security Disorder Index at time of action
  policy,       // ID of the Arbiter policy used
  action,       // The Warden actuation taken (STOP/THROTTLE/KILL)
  result,       // Success/Failure of the action
  time,         // Nanosecond timestamp
  confidence,   // AI/Model certainty score (0-1.0)
  replay,       // Replay ID for deterministic reproduction
  experiment    // R-ID if running in experimental mode
)
```

## Validation Gates
- [ ] Ledger entry integrity (Hash chaining).
- [ ] Retrieval latency < 1ms.
- [ ] Cryptographic link to Phoenix Trace.
