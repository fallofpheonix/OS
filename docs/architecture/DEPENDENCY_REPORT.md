---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Dependency and Coupling Report

Evaluation of coupling, layering violations, and circular references in the core repository.

## Layering Violations Detected

1. **Recovery to Replay Circular Coupling**:
   - **Path**: `core/Phoenix.Nucleus/PhoenixCore/recovery/engine.go`
   - **Imports**: `github.com/fallofpheonix/PhoenixValidation/replay`
   - **Violation**: The core runtime `recovery` package directly imports the `validation/replay` package, creating a circular loop since validation imports the core runtime.
   - **Remediation**: Refactor `recovery` to import the interface `contracts/replay/v1.ReplayEngine` instead of the validation package.

2. **Game Server to Guard Layering Violation**:
   - **Path**: `core/Phoenix.Nucleus/PhoenixCore/game/game_server.go`
   - **Imports**: `github.com/fallofpheonix/PhoenixGuard/engine`
   - **Violation**: Platform/API zone directly imports internal implementation details (`*engine.Warden`) of `PhoenixGuard`, bypassing the contract boundaries.
   - **Remediation**: Inject `contracts/security/v1.Actuator` instead.

3. **Tooling to Replay Coupling**:
   - **Path**: `core/Phoenix.Nucleus/PhoenixCore/common/tooling/report.go`
   - **Imports**: `github.com/fallofpheonix/PhoenixValidation/replay`
   - **Violation**: Tooling directly imports validation implementation internals.
   - **Remediation**: Use `contracts/replay/v1.ReplayEngine` to generate reports.

---

## Detailed Workspace Dependencies

