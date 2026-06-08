# Failure Report — PhoenixGuard & PhoenixValidation

Summary: after a successful build and per-module test run, two modules show failing test clusters that block quality verification: `PhoenixGuard` and `PhoenixValidation`. This report classifies failures, lists evidence, estimates fix effort, and recommends next actions. Do not run coverage until these are triaged.

Architecture framing: the core issue is contract ownership drift, now documented in [docs/architecture/FINAL_ARCHITECTURE.md](docs/architecture/FINAL_ARCHITECTURE.md) and [docs/architecture/ARCHITECTURE_DEBT.md](docs/architecture/ARCHITECTURE_DEBT.md).

---

## PhoenixGuard
- Failure Type: Compile-time test build error (tests fail to compile).
- Evidence:
  - `./warden_compat_test.go:131:21: cannot use a1 (variable of type *stubActuator) as Actuator value in argument to w.RegisterActuator: *stubActuator does not implement Actuator (missing method Kill)`
- Affected Packages: `github.com/fallofpheonix/PhoenixGuard` (tests), specifically `warden_compat_test.go` and `RegisterActuator` usage.
- Probable Root Cause: API/interface drift between production `Actuator` interface and test `stubActuator` (missing `Kill` method or signature mismatch). Could be a recent refactor that added/renamed interface method(s).
- Estimated Fix Complexity: Low → Medium (small code change to update the test stub or adjust interface usage). ~1–2 developer-days to confirm and patch tests.
- Risk If Unfixed: Medium — `PhoenixGuard` provides containment/actuation; failing tests reduce confidence in system safety checks.
- Recommended Action:
  1. Open `core/Phoenix.Nucleus/PhoenixGuard/warden_compat_test.go` and inspect `stubActuator` and `Actuator` interface definition in production code (`actuators` package).
  2. If `Actuator` recently added `Kill` (or similar), implement the method on `stubActuator` in the test or change test to use a compatible implementation.
  3. Re-run `go test ./core/Phoenix.Nucleus/PhoenixGuard -run TestName -v` to validate.

---

## PhoenixValidation
- Failure Types: Mix of compile-time API/type mismatches and runtime assertion failures during tests.
- Evidence (excerpts):
  - `chaos/fuzz_test.go:35:20: cannot use evt (variable of type *"github.com/fallofpheonix/PhoenixCore/proto/v1/event".EventEnvelope) as "github.com/fallofpheonix/PhoenixCore/event".Event value in argument to engine.Apply`
  - `replay/engine_test.go:35:27: engine.Reconstruct undefined (type *Engine has no field or method Reconstruct)`
  - Runtime assertion: `containment_proof_test.go:83: CONTAINMENT_PROOF_FAILED: PID 1234 was not isolated by the actuator` and Warden invariant violations logged during proofs.
- Affected Packages: `github.com/fallofpheonix/PhoenixValidation/{chaos,invariants,replay,proofs}` and associated test harnesses.
- Probable Root Cause(s):
  - API / type drift between generated proto types (`proto/v1/event.EventEnvelope`) and internal domain type (`PhoenixCore/event.Event`) — likely package path or type definition changes (proto regens or moved types).
  - Engine refactor: tests reference `Engine.Reconstruct` and `engine.currentState` which are no longer present (method/field renamed or encapsulated).
  - Some runtime test failures indicate behavior regressions (actuator didn't isolate PID) — could be test environment mismatch or actual logic failure.
- Estimated Fix Complexity: Medium → High (multi-package reconciliation). Rough estimate: 3–7 developer-days depending on whether proto types were intentionally changed (requires harmonizing types or adding adapter shims) and whether Engine API was refactored.
- Risk If Unfixed: High — validation and invariants failing reduce assurance of core system correctness; these tests exercise safety-critical behavior.
- Recommended Actions:
  1. Triage compile errors first:
     - Compare `PhoenixCore/proto/v1/event` types vs `PhoenixCore/event` interfaces; add thin adapters or update tests to use the correct types.
     - Search git history for recent changes to `Engine` API (Reconstruct/currentState) and update tests or reintroduce compatibility shims.
  2. Run failing tests in verbose mode and with focused logging to reproduce runtime invariant failures:
     - `go test ./core/Phoenix.Nucleus/PhoenixValidation/chaos -run TestName -v`
     - `go test ./core/Phoenix.Nucleus/PhoenixValidation/proofs -run TestContainmentProof -v`
  3. If failures are due to test harness expectations (mock actuators, external dependencies), update harness or add feature flags to run in shadow/simulated mode.
  4. After fixes, re-run full module tests for `PhoenixValidation`.

---

## Overall Assessment & Next Step (highest value)
- Priority: Investigate and fix the two failing modules before running coverage or security scans. These failures are narrow and high-impact; resolving them will unlock meaningful coverage and security analysis.
- Immediate next action I can perform for you now (pick one):
  - (A) Create a short patch to update `stubActuator` in `PhoenixGuard` tests if you want me to implement a minimal compatibility fix.
  - (B) Produce a focused triage for `PhoenixValidation` by searching for `Engine.Reconstruct`, `currentState`, and the `EventEnvelope`/`Event` type definitions and showing candidate diffs.

---

Generated: evidence-based failure summary from `build_logs/module-tests/`.
