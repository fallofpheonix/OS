# AI / Cloud / Distributed Micro-Features Catalog

This document catalogs AI/ML, Cyber, Cloud, and Distributed micro-features mapped to the extended roadmap. These features support building the Cloud Distributed OS and Sentinel Cloud capabilities.

## AI / ML features (AI-ML-001..AI-ML-025)
- AI-ML-001: Anomaly detector
- AI-ML-002: Risk classifier
- AI-ML-003: Threat labeler
- AI-ML-004: Log embedding
- AI-ML-005: Graph embedding
- AI-ML-006: Sequence predictor
- AI-ML-007: Time series model
- AI-ML-008: False positive scorer
- AI-ML-009: Confidence estimator
- AI-ML-010: Policy recommender
- AI-ML-011: Node ranking
- AI-ML-012: Attack prediction
- AI-ML-013: Model registry
- AI-ML-014: Model version tracker
- AI-ML-015: Model rollback
- AI-ML-016: Online learning
- AI-ML-017: Federated update
- AI-ML-018: Drift detector
- AI-ML-019: Feature store
- AI-ML-020: Training queue
- AI-ML-021: Inference cache
- AI-ML-022: Prompt memory
- AI-ML-023: RAG index
- AI-ML-024: Replay explanation
- AI-ML-025: SOC copilot

## Cybersecurity features (CYB-001..CYB-025)
- CYB-001: Process trust score
- CYB-002: Binary reputation
- CYB-003: Command anomaly
- CYB-004: Privilege jump detector
- CYB-005: Credential tracker
- CYB-006: Token misuse detector
- CYB-007: Persistence scanner
- CYB-008: Registry watcher
- CYB-009: Kernel event monitor
- CYB-010: File entropy monitor
- CYB-011: Ransom burst detector
- CYB-012: Lateral movement graph
- CYB-013: Beacon detector
- CYB-014: C2 suspicion score
- CYB-015: DNS anomaly
- CYB-016: Threat timeline
- CYB-017: MITRE mapper
- CYB-018: Evidence signer
- CYB-019: Incident builder
- CYB-020: Attack replay
- CYB-021: Sandbox launcher
- CYB-022: Honey token
- CYB-023: Decoy node
- CYB-024: Auto isolate
- CYB-025: Recovery planner

## Cloud features (CLD-001..CLD-025)
- CLD-001: VM launcher
- CLD-002: Container launcher
- CLD-003: Image registry
- CLD-004: Snapshot manager
- CLD-005: Volume attach
- CLD-006: Volume replicate
- CLD-007: Cloud scheduler
- CLD-008: Workload placement
- CLD-009: Policy push
- CLD-010: Global config
- CLD-011: Secrets manager
- CLD-012: Identity provider
- CLD-013: API gateway
- CLD-014: Load balancer
- CLD-015: Service discovery
- CLD-016: Autoscaling
- CLD-017: Cold migration
- CLD-018: Hot migration
- CLD-019: Cloud replay
- CLD-020: Telemetry aggregation
- CLD-021: Cluster dashboard
- CLD-022: Remote evidence
- CLD-023: Cloud SOC
- CLD-024: Distributed storage
- CLD-025: Cloud shell

## Distributed system features (DST-001..DST-025)
- DST-001: Node heartbeat
- DST-002: Leader election
- DST-003: Consensus vote
- DST-004: Node trust
- DST-005: Node join
- DST-006: Node leave
- DST-007: Replica tracker
- DST-008: Event replication
- DST-009: Distributed replay
- DST-010: Global process ID
- DST-011: Cross host lineage
- DST-012: Remote execution
- DST-013: Task migration
- DST-014: Failure detector
- DST-015: Clock sync
- DST-016: State sync
- DST-017: Shard manager
- DST-018: Partition detector
- DST-019: Recovery node
- DST-020: Distributed cache
- DST-021: Distributed graph
- DST-022: Evidence federation
- DST-023: Cluster entropy
- DST-024: Topology mapper
- DST-025: Global namespace

## Mapping to phases & priorities
- Phase 5 (Distributed Telemetry Layer): DST-008, DST-009, DST-011, DST-021, DST-022, DST-023
- Phase 6 (Cloud Control Plane): CLD-007, CLD-008, CLD-009, CLD-020, CLD-022, CLD-024
- Phase 11 (AI Runtime): AI-ML-001..AI-ML-025 (start with anomaly, drift, model registry)
- Cross-cutting Cyber basics (P0): CYB-001..CYB-010

## Implementation guidance
- Start with minimal, well-tested primitives: `DST-001` (heartbeat), `DST-015` (clock sync), `DST-016` (state sync), `CLD-020` (telemetry aggregation), `AI-ML-001` (anomaly detector), and `CYB-010` (file entropy monitor).
- Gate distributed/Cloud rollout behind replay determinism and evidence federation tests.
- Treat AI models as replaceable components with a model registry and versioning; enable model rollback and federated update for safety.

## Notes
- Each micro-feature should be tracked as an individual issue with an owner, tests, and benchmark targets.
- These lists are intentionally exhaustive; prioritize and stage implementation using the recommended P0–P3 buckets in the roadmap.
