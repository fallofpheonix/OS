---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Repository Cleanup Report

This report classifies all files in the core system context (`core/Phoenix.Nucleus`) to support the contract-first decomposition strategy.

## Classification Summary

| Classification | Count | Description |
| :--- | :--- | :--- |
| **ACTIVE** | 492 | Standard, operational implementation files |
| **LEGACY** | 7 | Backward-compatibility layers |
| **STUB** | 22 | Skeletons, placeholders, or commented-out definitions |
| **EXPERIMENTAL** | 147 | Test files, simulations, and experimental packages |
| **DUPLICATE** | 80 | Basename duplicates requiring consolidation |
| **ARCHIVE** | 0 | Archival folders (none in Phoenix.Nucleus core) |
| **ORPHAN** | 0 | Unused files (no active orphans detected) |

---

## Detailed File Registry

| File | Classification | Details / Reason |
| :--- | :--- | :--- |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [CLAUDE.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/CLAUDE.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [INVARIANTS.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/INVARIANTS.md) | `ACTIVE` | Configuration or Documentation |
| [OWNERSHIP.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/OWNERSHIP.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/README.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ai/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ai/agents/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [explainer.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ai/agents/explainer.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ai/causal/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [enforcer.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ai/causal/enforcer.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [enforcer_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ai/causal/enforcer_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [simulator_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ai/causal/simulator_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/arbiter/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [arbiter_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/arbiter/arbiter_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [consensus_bridge.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/arbiter/consensus_bridge.go) | `ACTIVE` | Standard implementation file |
| [policy_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/arbiter/policy_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [translator.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/arbiter/translator.go) | `ACTIVE` | Standard implementation file |
| [validator.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/arbiter/validator.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/audit/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/auth/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [authentication.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/auth/authentication.go) | `STUB` | Skeleton stub with minimal logic |
| [authorization.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/auth/authorization.go) | `STUB` | Skeleton stub with minimal logic |
| [encryption.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/auth/encryption.go) | `STUB` | Skeleton stub with minimal logic |
| [secrets.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/auth/secrets.go) | `STUB` | Skeleton stub with minimal logic |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/README.md) | `ACTIVE` | Configuration or Documentation |
| [RFC.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/RFC.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/artifacts/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [bus.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/bus.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/logs/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/normalizer/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [normalizer.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/normalizer/normalizer.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [normalizer_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/normalizer/normalizer_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/normalizer_module/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/normalizer_module/README.md) | `ACTIVE` | Configuration or Documentation |
| [RFC.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/normalizer_module/RFC.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/normalizer_module/bench/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/normalizer_module/debug/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/normalizer_module/debug/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/normalizer_module/replay/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/normalizer_module/src/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/normalizer_module/src/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/normalizer_module/tests/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/normalizer_module/tests/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [priority_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/priority_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [signing.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/signing.go) | `ACTIVE` | Standard implementation file |
| [signing_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/signing_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/src/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [bus.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/src/bus.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [bus_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/bus/src/bus_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/artifacts/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [storage.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/artifacts/storage.go) | `ACTIVE` | Standard implementation file |
| [storage_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/artifacts/storage_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/concurrency/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [worker_pool.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/concurrency/worker_pool.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/config/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [redlines.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/config/redlines.go) | `ACTIVE` | Standard implementation file |
| [world_hasher.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/hash/world_hasher.go) | `ACTIVE` | Standard implementation file |
| [world_hasher_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/hash/world_hasher_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/logical_clock/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [clock.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/logical_clock/clock.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/math/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [fixed_point.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/math/fixed_point.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/math/kalman/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [kalman.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/math/kalman/kalman.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/resource/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [allocator.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/resource/allocator.go) | `ACTIVE` | Standard implementation file |
| [quota.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/resource/quota.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/security/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [audit.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/security/audit.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/serialization/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [canonical.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/serialization/canonical.go) | `ACTIVE` | Standard implementation file |
| [endian_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/serialization/endian_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [validator.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/serialization/validator.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/snapshot/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [dag.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/snapshot/dag.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/tooling/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [report.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/common/tooling/report.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/compliance/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [audit.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/compliance/audit.go) | `STUB` | Skeleton stub with minimal logic |
| [framework.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/compliance/framework.go) | `STUB` | Skeleton stub with minimal logic |
| [engine.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/constitution/engine.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [engine_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/constitution/engine_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [invariant.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/constitution/invariant.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [policy.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/constitution/policy.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/file/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [audit.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/file/audit.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [file.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/file/file.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [replay.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/file/replay.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [snapshot.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/file/snapshot.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [isolation.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/isolation.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/network/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [audit.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/network/audit.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [network.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/network/network.go) | `ACTIVE` | Standard implementation file |
| [replay.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/network/replay.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [snapshot.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/network/snapshot.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [policy.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/policy.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [process.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/process.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [process_audit.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/process_audit.go) | `ACTIVE` | Standard implementation file |
| [process_snapshot.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/process_snapshot.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/rollback/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [audit.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/rollback/audit.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [registry.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/rollback/registry.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [replay.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/rollback/replay.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [restore.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/rollback/restore.go) | `ACTIVE` | Standard implementation file |
| [snapshot.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/containment/snapshot.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [compat.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/compat.go) | `LEGACY` | Maintained for backward compatibility |
| [event-bus.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/event-bus.yaml) | `ACTIVE` | Configuration or Documentation |
| [doc.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/events/v1/doc.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [envelope.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/events/v1/envelope.go) | `ACTIVE` | Standard implementation file |
| [event.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/events/v1/event.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [version.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/events/v1/version.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/fsm/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [arbiter.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/fsm/arbiter.yaml) | `ACTIVE` | Configuration or Documentation |
| [distributed.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/fsm/distributed.yaml) | `ACTIVE` | Configuration or Documentation |
| [guard.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/fsm/guard.yaml) | `ACTIVE` | Configuration or Documentation |
| [replay-engine.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/fsm/replay-engine.yaml) | `ACTIVE` | Configuration or Documentation |
| [warden.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/fsm/warden.yaml) | `ACTIVE` | Configuration or Documentation |
| [ledger.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/ledger.yaml) | `ACTIVE` | Configuration or Documentation |
| [doc.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/replay/v1/doc.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [replay.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/replay/v1/replay.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [actuator.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/security/v1/actuator.go) | `ACTIVE` | Standard implementation file |
| [doc.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/contracts/security/v1/doc.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/event/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [schema.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/event/schema.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/examples/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/examples/README.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/game/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [config.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/game/config.go) | `ACTIVE` | Standard implementation file |
| [game_server.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/game/game_server.go) | `ACTIVE` | Standard implementation file |
| [scoring.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/game/scoring.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/governance_engine/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [validator.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/governance_engine/validator.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/hardening/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [core.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/hardening/core.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/invariants/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [architecture-invariants.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/invariants/architecture-invariants.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ledger/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [main.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ledger/main.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ledger/src/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [concurrency_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ledger/src/concurrency_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [ledger.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ledger/src/ledger.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [ledger_allocation_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ledger/src/ledger_allocation_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [ledger_v2_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ledger/src/ledger_v2_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [pruning_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ledger/src/pruning_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [rollback_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/ledger/src/rollback_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/metrics/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitor/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitor/README.md) | `ACTIVE` | Configuration or Documentation |
| [RFC.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitor/RFC.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitor/artifacts/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [entropy.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitor/entropy.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [entropy_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitor/entropy_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [kalman_refinement.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitor/kalman_refinement.go) | `ACTIVE` | Standard implementation file |
| [monitor.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitor/monitor.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitor/replay/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [robust_stats.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitor/robust_stats.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitor/src/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [monitor.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitor/src/monitor.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [monitor_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitor/src/monitor_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitoring/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [alerting.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitoring/alerting.go) | `ACTIVE` | Standard implementation file |
| [logging.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitoring/logging.go) | `STUB` | Skeleton stub with minimal logic |
| [metrics.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitoring/metrics.go) | `STUB` | Skeleton stub with minimal logic |
| [tracing.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/monitoring/tracing.go) | `STUB` | Skeleton stub with minimal logic |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/nexus_coordination/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [bft.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/nexus_coordination/bft.go) | `ACTIVE` | Standard implementation file |
| [consensus.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/nexus_coordination/consensus.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/observability/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/advisory/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [advisory.pb.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/advisory/advisory.pb.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/distributed/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [distributed.pb.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/distributed/distributed.pb.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/enforcement/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [enforcement.pb.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/enforcement/enforcement.pb.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/event/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [event.pb.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/event/event.pb.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/fsm/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [fsm.pb.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/fsm/fsm.pb.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/ledger/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [ledger.pb.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/ledger/ledger.pb.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/memory/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [memory.pb.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/memory/memory.pb.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/simulation/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [simulation.pb.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/simulation/simulation.pb.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/trace/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [trace.pb.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/trace/trace.pb.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/truth/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [truth.pb.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/truth/truth.pb.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/validation/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [validation.pb.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/proto/v1/validation/validation.pb.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/recovery/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [backup.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/recovery/backup.go) | `STUB` | Skeleton stub with minimal logic |
| [disaster_recovery.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/recovery/disaster_recovery.go) | `STUB` | Skeleton stub with minimal logic |
| [engine.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/recovery/engine.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [loop.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/recovery/loop.go) | `ACTIVE` | Standard implementation file |
| [rollback.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/recovery/rollback.go) | `STUB` | Skeleton stub with minimal logic |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/runtime_lab/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [validator.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/runtime_lab/validator.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/scheduler/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [scheduler.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/scheduler/scheduler.go) | `ACTIVE` | Standard implementation file |
| [scheduler_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/scheduler/scheduler_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/schemas/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/schemas/v1/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/schemas/v1/README.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/control/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [fsm.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/control/fsm.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/detections/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/edr/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/game/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/game/stackelberg/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [solver.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/game/stackelberg/solver.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/hunting/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/ids/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/integrated_model/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/integrated_model/src/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [main.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/integrated_model/src/main.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/ips/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/physics/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/physics/disorder/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [sdi.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/physics/disorder/sdi.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [thermo.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/physics/thermo.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/response/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/siem/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/sigma/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/soar/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/xdr/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/security/yara/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/state/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [audit.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/state/audit.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [compat.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/state/compat.go) | `LEGACY` | Maintained for backward compatibility |
| [metrics.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/state/metrics.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [registry.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/state/registry.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [rollback.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/state/rollback.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [rollback_rules.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/state/rollback_rules.go) | `ACTIVE` | Standard implementation file |
| [snapshot.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/state/snapshot.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [transition.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/state/transition.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/tcs/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [degradation.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/tcs/degradation.go) | `ACTIVE` | Standard implementation file |
| [tcs.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/tcs/tcs.go) | `ACTIVE` | Standard implementation file |
| [tcs_dynamic_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/tcs/tcs_dynamic_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/collectors/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/collectors/file_metadata/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [file.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/collectors/file_metadata/file.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/collectors/process_exec/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [process.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/collectors/process_exec/process.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/detector/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [detector.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/detector/detector.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/ebpf/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/entropy_engine/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/entropy_engine/artifacts/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/entropy_engine_go/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [entropy.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/entropy_engine_go/entropy.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [main_bench.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/entropy_engine_go/main_bench.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/events/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [events.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/events/events.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/normalizer/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [normalizer.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/normalizer/normalizer.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/telemetry/pipelines/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/testing/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [chaos_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/testing/chaos_test.go) | `STUB` | Skeleton stub with minimal logic |
| [determinism_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/testing/determinism_test.go) | `STUB` | Skeleton stub with minimal logic |
| [integration_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/testing/integration_test.go) | `STUB` | Skeleton stub with minimal logic |
| [load_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/testing/load_test.go) | `STUB` | Skeleton stub with minimal logic |
| [unit_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/testing/unit_test.go) | `STUB` | Skeleton stub with minimal logic |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/tla/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/tla/README.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/truth/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [evidence.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/truth/evidence.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [fork_detector.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/truth/fork_detector.go) | `ACTIVE` | Standard implementation file |
| [ledger.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/truth/ledger.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [recovery.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/truth/recovery.go) | `ACTIVE` | Standard implementation file |
| [seal.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/truth/seal.go) | `ACTIVE` | Standard implementation file |
| [snapshot.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/truth/snapshot.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/versioning/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [compatibility-matrix.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/versioning/compatibility-matrix.md) | `ACTIVE` | Configuration or Documentation |
| [versioning.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixCore/versioning/versioning.md) | `ACTIVE` | Configuration or Documentation |
| [CLAUDE.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/CLAUDE.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [OWNERSHIP.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/OWNERSHIP.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/README.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/bench/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/consensus/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [poa.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/consensus/poa.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/discovery/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [beacon.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/discovery/beacon.go) | `ACTIVE` | Standard implementation file |
| [discovery.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/discovery/discovery.go) | `ACTIVE` | Standard implementation file |
| [registry.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/discovery/registry.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/docs/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [benchmarks.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/docs/benchmarks.md) | `ACTIVE` | Configuration or Documentation |
| [design.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/docs/design.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/identity/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [node.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/identity/node.go) | `ACTIVE` | Standard implementation file |
| [node_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/identity/node_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/include/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/ledger/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [consensus.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/ledger/consensus.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [distributed_ledger.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/ledger/distributed_ledger.go) | `ACTIVE` | Standard implementation file |
| [ledger.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/ledger/ledger.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [stub.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/ledger/stub.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/replication/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [sync.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/replication/sync.go) | `ACTIVE` | Standard implementation file |
| [sync_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/replication/sync_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/src/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixDistributed/tests/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [ARCHITECTURE_RULES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/ARCHITECTURE_RULES.md) | `ACTIVE` | Configuration or Documentation |
| [CLAUDE.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/CLAUDE.md) | `ACTIVE` | Configuration or Documentation |
| [DEPENDENCY_POLICY.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/DEPENDENCY_POLICY.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [PORT_REGISTRY.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/PORT_REGISTRY.yaml) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/README.md) | `ACTIVE` | Configuration or Documentation |
| [VISIBILITY_POLICY.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/VISIBILITY_POLICY.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/agent-governance/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/agent-governance/invariants/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/agent-governance/verification/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/benchmarks/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/bootstrap/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/bootstrap/templates/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/build-validation/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/contracts/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/contracts/events/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [EMISSION-POINTS.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/contracts/events/EMISSION-POINTS.md) | `ACTIVE` | Configuration or Documentation |
| [IMPLEMENTATION-ROADMAP.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/contracts/events/IMPLEMENTATION-ROADMAP.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/contracts/events/README.md) | `ACTIVE` | Configuration or Documentation |
| [SEQUENCING.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/contracts/events/SEQUENCING.md) | `ACTIVE` | Configuration or Documentation |
| [runtime-event.schema.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/contracts/events/runtime-event.schema.yaml) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/contracts/schemas/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/contracts/schemas/logging/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/extraction/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/git_governance/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/governance/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/health_engine/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DEPENDENCY_INDEX.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/repo_registry/DEPENDENCY_INDEX.yaml) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/repo_registry/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [FORK_INDEX.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/repo_registry/FORK_INDEX.yaml) | `ACTIVE` | Configuration or Documentation |
| [MASTER_REPO_INDEX.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/repo_registry/MASTER_REPO_INDEX.yaml) | `ACTIVE` | Configuration or Documentation |
| [MODULE_INDEX.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/repo_registry/MODULE_INDEX.yaml) | `ACTIVE` | Configuration or Documentation |
| [SERVICE_INDEX.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/repo_registry/SERVICE_INDEX.yaml) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/sync_engine/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [dependency-graph.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/sync_engine/dependency-graph.yaml) | `ACTIVE` | Configuration or Documentation |
| [repo-registry.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/sync_engine/repo-registry.yaml) | `ACTIVE` | Configuration or Documentation |
| [sync-config.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/sync_engine/sync-config.yaml) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/tracking/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [extractions.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/control_plane/tracking/extractions.yaml) | `ACTIVE` | Configuration or Documentation |
| [DEPENDENCY_GRAPH.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/dependency-map/DEPENDENCY_GRAPH.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/dependency-map/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [EXTRACTION_CANDIDATES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/dependency-map/EXTRACTION_CANDIDATES.md) | `ACTIVE` | Configuration or Documentation |
| [MODULE_USAGE_MAP.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/dependency-map/MODULE_USAGE_MAP.md) | `ACTIVE` | Configuration or Documentation |
| [PORT_ALLOCATION_MAP.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/dependency-map/PORT_ALLOCATION_MAP.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/events/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/runtime-state/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/runtime-state/README.md) | `ACTIVE` | Configuration or Documentation |
| [index.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/runtime-state/index.yaml) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/runtime/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/schemas/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/specs/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixFormal/tla/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [CLAUDE.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/CLAUDE.md) | `ACTIVE` | Configuration or Documentation |
| [DELIVERY.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/DELIVERY.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [OWNERSHIP.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/OWNERSHIP.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/README.md) | `ACTIVE` | Configuration or Documentation |
| [TRUST_MATRIX.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/TRUST_MATRIX.yaml) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/actuation/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [executor.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/actuation/executor.go) | `ACTIVE` | Standard implementation file |
| [executor_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/actuation/executor_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [sandbox.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/actuation/sandbox.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/actuators/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [ebpf.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/actuators/ebpf.go) | `ACTIVE` | Standard implementation file |
| [process.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/actuators/process.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/audit/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [jsonl_writer.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/audit/jsonl_writer.go) | `ACTIVE` | Standard implementation file |
| [violation_log.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/audit/violation_log.go) | `ACTIVE` | Standard implementation file |
| [violation_log_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/audit/violation_log_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [cert_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/cert_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [runtime-event.schema.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/contracts/events/runtime-event.schema.yaml) | `ACTIVE` | Configuration or Documentation |
| [control.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/control.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/control/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [fsm.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/control/fsm.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/core/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/core/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [ARCHITECTURAL_GUARDRAILS.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/docs/ARCHITECTURAL_GUARDRAILS.md) | `ACTIVE` | Configuration or Documentation |
| [ARCHITECTURE.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/docs/ARCHITECTURE.md) | `ACTIVE` | Configuration or Documentation |
| [COMPOSITE_OPERATION_SPEC.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/docs/COMPOSITE_OPERATION_SPEC.md) | `ACTIVE` | Configuration or Documentation |
| [COMPOSITION_PRESSURE_EVALUATION.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/docs/COMPOSITION_PRESSURE_EVALUATION.md) | `ACTIVE` | Configuration or Documentation |
| [COMPOSITION_READINESS_ANALYSIS.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/docs/COMPOSITION_READINESS_ANALYSIS.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/docs/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [INVARIANTS.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/docs/INVARIANTS.md) | `ACTIVE` | Configuration or Documentation |
| [OBSERVATION_CHECKLIST.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/docs/OBSERVATION_CHECKLIST.md) | `ACTIVE` | Configuration or Documentation |
| [ORCHESTRATION_IMPLEMENTATION.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/docs/ORCHESTRATION_IMPLEMENTATION.md) | `ACTIVE` | Configuration or Documentation |
| [ORCHESTRATION_SEMANTICS.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/docs/ORCHESTRATION_SEMANTICS.md) | `ACTIVE` | Configuration or Documentation |
| [PHASE_2_PRESERVATION.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/docs/PHASE_2_PRESERVATION.md) | `ACTIVE` | Configuration or Documentation |
| [PSYCHOLOGICAL_ARCHITECTURE.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/docs/PSYCHOLOGICAL_ARCHITECTURE.md) | `ACTIVE` | Configuration or Documentation |
| [RELEASE.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/docs/RELEASE.md) | `ACTIVE` | Configuration or Documentation |
| [USAGE.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/docs/USAGE.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/emergency/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [killswitch.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/emergency/killswitch.go) | `ACTIVE` | Standard implementation file |
| [killswitch_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/emergency/killswitch_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/engine/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [warden.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/engine/warden.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [warden_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/engine/warden_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/game/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/game/stackelberg/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [solver.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/game/stackelberg/solver.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/hunting/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/infrastructure/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/infrastructure/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/infrastructure/logging/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/infrastructure/logging/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/integrated_model/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/integrated_model/src/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [main.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/integrated_model/src/main.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/interfaces/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/interfaces/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/interfaces/cli/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/interfaces/cli/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [invariant.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/invariant.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [invariant_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/invariant_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/physics/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/physics/disorder/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [sdi.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/physics/disorder/sdi.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [thermo.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/physics/thermo.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/policies/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [trust_matrix.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/policies/trust_matrix.go) | `ACTIVE` | Standard implementation file |
| [trust_matrix_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/policies/trust_matrix_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/response/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/runtime/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/runtime/README.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/runtime/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/runtime/filesystem/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/runtime/filesystem/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/runtime/orchestration/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/runtime/orchestration/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/runtime/shell/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/runtime/shell/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/runtime/tracing/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/runtime/tracing/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/security/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [incident_response.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/security/incident_response.go) | `ACTIVE` | Standard implementation file |
| [threat_detection.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/security/threat_detection.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/tests/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/tests/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/tests/integration/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/tests/integration/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/tests/runtime/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/tests/runtime/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/validation/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [security_validator.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/validation/security_validator.go) | `ACTIVE` | Standard implementation file |
| [warden.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/warden.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [warden_compat.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/warden_compat.go) | `LEGACY` | Maintained for backward compatibility |
| [warden_compat_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/warden_compat_test.go) | `LEGACY` | Maintained for backward compatibility |
| [warden_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixGuard/warden_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [CLAUDE.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/CLAUDE.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [OWNERSHIP.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/OWNERSHIP.md) | `ACTIVE` | Configuration or Documentation |
| [REPO_OVERVIEW.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/REPO_OVERVIEW.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/boot/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/bridge/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [bridge_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/bridge/bridge_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [event_stream.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/bridge/event_stream.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/documents/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/documents/vault/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/drivers/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [ebpf_loader.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/ebpf_loader.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [ebpf_probe.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/ebpf_probe.go) | `ACTIVE` | Standard implementation file |
| [ebpf_test_harness.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/ebpf_test_harness.go) | `ACTIVE` | Standard implementation file |
| [enforcer.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/enforcer.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/git-hooks/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/golang/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/golang/golangci-lint/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/hooks/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [ebpf_enforcer.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/hooks/ebpf_enforcer.go) | `ACTIVE` | Standard implementation file |
| [isolation.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/hooks/isolation.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [profiles.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/hooks/profiles.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/ipc/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/live/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [clock_normalizer.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/live/clock_normalizer.go) | `ACTIVE` | Standard implementation file |
| [clock_sync.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/live/clock_sync.go) | `ACTIVE` | Standard implementation file |
| [ebpf_loader.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/live/ebpf_loader.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [event_decoder.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/live/event_decoder.go) | `ACTIVE` | Standard implementation file |
| [live_determinism_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/live/live_determinism_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [overflow_handler.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/live/overflow_handler.go) | `ACTIVE` | Standard implementation file |
| [probe_loader.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/live/probe_loader.go) | `ACTIVE` | Standard implementation file |
| [probe_manager.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/live/probe_manager.go) | `ACTIVE` | Standard implementation file |
| [ring_reader.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/live/ring_reader.go) | `ACTIVE` | Standard implementation file |
| [sync_manager.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/live/sync_manager.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/memory/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/observability/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/observability/otel/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [otel-config.yaml](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/observability/otel/otel-config.yaml) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/observability/prometheus/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/probes/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [gen.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/probes/gen.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/probes/src/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/prototypes/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/python/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/python/mypy/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/python/ruff/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/python/ruff/profiles/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/runtime/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [affinity_runner.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/runtime/affinity_runner.go) | `ACTIVE` | Standard implementation file |
| [cgroups.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/runtime/cgroups.go) | `ACTIVE` | Standard implementation file |
| [cgroups_stub.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/runtime/cgroups_stub.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [cgroups_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/runtime/cgroups_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [clock_skew.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/runtime/clock_skew.go) | `ACTIVE` | Standard implementation file |
| [namespaces.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/runtime/namespaces.go) | `ACTIVE` | Standard implementation file |
| [namespaces_stub.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/runtime/namespaces_stub.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [probe_injector.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/runtime/probe_injector.go) | `ACTIVE` | Standard implementation file |
| [replay_bridge.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/runtime/replay_bridge.go) | `ACTIVE` | Standard implementation file |
| [ring_monitor.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/runtime/ring_monitor.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/rust/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/rust/cargo/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/rust/clippy/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/rust/rustfmt/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/sandbox/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [enforcement.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/sandbox/enforcement.go) | `ACTIVE` | Standard implementation file |
| [simulator.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/sandbox/simulator.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/scheduler/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/scripts/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/shared-libraries/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/shared-libraries/auth-orchestrator/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/shared-libraries/transaction-processor/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/src/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/templates/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [ARCHITECTURE.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/templates/project-base/ARCHITECTURE.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/templates/project-base/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/templates/project-base/README.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/templates/project-base/scripts/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/tooling/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/tooling/vault-root-hidden/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [types.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/types.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixKernel/vfs/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixMind/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixMind/security/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [prompt_security.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixMind/security/prompt_security.go) | `ACTIVE` | Standard implementation file |
| [CLAUDE.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/CLAUDE.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [OWNERSHIP.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/OWNERSHIP.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/README.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/engine/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [dag.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/engine/dag.go) | `STUB` | Skeleton stub with minimal logic |
| [engine_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/engine/engine_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [lineage.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/engine/lineage.go) | `STUB` | Skeleton stub with minimal logic |
| [mapper.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/engine/mapper.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/engine/process_graphs/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/engine/process_graphs/artifacts/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [process_graphs.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/engine/process_graphs/process_graphs.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/engine/process_lineage/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [lineage.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/engine/process_lineage/lineage.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/engine/syscall_graphs/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTrace/engine/traces/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [CLAUDE.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTruth/CLAUDE.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTruth/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [OWNERSHIP.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTruth/OWNERSHIP.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTruth/README.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTruth/engine/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [confidence.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTruth/engine/confidence.go) | `ACTIVE` | Standard implementation file |
| [contradiction.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTruth/engine/contradiction.go) | `ACTIVE` | Standard implementation file |
| [engine_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTruth/engine/engine_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [evaluator.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTruth/engine/evaluator.go) | `ACTIVE` | Standard implementation file |
| [evidence.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTruth/engine/evidence.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [truth.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixTruth/engine/truth.go) | `ACTIVE` | Standard implementation file |
| [CLAUDE.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/CLAUDE.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [OWNERSHIP.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/OWNERSHIP.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/README.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/__pycache__/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/chaos/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [fuzz_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/chaos/fuzz_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/contract-tests/README.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/determinism/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [determinism.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/determinism/determinism.go) | `STUB` | Skeleton stub with minimal logic |
| [determinism_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/determinism/determinism_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [distribution_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/distribution_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/evidence/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [chain_integrity_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/evidence/chain_integrity_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [fork_reject_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/evidence/fork_reject_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [ledger_repeat_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/evidence/ledger_repeat_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [mutation_block_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/evidence/mutation_block_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [seal_consistency_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/evidence/seal_consistency_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [snapshot_recovery_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/evidence/snapshot_recovery_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/formal/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [hash_chain_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/formal/hash_chain_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [ledger_invariant_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/formal/ledger_invariant_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [ordering_invariant_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/formal/ordering_invariant_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [replay_consistency_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/formal/replay_consistency_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [rollback_consistency_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/formal/rollback_consistency_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [infrastructure_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/infrastructure_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/integration/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/integration/ai/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [brain_integrity_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/integration/ai/brain_integrity_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/integration/runtime/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [runtime_graph_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/integration/runtime/runtime_graph_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/invariants/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [clock_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/invariants/clock_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [affinity_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/affinity_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [clock_skew_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/clock_skew_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [cross_core_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/cross_core_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [drop_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/drop_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [ebpf_ring_stress_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/ebpf_ring_stress_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [event_loss_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/event_loss_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [live_probe_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/live_probe_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [logical_clock_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/logical_clock_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [ordering_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/ordering_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [overflow_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/overflow_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [replay_kernel_sync_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/replay_kernel_sync_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [replay_sync_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/replay_sync_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [runtime_ring_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/runtime_ring_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [simulator_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/kernel/simulator_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proof/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proof/proofs/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [replay_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proof/proofs/replay_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [rollback_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proof/proofs/rollback_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [state_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proof/proofs/state_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [transition_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proof/proofs/transition_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proofs/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [containment_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proofs/containment_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [federation_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proofs/federation_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [ledger_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proofs/ledger_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [ordering_proof.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proofs/ordering_proof.go) | `ACTIVE` | Standard implementation file |
| [recovery_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proofs/recovery_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [replay_identity_proof.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proofs/replay_identity_proof.go) | `ACTIVE` | Standard implementation file |
| [replay_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proofs/replay_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [rollback_proof.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proofs/rollback_proof.go) | `ACTIVE` | Standard implementation file |
| [transition_proof.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/proofs/transition_proof.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/replay/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [authority.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/replay/authority.go) | `ACTIVE` | Standard implementation file |
| [engine.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/replay/engine.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [engine_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/replay/engine_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [replay.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/replay/replay.go) | `STUB` | Skeleton stub with minimal logic |
| [replay_execution_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/replay/replay_execution_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [verifier.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/replay/verifier.go) | `STUB` | Skeleton stub with minimal logic |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/replay/verifier/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [verifier.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/replay/verifier/verifier.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/runtime/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [brain_integrity_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/runtime/brain_integrity_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [recursive_debug_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/runtime/recursive_debug_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/runtime_graph/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [arbiter_warden.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/runtime_graph/arbiter_warden.go) | `ACTIVE` | Standard implementation file |
| [containment_recovery.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/runtime_graph/containment_recovery.go) | `ACTIVE` | Standard implementation file |
| [illegal_path.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/runtime_graph/illegal_path.go) | `ACTIVE` | Standard implementation file |
| [replay_truth.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/runtime_graph/replay_truth.go) | `ACTIVE` | Standard implementation file |
| [telemetry_replay.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/runtime_graph/telemetry_replay.go) | `ACTIVE` | Standard implementation file |
| [truth_arbiter.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/runtime_graph/truth_arbiter.go) | `ACTIVE` | Standard implementation file |
| [warden_containment.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/runtime_graph/warden_containment.go) | `ACTIVE` | Standard implementation file |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [batch_d_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/batch_d_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [beacon_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/beacon_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [bus_exploit_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/bus_exploit_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [containment_attack_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/containment_attack_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [exfil_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/exfil_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [forkbomb_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/forkbomb_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [fsm_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/fsm_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [mutation_attack_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/mutation_attack_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [portscan_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/portscan_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [privilege_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/privilege_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [replay_attack_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/replay_attack_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [solver_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/solver_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [tamper_attack_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/tamper_attack_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [tcs_exploit_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/tcs_exploit_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [thermo_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/thermo_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [timeline_attack_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/security/timeline_attack_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/soak/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [containment_24h_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/soak/containment_24h_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [drift_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/soak/drift_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [recovery_24h_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/soak/recovery_24h_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [replay_24h_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/soak/replay_24h_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/truth/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [truth_verification_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/truth/truth_verification_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [boot_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/boot_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [bus_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/bus_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [concurrency_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/concurrency_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [containment_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/containment_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [detector_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/detector_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [determinism_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/determinism_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [endian_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/endian_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [entropy_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/entropy_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [file_determinism_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/file_determinism_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [file_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/file_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [file_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/file_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [ledger_allocation_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/ledger_allocation_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [ledger_compat.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/ledger_compat.go) | `LEGACY` | Maintained for backward compatibility |
| [ledger_v2_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/ledger_v2_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [lineage_compat.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/lineage_compat.go) | `LEGACY` | Maintained for backward compatibility |
| [lineage_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/lineage_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [metrics_export_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/metrics_export_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [metrics_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/metrics_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [network_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/network_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [network_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/network_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [normalizer_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/normalizer_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [policy_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/policy_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [priority_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/priority_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [process_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/process_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [process_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/process_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [pruning_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/pruning_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [recovery_hardening_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/recovery_hardening_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [recovery_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/recovery_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [registry_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/registry_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [replay_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/replay_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [rollback_determinism_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/rollback_determinism_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [rollback_integration_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/rollback_integration_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [rollback_policy_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/rollback_policy_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [rollback_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/rollback_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [rollback_repeatability_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/rollback_repeatability_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [rollback_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/rollback_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [runtime_graph_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/runtime_graph_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [snapshot_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/snapshot_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [state_compat.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/state_compat.go) | `LEGACY` | Maintained for backward compatibility |
| [state_proof_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/PhoenixValidation/unit/state_proof_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/README.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/authority/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [audit.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/authority/audit.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [audit_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/authority/audit_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [conservation.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/authority/conservation.go) | `ACTIVE` | Standard implementation file |
| [manager.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/authority/manager.go) | `ACTIVE` | Standard implementation file |
| [policy.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/authority/policy.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/capability/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [token.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/capability/token.go) | `ACTIVE` | Standard implementation file |
| [token_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/capability/token_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/docs/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [ADR-001-Canonical-Contracts.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/docs/adr/ADR-001-Canonical-Contracts.md) | `ACTIVE` | Configuration or Documentation |
| [ADR-002-Warden-FSM.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/docs/adr/ADR-002-Warden-FSM.md) | `ACTIVE` | Configuration or Documentation |
| [ADR-003-Event-Bus.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/docs/adr/ADR-003-Event-Bus.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/docs/adr/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [README.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/docs/adr/README.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/fracture/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/ledger/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [chain.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/ledger/chain.go) | `ACTIVE` | Standard implementation file |
| [chain_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/ledger/chain_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [event.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/ledger/event.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [genesis.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/ledger/genesis.go) | `ACTIVE` | Standard implementation file |
| [persistor.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/ledger/persistor.go) | `ACTIVE` | Standard implementation file |
| [persistor_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/ledger/persistor_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [reality.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/ledger/reality.go) | `ACTIVE` | Standard implementation file |
| [reality_test.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/ledger/reality_test.go) | `EXPERIMENTAL` | Test harness or simulation module |
| [snapshot.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/ledger/snapshot.go) | `DUPLICATE` | Basename duplicate; potential redundancy |
| [DIRECTORY_NOTES.md](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/recovery/DIRECTORY_NOTES.md) | `ACTIVE` | Configuration or Documentation |
| [boot.go](file:////Users/fallofpheonix/os/core/Phoenix.Nucleus/recovery/boot.go) | `ACTIVE` | Standard implementation file |
