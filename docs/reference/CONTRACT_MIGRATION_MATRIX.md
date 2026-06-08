---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Contract Migration Matrix

This matrix maps active interfaces to target contract packages and registers their migration status.

| Interface | Current Package | Future Contract | Migration Status |
| :--- | :--- | :--- | :--- |
| `PoAEngineInterface` | `core/Phoenix.Nucleus/PhoenixCore/arbiter` | `contracts/arbiter/v1` | pending |
| `Normalizer` | `core/Phoenix.Nucleus/PhoenixCore/bus/normalizer` | `contracts/bus/v1` | pending |
| `SchemaValidator` | `core/Phoenix.Nucleus/PhoenixCore/common/serialization` | `contracts/serialization/v1` | pending |
| `AuditProvider` | `core/Phoenix.Nucleus/PhoenixCore/containment/rollback` | `contracts/containment/v1` | pending |
| `ILedger` | `core/Phoenix.Nucleus/PhoenixCore/contracts` | `contracts/ledger/v1` | partial |
| `EventEnvelope` | `core/Phoenix.Nucleus/PhoenixCore/contracts/events/v1` | `contracts/events/v1` | instantiated |
| `Event` | `core/Phoenix.Nucleus/PhoenixCore/contracts/events/v1` | `contracts/events/v1` | instantiated |
| `Serializer` | `core/Phoenix.Nucleus/PhoenixCore/contracts/events/v1` | `contracts/events/v1` | instantiated |
| `Reconstructor` | `core/Phoenix.Nucleus/PhoenixCore/contracts/replay/v1` | `contracts/replay/v1` | instantiated |
| `ReplayEngine` (contract) | `core/Phoenix.Nucleus/PhoenixCore/contracts/replay/v1` | `contracts/replay/v1` | instantiated |
| `Snapshot` | `core/Phoenix.Nucleus/PhoenixCore/contracts/replay/v1` | `contracts/replay/v1` | instantiated |
| `Actuator` (contract) | `core/Phoenix.Nucleus/PhoenixCore/contracts/security/v1` | `contracts/security/v1` | instantiated |
| `Containment` | `core/Phoenix.Nucleus/PhoenixCore/contracts/security/v1` | `contracts/security/v1` | instantiated |
| `Escalation` | `core/Phoenix.Nucleus/PhoenixCore/contracts/security/v1` | `contracts/security/v1` | instantiated |
| `ResourceAllocator` | `core/Phoenix.Nucleus/PhoenixCore/ledger/src` | `contracts/ledger/v1` | pending |
| `AlertManager` | `core/Phoenix.Nucleus/PhoenixCore/monitoring` | `contracts/observability/v1` | pending |
| `Alerter` | `core/Phoenix.Nucleus/PhoenixCore/monitoring` | `contracts/observability/v1` | pending |
| `PeerDiscovery` | `core/Phoenix.Nucleus/PhoenixDistributed/discovery` | `contracts/distributed/v1` | pending |
| `ConsensusLedger` | `core/Phoenix.Nucleus/PhoenixDistributed/ledger` | `contracts/distributed/v1` | pending |
| `Invariant` | `core/Phoenix.Nucleus/PhoenixGuard` | `contracts/security/v1` | pending |
| `IncidentResponseManager` | `core/Phoenix.Nucleus/PhoenixGuard/security` | `contracts/security/v1` | pending |
| `ThreatDetector` | `core/Phoenix.Nucleus/PhoenixGuard/security` | `contracts/security/v1` | pending |
| `Actuator` (legacy) | `core/Phoenix.Nucleus/PhoenixGuard` | `contracts/security/v1` | pending |
| `CertificateValidator` | `core/Phoenix.Nucleus/PhoenixGuard` | `contracts/security/v1` | pending |
| `GraphProvider` | `core/Phoenix.Nucleus/PhoenixGuard` | `contracts/security/v1` | pending |
| `EventPublisher` | `core/Phoenix.Nucleus/PhoenixKernel` | `contracts/kernel/v1` | pending |
| `PromptGuard` | `core/Phoenix.Nucleus/PhoenixMind/security` | `contracts/security/v1` | pending |
