# Risk Register (summary)

P0 Critical
- Evidence chain integrity compromise (P0): If ledger can be spoofed, all actions are unverifiable.
  - Mitigation: Audit cryptographic signing; require remote attestation; add evidence replay tests.

- Kernel Fast-Path errors causing service disruption (P0): In-kernel blocking may accidentally block benign workloads.
  - Mitigation: Feature-flag kernel changes, staged rollout, canary nodes, kill-switch and whitelist.

P1 Major
- Scheduler weight misconfiguration causing DoS (P1): Threat-weighted scheduling could starve critical processes.
  - Mitigation: Safe defaults, max-throttle limits, circuit-breaker, observability.

- Distributed control inconsistency (P1): Divergent controller state across nodes may cause instability.
  - Mitigation: Consensus protocols, leader leases, versioned policies.

P2 Moderate
- Performance regressions from additional telemetry (P2).
  - Mitigation: Sampling, adaptive telemetry rates, overhead budgets.

P3 Future
- Regulatory and privacy compliance for telemetry (P3).
  - Mitigation: Data minimization, on-device aggregation, PI redaction.
