# PHOENIX_VALIDATION_ROOT_CAUSE_REPORT

Purpose: classify every failing record from the `PhoenixValidation` test run, identify root causes, and list the required fix and risk level. This is a diagnosis only — no patches applied.

Architecture framing: the failures below are contract drift, so they map to [FINAL_ARCHITECTURE.md](docs/architecture/FINAL_ARCHITECTURE.md) and the debt register rather than to isolated implementation bugs.

Source of evidence: `build_logs/module-tests/core_Phoenix.Nucleus_PhoenixValidation.log` and the validation package sources under `core/Phoenix.Nucleus/PhoenixValidation/`.

Summary of failing items found:

1) chaos/fuzz_test.go:35
- File: `core/Phoenix.Nucleus/PhoenixValidation/chaos/fuzz_test.go`
- Error (log):
  - `cannot use evt (variable of type *"github.com/fallofpheonix/PhoenixCore/proto/v1/event".EventEnvelope) as "github.com/fallofpheonix/PhoenixCore/event".Event value in argument to engine.Apply`
- Error Type: Compile-time type mismatch (Proto type vs internal domain type)
- Root Cause: `engine.Apply` expects `event.Event` (the internal struct in `PhoenixCore/event`), while tests construct and pass `proto/v1/event.EventEnvelope` (generated protobuf type). There is no adapter/conversion present.
- Required Fix: Introduce an adapter function to convert `proto/v1/event.EventEnvelope` → `event.Event` (or overload/extend `Engine.Apply` to accept the proto type), and update tests to use the canonical adapter. Alternatively, change test inputs to build `event.Event` values.
- Risk Level: Medium (affects replay ingestion correctness and test harness compatibility).

2) invariants/clock_test.go:30,41
- File: `core/Phoenix.Nucleus/PhoenixValidation/invariants/clock_test.go`
- Error (log):
  - same compile-time mismatch as in `chaos/fuzz_test.go` for `evt1` and `evt2` when calling `engine.Apply`.
- Error Type: Compile-time type mismatch
- Root Cause: same as #1; tests use `proto/v1/event.EventEnvelope` but `Engine.Apply` requires `event.Event`.
- Required Fix: same as #1. After conversion, re-run invariant tests to verify monotonic clock checks.
- Risk Level: Medium

3) replay/engine_test.go: lines referencing `Reconstruct` and `currentState`
- File: `core/Phoenix.Nucleus/PhoenixValidation/replay/engine_test.go`
- Errors (log):
  - `engine.Reconstruct undefined (type *Engine has no field or method Reconstruct)`
  - `engine.currentState undefined (type *Engine has no field or method currentState)`
- Error Type: API/ABI drift (test expects methods/fields not present in current Engine implementation)
- Root Cause: The `Engine` implementation currently provides `Replay(events []event.Event)` and exposes an exported `State` field on `Engine` (type `State`). Tests expect a `Reconstruct` method and an internal `currentState` structure (with `.Warden.CurrentState`). Either the engine was refactored (method renamed) or tests reference a previous internal shape.
- Required Fix: Pick one of:
  - Reintroduce compatibility shims: implement `Reconstruct(events []*proto.EventEnvelope)` that converts inputs and delegates to `Replay`, and expose an accessor `CurrentState()` returning the expected structure for tests; or
  - Update tests to use `Replay` and read from the exported `Engine.State` (and adapt assertions to the new shape).
- Risk Level: High (the replay engine is central to correctness proofs and deterministic reconstruction).

4) proofs/containment_proof_test.go: runtime invariant/assertion failure
- File: `core/Phoenix.Nucleus/PhoenixValidation/proofs/containment_proof_test.go`
- Failure (log):
  - `CONTAINMENT_PROOF_FAILED: PID 1234 was not isolated by the actuator`
  - Prior logs show Warden invariant violations and emergency halt: `[VIOLATION] [Warden Panic] Invariant Violation: insufficient evidence weight to reach WATCH (0.10 < 0.50)` and `[Warden EMERGENCY HALT] Quenching System: invariant violation...`
