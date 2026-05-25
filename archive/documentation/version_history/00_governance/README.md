# PhoenixOS Documentation

This directory contains the authoritative documentation for the PhoenixOS project, organized into a 12-layer architecture to ensure complete system traceability.

## Documentation Structure

### [L0: Governance](./00_governance/)
Project Vision, Roadmap, Master Status, and Decision Logs (ADRs).

### [L1: Architecture](./01_architecture/)
System Architecture, Component Maps, RFCs, and Technical Specifications.

### [L2: Integration](./02_integration/)
External Repository Tracking, Merge Policies, and Upstream Synchronization.

### [L3: Agents](./03_agents/)
AI Agent Registry, Interaction Maps, and Behavior Protocols.

### [L4: Security](./04_security/)
Threat Models, Attack Surfaces, Security Boundaries, and Incident Response.

### [L5: Kernel](./05_kernel/)
Boot Sequence, Process/Memory Models, Scheduler, and IPC Design.

### [L6: Research](./06_research/)
Experimental Logs, Math Models, Game Theory, and Optimization Logs.

### [L7: Validation](./07_validation/)
Test Plans, Matrices, Benchmarks, and Verification Reports.

### [L8: Deployment](./08_deployment/)
Build Pipelines, CI/CD, Release Processes, and Installation Guides.

### [L9: Operations](./09_operations/)
Telemetry Schemas, Metrics, Observability, and Event Stream Definitions.

### [L10: Runtime AI](./10_runtime_ai/)
Knowledge Graphs, Model Lifecycle, Learning Policies, and Adaptation Rules.

### [L11: Emergency](./11_emergency/)
Disaster Recovery, Safe Modes, and Manual Override Protocols.

---

## Shared Resources
- **[Decisions](./decisions/):** Repository of Architectural Decision Records (ADRs).
- **[Reports](./reports/):** System health and validation reports.
- **[Diagrams](./diagrams/):** Visual architectural maps.
- **[Experiments](./experiments/):** Raw experimental data and logs.
- **[Archives](./archives/):** Deprecated or historical documentation.

## Documentation Policy
1. **Traceability:** No code change without a corresponding document update.
2. **Provenance:** Maintain clear dependency and external repo tracking.
3. **Auditability:** Every subsystem must maintain state and assumptions.
4. **Immutability:** Immutable audit trails for all agent-driven modifications.