| Source Path | Imports / External Boundaries |
| :--- | :--- |
| `core/Phoenix.Nucleus/PhoenixCore/ai/agents/explainer.go` | `github.com/fallofpheonix/PhoenixCore/ledger/src` |
| `core/Phoenix.Nucleus/PhoenixCore/common/concurrency/worker_pool.go` | `github.com/fallofpheonix/PhoenixCore/common/resource` |
| `core/Phoenix.Nucleus/PhoenixCore/common/hash/world_hasher_test.go` | `github.com/fallofpheonix/PhoenixCore/common/hash` |
| `core/Phoenix.Nucleus/PhoenixCore/common/security/audit.go` | `github.com/fallofpheonix/PhoenixCore/ledger/src` |
| `core/Phoenix.Nucleus/PhoenixCore/common/tooling/report.go` | `github.com/fallofpheonix/PhoenixCore/ledger/src`, `github.com/fallofpheonix/PhoenixValidation/replay` |
| `core/Phoenix.Nucleus/PhoenixCore/constitution/engine.go` | `github.com/fallofpheonix/PhoenixCore/event` |
| `core/Phoenix.Nucleus/PhoenixCore/constitution/engine_test.go` | `github.com/fallofpheonix/PhoenixCore/event` |
| `core/Phoenix.Nucleus/PhoenixCore/constitution/invariant.go` | `github.com/fallofpheonix/PhoenixCore/event` |
| `core/Phoenix.Nucleus/PhoenixCore/constitution/policy.go` | `github.com/fallofpheonix/PhoenixCore/event` |
| `core/Phoenix.Nucleus/PhoenixCore/containment/rollback/restore.go` | `github.com/fallofpheonix/PhoenixCore/containment`, `github.com/fallofpheonix/PhoenixCore/containment/file`, `github.com/fallofpheonix/PhoenixCore/containment/network` |
| `core/Phoenix.Nucleus/PhoenixCore/contracts/replay/v1/replay.go` | `github.com/fallofpheonix/PhoenixCore/contracts/events/v1` |
| `core/Phoenix.Nucleus/PhoenixCore/game/game_server.go` | `github.com/fallofpheonix/PhoenixCore/ledger/src`, `github.com/fallofpheonix/PhoenixGuard/engine` |
| `core/Phoenix.Nucleus/PhoenixCore/ledger/src/concurrency_test.go` | `github.com/fallofpheonix/PhoenixCore/common/resource` |
| `core/Phoenix.Nucleus/PhoenixCore/ledger/src/ledger_v2_test.go` | `github.com/fallofpheonix/PhoenixCore/common/resource` |
| `core/Phoenix.Nucleus/PhoenixCore/ledger/src/pruning_test.go` | `github.com/fallofpheonix/PhoenixCore/common/resource` |
| `core/Phoenix.Nucleus/PhoenixCore/ledger/src/rollback_test.go` | `github.com/fallofpheonix/PhoenixCore/common/resource` |
| `core/Phoenix.Nucleus/PhoenixCore/monitor/monitor.go` | `github.com/fallofpheonix/PhoenixCore/bus`, `github.com/fallofpheonix/PhoenixCore/common/math/kalman` |
| `core/Phoenix.Nucleus/PhoenixCore/recovery/engine.go` | `github.com/fallofpheonix/PhoenixCore/constitution`, `github.com/fallofpheonix/PhoenixCore/event`, `github.com/fallofpheonix/PhoenixValidation/replay` |
| `core/Phoenix.Nucleus/PhoenixCore/recovery/loop.go` | `github.com/fallofpheonix/PhoenixCore/bus`, `github.com/fallofpheonix/PhoenixCore/containment/rollback` |
| `core/Phoenix.Nucleus/PhoenixCore/security/integrated_model/src/main.go` | `github.com/fallofpheonix/PhoenixCore/security/physics`, `github.com/fallofpheonix/PhoenixCore/telemetry/entropy_engine_go`, `github.com/fallofpheonix/PhoenixTrace/engine/process_graphs` |
| `core/Phoenix.Nucleus/PhoenixCore/security/physics/thermo.go` | `github.com/fallofpheonix/PhoenixCore/security/physics/disorder` |
| `core/Phoenix.Nucleus/PhoenixCore/telemetry/detector/detector.go` | `github.com/fallofpheonix/PhoenixCore/telemetry/events`, `github.com/fallofpheonix/PhoenixTrace/engine/process_lineage` |
| `core/Phoenix.Nucleus/PhoenixCore/telemetry/normalizer/normalizer.go` | `github.com/fallofpheonix/PhoenixCore/bus`, `github.com/fallofpheonix/PhoenixCore/common/logical_clock` |
| `core/Phoenix.Nucleus/PhoenixDistributed/consensus/poa.go` | `github.com/fallofpheonix/PhoenixDistributed/identity` |
| `core/Phoenix.Nucleus/PhoenixDistributed/ledger/consensus.go` | `github.com/fallofpheonix/PhoenixCore/ledger/src`, `github.com/fallofpheonix/PhoenixDistributed/discovery` |
| `core/Phoenix.Nucleus/PhoenixDistributed/ledger/distributed_ledger.go` | `github.com/fallofpheonix/PhoenixCore/ledger/src` |
| `core/Phoenix.Nucleus/PhoenixGuard/actuation/sandbox.go` | `github.com/fallofpheonix/PhoenixKernel/runtime` |
| `core/Phoenix.Nucleus/PhoenixGuard/actuators/ebpf.go` | `github.com/fallofpheonix/PhoenixGuard`, `github.com/fallofpheonix/PhoenixKernel` |
| `core/Phoenix.Nucleus/PhoenixGuard/actuators/process.go` | `github.com/fallofpheonix/PhoenixGuard` |
| `core/Phoenix.Nucleus/PhoenixGuard/cert_test.go` | `github.com/fallofpheonix/PhoenixCore/common/resource`, `github.com/fallofpheonix/PhoenixCore/ledger/src` |
| `core/Phoenix.Nucleus/PhoenixGuard/emergency/killswitch.go` | `github.com/fallofpheonix/PhoenixGuard/engine` |
| `core/Phoenix.Nucleus/PhoenixGuard/emergency/killswitch_test.go` | `github.com/fallofpheonix/PhoenixGuard/engine` |
| `core/Phoenix.Nucleus/PhoenixGuard/integrated_model/src/main.go` | `github.com/fallofpheonix/PhoenixCore/security/physics`, `github.com/fallofpheonix/PhoenixCore/telemetry/entropy_engine`, `github.com/fallofpheonix/PhoenixTrace/engine/process_graphs` |
| `core/Phoenix.Nucleus/PhoenixGuard/physics/thermo.go` | `github.com/fallofpheonix/PhoenixCore/security/physics/disorder` |
| `core/Phoenix.Nucleus/PhoenixGuard/warden.go` | `github.com/fallofpheonix/PhoenixCore/bus`, `github.com/fallofpheonix/PhoenixCore/common/config` |
| `core/Phoenix.Nucleus/PhoenixGuard/warden_compat.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixGuard/warden_compat_test.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixGuard/warden_test.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixKernel/bridge/bridge_test.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixKernel/bridge/event_stream.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixKernel/runtime/probe_injector.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixKernel/runtime/replay_bridge.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixKernel/runtime/ring_monitor.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixTruth/engine/contradiction.go` | `github.com/fallofpheonix/PhoenixCore/proto/v1/ledger`, `github.com/fallofpheonix/PhoenixCore/proto/v1/truth` |
| `core/Phoenix.Nucleus/PhoenixTruth/engine/engine_test.go` | `github.com/fallofpheonix/PhoenixCore/proto/v1/ledger`, `github.com/fallofpheonix/PhoenixTruth/engine` |
| `core/Phoenix.Nucleus/PhoenixTruth/engine/evaluator.go` | `github.com/fallofpheonix/PhoenixCore/proto/v1/ledger`, `github.com/fallofpheonix/PhoenixCore/proto/v1/truth` |
| `core/Phoenix.Nucleus/PhoenixValidation/chaos/fuzz_test.go` | `github.com/fallofpheonix/PhoenixCore/proto/v1/event`, `github.com/fallofpheonix/PhoenixValidation/replay` |
| `core/Phoenix.Nucleus/PhoenixValidation/integration/ai/brain_integrity_test.go` | `github.com/fallofpheonix/PhoenixCore/bus`, `github.com/fallofpheonix/PhoenixMind/intelligence` |
| `core/Phoenix.Nucleus/PhoenixValidation/invariants/clock_test.go` | `github.com/fallofpheonix/PhoenixCore/proto/v1/event`, `github.com/fallofpheonix/PhoenixValidation/replay` |
| `core/Phoenix.Nucleus/PhoenixValidation/kernel/affinity_test.go` | `github.com/fallofpheonix/PhoenixKernel/runtime` |
| `core/Phoenix.Nucleus/PhoenixValidation/kernel/clock_skew_test.go` | `github.com/fallofpheonix/PhoenixKernel/runtime` |
| `core/Phoenix.Nucleus/PhoenixValidation/kernel/ebpf_ring_stress_test.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixValidation/kernel/event_loss_test.go` | `github.com/fallofpheonix/PhoenixKernel/sandbox` |
| `core/Phoenix.Nucleus/PhoenixValidation/kernel/live_probe_test.go` | `github.com/fallofpheonix/PhoenixCore/bus`, `github.com/fallofpheonix/PhoenixKernel/runtime` |
| `core/Phoenix.Nucleus/PhoenixValidation/kernel/logical_clock_test.go` | `github.com/fallofpheonix/PhoenixCore/common/logical_clock` |
| `core/Phoenix.Nucleus/PhoenixValidation/kernel/replay_kernel_sync_test.go` | `github.com/fallofpheonix/PhoenixKernel/sandbox` |
| `core/Phoenix.Nucleus/PhoenixValidation/kernel/runtime_ring_test.go` | `github.com/fallofpheonix/PhoenixCore/bus`, `github.com/fallofpheonix/PhoenixKernel/runtime` |
| `core/Phoenix.Nucleus/PhoenixValidation/kernel/simulator_test.go` | `github.com/fallofpheonix/PhoenixKernel/sandbox` |
| `core/Phoenix.Nucleus/PhoenixValidation/proofs/containment_proof_test.go` | `github.com/fallofpheonix/PhoenixGuard` |
| `core/Phoenix.Nucleus/PhoenixValidation/proofs/federation_proof_test.go` | `github.com/fallofpheonix/PhoenixDistributed/identity` |
| `core/Phoenix.Nucleus/PhoenixValidation/proofs/ledger_proof_test.go` | `github.com/fallofpheonix/Phoenix.Nucleus/ledger` |
| `core/Phoenix.Nucleus/PhoenixValidation/proofs/recovery_proof_test.go` | `github.com/fallofpheonix/PhoenixCore/constitution`, `github.com/fallofpheonix/PhoenixCore/event`, `github.com/fallofpheonix/PhoenixCore/recovery`, `github.com/fallofpheonix/PhoenixValidation/replay` |
| `core/Phoenix.Nucleus/PhoenixValidation/proofs/replay_proof_test.go` | `github.com/fallofpheonix/PhoenixCore/event`, `github.com/fallofpheonix/PhoenixValidation/replay` |
| `core/Phoenix.Nucleus/PhoenixValidation/replay/engine.go` | `github.com/fallofpheonix/PhoenixCore/common/hash`, `github.com/fallofpheonix/PhoenixCore/common/serialization`, `github.com/fallofpheonix/PhoenixCore/event`, `github.com/fallofpheonix/PhoenixCore/contracts/events/v1` |
| `core/Phoenix.Nucleus/PhoenixValidation/replay/engine_test.go` | `github.com/fallofpheonix/PhoenixCore/proto/v1/event`, `github.com/fallofpheonix/PhoenixCore/proto/v1/fsm` |
| `core/Phoenix.Nucleus/PhoenixValidation/replay/replay_execution_test.go` | `github.com/fallofpheonix/PhoenixCore/containment`, `github.com/fallofpheonix/PhoenixCore/containment/file`, `github.com/fallofpheonix/PhoenixCore/containment/network`, `github.com/fallofpheonix/PhoenixCore/containment/rollback` |
| `core/Phoenix.Nucleus/PhoenixValidation/replay/verifier/verifier.go` | `github.com/fallofpheonix/PhoenixCore/bus`, `github.com/fallofpheonix/PhoenixCore/event`, `github.com/fallofpheonix/PhoenixValidation/replay` |
| `core/Phoenix.Nucleus/PhoenixValidation/runtime/brain_integrity_test.go` | `github.com/fallofpheonix/PhoenixMind/intelligence` |
| `core/Phoenix.Nucleus/PhoenixValidation/security/beacon_test.go` | `github.com/fallofpheonix/PhoenixCore/bus`, `github.com/fallofpheonix/PhoenixCore/containment` |
| `core/Phoenix.Nucleus/PhoenixValidation/security/bus_exploit_test.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixValidation/security/exfil_test.go` | `github.com/fallofpheonix/PhoenixCore/bus`, `github.com/fallofpheonix/PhoenixCore/containment` |
| `core/Phoenix.Nucleus/PhoenixValidation/security/forkbomb_test.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixValidation/security/fsm_test.go` | `github.com/fallofpheonix/PhoenixGuard/control` |
| `core/Phoenix.Nucleus/PhoenixValidation/security/portscan_test.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixValidation/security/solver_test.go` | `github.com/fallofpheonix/PhoenixCore/security/game/stackelberg` |
| `core/Phoenix.Nucleus/PhoenixValidation/security/tcs_exploit_test.go` | `github.com/fallofpheonix/PhoenixCore/tcs` |
| `core/Phoenix.Nucleus/PhoenixValidation/security/thermo_test.go` | `github.com/fallofpheonix/PhoenixCore/security/physics` |
| `core/Phoenix.Nucleus/PhoenixValidation/security/timeline_attack_test.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixValidation/truth/truth_verification_test.go` | `github.com/fallofpheonix/PhoenixCore/ledger/src` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/boot_test.go` | `github.com/fallofpheonix/PhoenixCore/arbiter`, `github.com/fallofpheonix/PhoenixCore/common/resource`, `github.com/fallofpheonix/PhoenixCore/ledger/src`, `github.com/fallofpheonix/PhoenixGuard/engine`, `github.com/fallofpheonix/phoenix-os/phoenix_os/boot` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/bus_test.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/concurrency_test.go` | `github.com/fallofpheonix/PhoenixCore/common/resource`, `github.com/fallofpheonix/PhoenixCore/ledger/src` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/containment_proof_test.go` | `github.com/fallofpheonix/PhoenixCore/containment` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/detector_test.go` | `github.com/fallofpheonix/PhoenixCore/telemetry/detector`, `github.com/fallofpheonix/PhoenixCore/telemetry/events`, `github.com/fallofpheonix/PhoenixTrace/engine/process_lineage` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/entropy_test.go` | `github.com/fallofpheonix/PhoenixCore/telemetry/entropy_engine_go` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/file_determinism_test.go` | `github.com/fallofpheonix/PhoenixCore/containment/file` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/file_proof_test.go` | `github.com/fallofpheonix/PhoenixCore/containment/file` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/file_test.go` | `github.com/fallofpheonix/PhoenixCore/containment/file` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/ledger_compat.go` | `github.com/fallofpheonix/PhoenixCore/ledger/src` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/ledger_v2_test.go` | `github.com/fallofpheonix/PhoenixCore/common/resource` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/lineage_compat.go` | `github.com/fallofpheonix/PhoenixTrace/engine/process_lineage` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/network_proof_test.go` | `github.com/fallofpheonix/PhoenixCore/containment/network` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/network_test.go` | `github.com/fallofpheonix/PhoenixCore/containment/network` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/normalizer_test.go` | `github.com/fallofpheonix/PhoenixCore/bus/normalizer` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/policy_test.go` | `github.com/fallofpheonix/PhoenixCore/containment` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/priority_test.go` | `github.com/fallofpheonix/PhoenixCore/bus` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/process_proof_test.go` | `github.com/fallofpheonix/PhoenixCore/containment` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/process_test.go` | `github.com/fallofpheonix/PhoenixCore/containment` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/pruning_test.go` | `github.com/fallofpheonix/PhoenixCore/common/resource` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/recovery_hardening_test.go` | `github.com/fallofpheonix/PhoenixCore/state` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/replay_test.go` | `github.com/fallofpheonix/PhoenixCore/containment/rollback` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/rollback_determinism_test.go` | `github.com/fallofpheonix/PhoenixCore/containment`, `github.com/fallofpheonix/PhoenixCore/containment/file`, `github.com/fallofpheonix/PhoenixCore/containment/network`, `github.com/fallofpheonix/PhoenixCore/containment/rollback` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/rollback_integration_test.go` | `github.com/fallofpheonix/PhoenixCore/containment`, `github.com/fallofpheonix/PhoenixCore/containment/file`, `github.com/fallofpheonix/PhoenixCore/containment/network`, `github.com/fallofpheonix/PhoenixCore/containment/rollback` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/rollback_proof_test.go` | `github.com/fallofpheonix/PhoenixCore/containment`, `github.com/fallofpheonix/PhoenixCore/containment/file`, `github.com/fallofpheonix/PhoenixCore/containment/network`, `github.com/fallofpheonix/PhoenixCore/containment/rollback` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/snapshot_test.go` | `github.com/fallofpheonix/PhoenixCore/truth` |
| `core/Phoenix.Nucleus/PhoenixValidation/unit/state_compat.go` | `github.com/fallofpheonix/PhoenixCore/state` |
| `core/Phoenix.Nucleus/authority/audit.go` | `github.com/fallofpheonix/Phoenix.Nucleus/ledger` |
| `core/Phoenix.Nucleus/authority/audit_test.go` | `github.com/fallofpheonix/Phoenix.Nucleus/ledger` |
| `core/Phoenix.Nucleus/authority/conservation.go` | `github.com/fallofpheonix/Phoenix.Nucleus/ledger` |
| `core/Phoenix.Nucleus/authority/policy.go` | `github.com/fallofpheonix/Phoenix.Nucleus/ledger` |
| `core/Phoenix.Nucleus/capability/token.go` | `github.com/fallofpheonix/Phoenix.Nucleus/ledger` |
| `core/Phoenix.Nucleus/recovery/boot.go` | `github.com/fallofpheonix/Phoenix.Nucleus/ledger` |
