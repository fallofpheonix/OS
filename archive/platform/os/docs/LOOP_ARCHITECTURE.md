---\nStatus: Partial\nImplementation: 60%\nConfidence: Tested\n---\n# Loop Architecture

## Overview

The PhoenixOS continuous execution loop orchestrates automated debugging, validation, repair, and improvement cycles, with human approval gates. It operates under strict rules to prevent direct AI-to-production patches.

---

## Core Loop Components

### phoenixmind-loop/

This top-level module houses the orchestration components for the continuous loop.

*   `watcher/`: Monitors filesystem changes and triggers the loop.
*   `builder/`: Handles incremental builds and captures failures.
*   `tester/`: Executes comprehensive test suites.
*   `observer/`: Provides continuous runtime monitoring and drift detection.
*   `replay/`: Ensures determinism and provides replay validation.
*   `truth/`: Establishes ground truth from logs and evidence.
*   `repair/`: Generates repair proposals within cognition engine.
*   `sandbox/`: Executes repair proposals in an isolated environment.
*   `eval/`: Evaluates system performance and repair effectiveness.
*   `gate/`: Manages the human approval process.

---

## Continuous Execution Flow

```
filesystem change
        ↓
watcher (phoenixmind-loop/watcher)
        ↓
incremental build (phoenixmind-loop/builder)
        ↓
unit tests (phoenixmind-loop/tester)
        ↓
integration tests (phoenixmind-loop/tester)
        ↓
runtime execution
        ↓
log capture (phoenixmind-observability)
        ↓
replay verify (phoenix_os/replay)
        ↓
truth validate (phoenixmind-validator)
        ↓
repair proposal (experimental/cognition_engine/repair)
        ↓
sandbox run (phoenixmind-sandbox)
        ↓
evaluation (phoenixmind-evals)
        ↓
approval (human gate)
        ↓
merge
```

---

## Stage Details

### Stage 1: Watch System (phoenixmind-loop/watcher)
*   Monitors: `phoenix_os/`, `validator/`, `runtime/`, `security/`, `observability/`.
*   Events: `file change`, `new test`, `build fail`, `panic`, `coverage drop`, `drift increase`.

### Stage 2: Build Runner (phoenixmind-loop/builder)
*   Flow: `repo change` → `affected modules` → `build` → `capture failure`.
*   Output Example: `{"module":"validator", "build":"FAIL", "error":"missing import"}`.

### Stage 3: Test Engine (phoenixmind-loop/tester)
*   Runs: `go test`, `fuzz`, `race`, `chaos`, `replay`.
*   Collects: `pass rate`, `coverage`, `panic count`, `latency`, `determinism`.

### Stage 4: Observability (phoenixmind-observability)
*   Flow: `runtime` → `jsonl logs` → `history` → `baseline` → `drift` → `trend`.
*   Writes: `runtime_audit.jsonl`, `MODULE_STATUS.json`, `DRIFT_HISTORY.json`, `OBS-*.json`.

### Stage 5: Replay Validation (phoenix_os/replay)
*   Flow: `run A` → `hash` → `run B` → `hash` → `compare`.
*   Results: `VALIDATED` (if same hash), `ESCALATED` (if different hash).

### Stage 6: Truth Engine (phoenixmind-validator)
*   Flow: `logs` → `evidence` → `confidence` → `resolver` → `truth`.
*   Outputs: `UNKNOWN`, `OBSERVED`, `VALIDATED`, `WARNING`, `ESCALATED`, `BLOCKED`.

### Stage 7: Repair Engine (experimental/cognition_engine)
*   Flow: `failure` → `AST parse` → `semantic diff` → `repair idea` → `risk score` → `simulation`.
*   Constraint: Never `repair` → `runtime write`. Only `repair` → `sandbox`.

### Stage 8: Sandbox Execution (phoenixmind-sandbox)
*   Flow: `proposal` → `sandbox run` → `tests` → `replay` → `truth` → `score`.
*   Pass Conditions: `coverage up`, `panic down`, `hash same`, `latency stable`.
*   Reject Conditions: `regression`, `hash mismatch`, `runtime mutation`, `security violation`.

### Stage 9: Evaluation (phoenixmind-evals)
*   Metrics: `build success %`, `test %`, `replay %`, `truth %`, `drift %`, `security %`, `repair success %`.

### Stage 10: Human Gate
*   Merge if: `tests PASS`, `replay PASS`, `truth PASS`, `security PASS`, `sandbox PASS`, `evaluation PASS`.
*   Then: `approve` → `merge`.

---

## Rules for Continuous Operation

*   Do **NOT** run: `code` → `LLM` → `patch` → `production`.
*   AI is advisory; never directly controls kernel or actuation FSM.
*   Cognition components do not directly modify runtime.

---

## Recommended Implementation Order

1.  watcher
2.  builder
3.  tester
4.  observability
5.  replay
6.  truth
7.  sandbox
8.  eval
9.  repair
10. training (later)
