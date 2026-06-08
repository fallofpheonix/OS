---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Interface Inventory

This document maps all active, uncommented interfaces in the core system context (`core/Phoenix.Nucleus`) to their owners, consumers, target contracts, and migration status. Commented-out skeleton/placeholder interfaces are excluded.

| Interface | Package | Owner | Consumers | Target Contract | Migration Status |
| :--- | :--- | :--- | :--- | :--- | :--- |
| `PoAEngineInterface` | `core/Phoenix.Nucleus/PhoenixCore/arbiter` | Governance Team | PhoenixCore, PhoenixDistributed | `contracts/arbiter/v1` | pending |
| `Normalizer` | `core/Phoenix.Nucleus/PhoenixCore/bus/normalizer` | Core Runtime Team | bus, telemetry | `contracts/bus/v1` | pending |
| `SchemaValidator` | `core/Phoenix.Nucleus/PhoenixCore/common/serialization` | Core Runtime Team | event, ledger, validation | `contracts/serialization/v1` | pending |
| `AuditProvider` | `core/Phoenix.Nucleus/PhoenixCore/containment/rollback` | Core Runtime Team | containment, recovery | `contracts/containment/v1` | pending |
| `ILedger` | `core/Phoenix.Nucleus/PhoenixCore/contracts` | Ledger Team | PhoenixValidation, PhoenixTruth | `contracts/ledger/v1` | partial |
| `EventEnvelope` | `core/Phoenix.Nucleus/PhoenixCore/contracts/events/v1` | Architecture Team | PhoenixValidation, PhoenixCore, PhoenixGuard | `contracts/events/v1` | instantiated |
| `Event` | `core/Phoenix.Nucleus/PhoenixCore/contracts/events/v1` | Architecture Team | PhoenixValidation, PhoenixCore, PhoenixGuard | `contracts/events/v1` | instantiated |
| `Serializer` | `core/Phoenix.Nucleus/PhoenixCore/contracts/events/v1` | Architecture Team | PhoenixValidation, PhoenixCore, PhoenixGuard | `contracts/events/v1` | instantiated |
| `Reconstructor` | `core/Phoenix.Nucleus/PhoenixCore/contracts/replay/v1` | Validation Team | PhoenixValidation, PhoenixCore | `contracts/replay/v1` | instantiated |
| `ReplayEngine` (contract) | `core/Phoenix.Nucleus/PhoenixCore/contracts/replay/v1` | Validation Team | PhoenixValidation, PhoenixCore | `contracts/replay/v1` | instantiated |
| `Snapshot` | `core/Phoenix.Nucleus/PhoenixCore/contracts/replay/v1` | Validation Team | PhoenixValidation, PhoenixCore | `contracts/replay/v1` | instantiated |
| `Actuator` (contract) | `core/Phoenix.Nucleus/PhoenixCore/contracts/security/v1` | Security Team | PhoenixGuard, PhoenixValidation | `contracts/security/v1` | instantiated |
| `Containment` | `core/Phoenix.Nucleus/PhoenixCore/contracts/security/v1` | Security Team | PhoenixGuard, PhoenixValidation | `contracts/security/v1` | instantiated |
| `Escalation` | `core/Phoenix.Nucleus/PhoenixCore/contracts/security/v1` | Security Team | PhoenixGuard, PhoenixValidation | `contracts/security/v1` | instantiated |
| `ResourceAllocator` | `core/Phoenix.Nucleus/PhoenixCore/ledger/src` | Ledger Team | ledger, runtime | `contracts/ledger/v1` | pending |
| `AlertManager` | `core/Phoenix.Nucleus/PhoenixCore/monitoring` | Observability Team | runtime, guard, validation | `contracts/observability/v1` | pending |
| `Alerter` | `core/Phoenix.Nucleus/PhoenixCore/monitoring` | Observability Team | runtime, guard, validation | `contracts/observability/v1` | pending |
| `PeerDiscovery` | `core/Phoenix.Nucleus/PhoenixDistributed/discovery` | Core Runtime Team | distributed, runtime | `contracts/distributed/v1` | pending |
| `ConsensusLedger` | `core/Phoenix.Nucleus/PhoenixDistributed/ledger` | Ledger Team | distributed, runtime | `contracts/distributed/v1` | pending |
| `Invariant` | `core/Phoenix.Nucleus/PhoenixGuard` | Security Team | PhoenixGuard, warden | `contracts/security/v1` | pending |
| `IncidentResponseManager` | `core/Phoenix.Nucleus/PhoenixGuard/security` | Security Team | PhoenixGuard, security | `contracts/security/v1` | pending |
| `ThreatDetector` | `core/Phoenix.Nucleus/PhoenixGuard/security` | Security Team | PhoenixGuard, security | `contracts/security/v1` | pending |
| `Actuator` (legacy) | `core/Phoenix.Nucleus/PhoenixGuard` | Security Team | PhoenixGuard, warden | `contracts/security/v1` | pending |
| `CertificateValidator` | `core/Phoenix.Nucleus/PhoenixGuard` | Security Team | PhoenixGuard, warden_compat | `contracts/security/v1` | pending |
| `GraphProvider` | `core/Phoenix.Nucleus/PhoenixGuard` | Security Team | PhoenixGuard, warden_compat | `contracts/security/v1` | pending |
| `EventPublisher` | `core/Phoenix.Nucleus/PhoenixKernel` | Core Runtime Team | kernel, runtime | `contracts/kernel/v1` | pending |
| `PromptGuard` | `core/Phoenix.Nucleus/PhoenixMind/security` | Security Team | cognition, security | `contracts/security/v1` | pending |
