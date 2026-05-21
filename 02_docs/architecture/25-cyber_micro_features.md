# Cybersecurity Micro-Features Catalog

This document catalogs passive defense, active defense, attack-analysis, attack-simulation, support/resilience, deception, and strategic micro-features for SentinelOS. These map to the Cyber Operations layer and the roadmap's phased implementation.

## Passive defense features (PDEF-001..PDEF-020)
- PDEF-001: Process watcher
- PDEF-002: File watcher
- PDEF-003: Registry monitor
- PDEF-004: DNS observer
- PDEF-005: TLS metadata capture
- PDEF-006: Network flow collector
- PDEF-007: Kernel event tap
- PDEF-008: Memory pressure tracker
- PDEF-009: Credential usage tracker
- PDEF-010: Session observer
- PDEF-011: API invocation logger
- PDEF-012: Privilege escalation watcher
- PDEF-013: Entropy scanner
- PDEF-014: Persistence monitor
- PDEF-015: USB observer
- PDEF-016: Cloud audit collector
- PDEF-017: Container monitor
- PDEF-018: Namespace observer
- PDEF-019: Process lineage builder
- PDEF-020: Trust scorer

Combines into: Passive Security Runtime

## Active defense features (ADEF-001..ADEF-020)
- ADEF-001: Process throttle
- ADEF-002: Priority reduction
- ADEF-003: Network isolate
- ADEF-004: Remote freeze
- ADEF-005: Node quarantine
- ADEF-006: Auto snapshot
- ADEF-007: Policy injection
- ADEF-008: Dynamic firewall
- ADEF-009: Credential revoke
- ADEF-010: Session kill
- ADEF-011: Container pause
- ADEF-012: Decoy deploy
- ADEF-013: Honey credential
- ADEF-014: Traffic redirect
- ADEF-015: Replay trigger
- ADEF-016: Recovery switch
- ADEF-017: Safe mode entry
- ADEF-018: Cluster isolation
- ADEF-019: Rollback action
- ADEF-020: Defense escalation

Combines into: Adaptive Defense Engine

## Passive attack-analysis features (PATT-001..PATT-020)
- PATT-001: Recon pattern detector
- PATT-002: Port scan tracker
- PATT-003: DNS anomaly detector
- PATT-004: Beacon interval detector
- PATT-005: Lateral path graph
- PATT-006: Credential abuse score
- PATT-007: Privilege chain tracker
- PATT-008: File mutation graph
- PATT-009: Command rarity
- PATT-010: Payload classifier
- PATT-011: Timeline generator
- PATT-012: MITRE mapper
- PATT-013: IOC matcher
- PATT-014: Campaign correlator
- PATT-015: Risk propagation
- PATT-016: Exfiltration estimator
- PATT-017: Evidence builder
- PATT-018: Threat confidence
- PATT-019: Host impact score
- PATT-020: Replay reconstruction

Combines into: Threat Intelligence Runtime

## Active attack simulation features (ASIM-001..ASIM-020)
- ASIM-001: Recon simulator
- ASIM-002: Phishing scenario
- ASIM-003: Credential attack model
- ASIM-004: Privilege escalation simulation
- ASIM-005: Lateral movement simulation
- ASIM-006: Ransom scenario
- ASIM-007: Persistence simulation
- ASIM-008: Beacon emulation
- ASIM-009: Exfiltration scenario
- ASIM-010: Insider threat model
- ASIM-011: APT workflow simulation
- ASIM-012: SOC stress mode
- ASIM-013: Recovery drill
- ASIM-014: Incident replay
- ASIM-015: Blue team challenge
- ASIM-016: Red team arena
- ASIM-017: Purple team mode
- ASIM-018: Digital twin attack
- ASIM-019: Cloud breach scenario
- ASIM-020: Distributed incident replay

Combines into: Cyber Range System

## Support / resilience features (SUP-001..SUP-020)
- SUP-001: Evidence archive
- SUP-002: Replay database
- SUP-003: Threat memory
- SUP-004: Knowledge graph
- SUP-005: Case similarity search
- SUP-006: Recovery planner
- SUP-007: Incident notebook
- SUP-008: SOC dashboard
- SUP-009: Risk history
- SUP-010: Policy library
- SUP-011: Rule versioning
- SUP-012: Model registry
- SUP-013: Experiment storage
- SUP-014: Forensic export
- SUP-015: Trust ledger
- SUP-016: Audit chain
- SUP-017: Cluster history
- SUP-018: Validation store
- SUP-019: Decision log
- SUP-020: Playbook engine

Combines into: Security Support Layer

## Deception features (DEC-001..DEC-015)
- DEC-001: Honey files
- DEC-002: Honey users
- DEC-003: Honey credentials
- DEC-004: Honey shares
- DEC-005: Honey process
- DEC-006: Fake services
- DEC-007: Virtual host
- DEC-008: Decoy database
- DEC-009: Shadow node
- DEC-010: Mirror network
- DEC-011: Synthetic traffic
- DEC-012: Fake telemetry
- DEC-013: False asset map
- DEC-014: Trap session
- DEC-015: Adaptive deception

Combines into: Deception Runtime

## Strategic layer
Attack analysis + Passive defense + Active defense + Simulation + Support + Deception = Cyber Operations Layer

## Implementation order (recommended)
- P0: Passive defense, Support, Evidence
- P1: Threat analysis, Replay, Deception
- P2: Active defense
- P3: Simulation range
- P4: Autonomous cyber operations

## Notes
- Track each micro-feature as an independent issue with owner, tests, and benchmarks.
- Prioritize P0 features to establish secure telemetry and evidence first.
- Avoid shipping in-kernel enforcement until replay determinism and evidence verification are in place.