- Error Type: Runtime behavior / assertion failure (test expectation mismatch with runtime behavior)
- Root Cause: Implementation decision: when an invariant Verify fails inside `Warden.Actuate`, the implementation calls `emergencyHalt` which sets `StateCompromised` and returns without executing the physical actuation ladder. The test expects that the Warden will both enter `StateCompromised` and still call actuators (or that actuation occurs as part of emergency handling). This mismatch (test expecting actuator invocation on invariant violation) causes the assertion that `act.IsolatedPID != 1234`.
- Required Fix: Decide desired behavior and make tests or runtime consistent:
  - If the desired behavior is to still notify/trigger actuators on invariant failure, change `Actuate` to perform actuation even after emergencyHalt (careful: could be dangerous). Or,
  - If the desired behavior is to stop actuation on invariant failure (current implementation), update the test to assert no actuator calls and only verify `StateCompromised` and appropriate logs.
- Risk Level: High (involves safety-critical behavior and physical actuation semantics).

5) proofs/proofs and replay/proofs runtime failures
- Files: `core/Phoenix.Nucleus/PhoenixValidation/proofs/*` and `core/Phoenix.Nucleus/PhoenixValidation/proofs` (see test log lines referencing `profs` fail)
- Failures: multiple `FAIL` entries in proofs and replay subpackages; some tests call engine methods that don't exist or rely on actuators/internal fields.
- Error Type: Mix of compile-time API drift and runtime invariant failures (see items 1–4 above).
- Root Cause: Consolidation of above issues — tests rely on older APIs and on semantics that current runtime doesn't match.
- Required Fix: Multi-step: (a) reconcile proto vs domain types; (b) reconcile Engine API; (c) reconcile Warden actuation semantics in test harness.
- Risk Level: High

Dependency mapping around `PhoenixValidation` (owners and responsibility)

- `EventEnvelope` owner: `core/Phoenix.Nucleus/PhoenixCore/proto/v1/event` (generated protobuf). This is the canonical envelope used across system contracts (see `proto/v1/event/event.pb.go` and `proto/v1/event.proto`).
- `Event` owner: `core/Phoenix.Nucleus/PhoenixCore/event` (internal domain struct defined in `event/schema.go`). Tests and runtime must agree on conversions between these representations.
- `Replay Engine` owner: `core/Phoenix.Nucleus/PhoenixValidation/replay` (source: `replay/engine.go`). This package implements `Engine` with methods `Replay([]event.Event)`, `Apply(event.Event)`, `CalculateStateHash()`.
- `Reconstruction Logic` expected by tests: tests reference `Engine.Reconstruct([]*proto.EventEnvelope)` and `engine.currentState` — these are expected interfaces/fields from a prior design. Current owner remains `PhoenixValidation/replay`, but the API has changed.

Immediate non-invasive recommendations (no code changes yet):
1. Add a *diagnostic* adapter test that attempts to convert a minimal `proto.EventEnvelope` to `event.Event` using the canonical mappings (EventId→EventID, ReplaySequence→LogicalTime, Payload mapping). This will validate conversion feasibility before larger changes.
2. Decide a canonical inbound event type for the Replay Engine: either the proto envelope (public contract) or the internal `event.Event`. Prefer protobuf as external contract and write thin converters into `replay` package.
3. Update test expectations in `replay/engine_test.go` and `proofs/containment_proof_test.go` to align with current runtime semantics OR update runtime to match documented semantics (team decision required).

Appendix: exact failing lines (from test log)
- `chaos/fuzz_test.go:35:20`
- `invariants/clock_test.go:30:25`
- `invariants/clock_test.go:41:22`
- `replay/engine_test.go:35:27` and `replay/engine_test.go:41:12` etc.
- `proofs/containment_proof_test.go:83`

Generated: /PHOENIX_VALIDATION_ROOT_CAUSE_REPORT.md
