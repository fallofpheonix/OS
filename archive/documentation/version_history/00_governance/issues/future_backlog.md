# ISSUE BACKLOG: Future Roadmap (P5-P7)

## P5: Distributed Control Plane (Nexus)

### Issue #50: [Distributed] Multi-node Replay Synchronization
- Implement Lamport/Vector clock synchronization across nodes to ensure event ordering in cluster-wide replays.
- **Risk:** High

### Issue #51: [Distributed] Cluster Actuation Coordination
- Build consensus mechanism to ensure isolation actions are consistent across multiple nodes.
- **Risk:** High

### Issue #52: [Distributed] Event Sourcing Consensus
- Move from single-node evidence ledger to distributed event sourcing with raft-based consensus.
- **Risk:** High

## P6: Advisory AI Layer (Arbiter Intelligence)

### Issue #60: [AI] Natural Language Incident Summarization
- Build an AI-assisted explanation layer that summarizes complex causal DAGs into plain English for operators.
- **Risk:** Low (Non-authoritative)

### Issue #61: [AI] Multi-stage Verification for AI Policy
- Implement a 'Deterministic Validator' layer that checks AI-suggested policies against FSM safety constraints before submission to human operators.
- **Risk:** Medium

## P7: Simulation & Game Theory Engine

### Issue #70: [Research] Game-Theory Containment Planning
- Model attacker vs. defender game scenarios to optimize node-isolation thresholds.
- **Risk:** Low (Research-focused)

### Issue #71: [Simulation] Cyber-Range Replay Harness
- Build the infrastructure to replay ledger-captured attacks in a sandbox environment to test Warden's efficacy.
- **Risk:** Medium
