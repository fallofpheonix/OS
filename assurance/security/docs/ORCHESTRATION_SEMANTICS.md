---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Orchestration Semantics: Determinism Through Composition

**Purpose**: Define the explicit runtime semantics that govern multi-step control-plane. These semantics ensure that trust (canonicalization, containment, resource governance) survives composition.

**Principle**: Trust preservation under composition is the actual control-plane problem. If determinism, auditability, and explicitness survive multi-operation execution, the system evolves correctly.

---

## Core Orchestration Invariant

**Deterministic Replay**

```
same control-plane + same inputs → same execution order + same traces + same failures + same final state
```

This becomes the defining control-plane invariant. Any control-plane that cannot be deterministically replayed is not trustworthy.

---

## Failure Propagation Semantics

Orchestration halts deterministically on failure. No continuation, no branching, no conditional logic.

### Failure Halt Matrix

| Step | Failure | Behavior | Trace | State |
|------|---------|----------|-------|-------|
| Read Config | File not found | HALT (don't proceed to validate) | traces read failure | final_state = FAILED |
| Read Config | Oversized file | HALT (don't proceed to validate) | traces read failure | final_state = FAILED |
| Read Config | Binary content | HALT (don't proceed to validate) | traces read failure | final_state = FAILED |
| Validate Structure | Schema invalid | HALT (don't proceed to execute) | traces read + validate failure | final_state = FAILED |
| Validate Structure | Constraint violated | HALT (don't proceed to execute) | traces read + validate failure | final_state = FAILED |
| Execute Command | Command not found | HALT (don't proceed to capture) | traces read + validate + execute failure | final_state = FAILED |
| Execute Command | Timeout | HALT (don't proceed to capture) | traces read + validate + execute timeout | final_state = FAILED |
| Capture Result | Capture fails (impossible if execute succeeded) | N/A (execute guarantees result) | N/A | N/A |

### Propagation Rule

**Rule: Fail-Fast Deterministic**

If ANY step fails:
1. Record the failure trace
2. Do NOT execute subsequent steps
3. Return `OrchestrationResult(final_state=FAILED, steps=[...all traces so far...])`
4. Same failure always produces identical result (deterministic replay invariant)

### What This Prevents

- ❌ No conditional branching (if step 2 fails, try alternate logic)
- ❌ No automatic retries (exponential backoff, circuit breakers)
- ❌ No partial recovery (capture some results, continue)
- ❌ No compensation logic (undo step 1 if step 3 fails)
- ❌ No event-driven recovery (emit failure, wait for external signal)

**Why**: Those features destroy determinism. Deterministic replay requires: same input → same path every time.

---

## State Authoritativeness

Which runtime traces are canonical? Which state is considered the "truth"?

### Authoritative Sources

1. **Filesystem reads** - File content is authoritative
2. **Validation results** - Schema check outcome is authoritative
3. **Command execution** - Exit code and stdout/stderr are authoritative
4. **Traces** - Immutable records are authoritative

### Non-Authoritative Sources

- ❌ Application interpretation of results (subject to logic bugs)
- ❌ Caching or memoization (non-deterministic across runs)
- ❌ External signals (control-plane should not listen for events)
- ❌ Configuration or registry state (dynamic state creates ordering dependence)

### Rule: Always Trust Runtime Traces

Orchestration reasoning must be based on:
- Read result (success/failure)
- Validate result (success/failure)
- Execute result (exit code, stdout, stderr)
- Aggregate trace (composition of runtime traces)

Never based on:
- Application state
- Caching decisions
- Event system signals
- Configuration lookups

---

## Ordering Guarantees

Orchestration executes in fixed sequence. No dynamic reordering, no parallel execution, no conditional skipping.

### Fixed Sequence

```
Step 1: Read Configuration
         ↓ (deterministic, not conditional)
Step 2: Validate Structure
         ↓ (deterministic, not conditional)
Step 3: Execute Command
         ↓ (deterministic, not conditional)
Step 4: Capture Result
```

### Execution Properties

- **Deterministic**: Same control-plane always executes in same order
- **Sequential**: No step executes until previous step completes
- **Atomic at control-plane boundary**: All steps or none (due to fail-fast)
- **Order-preserving traces**: Aggregate trace reflects execution order

### What This Prevents

- ❌ No dynamic scheduling (choose next step based on previous result)
- ❌ No parallel execution (step 1 and step 2 simultaneously)
- ❌ No conditional skipping (if step 1 succeeds, skip step 2)
- ❌ No alternative paths (if step 1 fails, try step 1b)

**Why**: Dynamic ordering destroys deterministic replay. Fixed order + fail-fast = deterministic.

---

## Trace Aggregation Rules

Orchestration aggregates runtime traces. Aggregation is composition, not mutation.

### Aggregation Properties

1. **Immutability**: Aggregated trace cannot be modified after creation
2. **Composition**: Aggregate trace contains references to step traces (not reconstruction)
3. **Completeness**: Aggregate trace captures all steps executed (including failures)
4. **Ordering**: Aggregate trace preserves step execution order
5. **Determinism**: Same control-plane produces identical aggregate trace

### Aggregation Structure

```python
@dataclass(frozen=True)
class OrchestrationStep:
    step_name: str  # "read", "validate", "execute", "capture"
    result: FileOperationResult | ExecutionResult  # runtime result
    trace: RuntimeTrace  # immutable runtime trace

@dataclass(frozen=True)
class OrchestrationResult:
    steps: tuple[OrchestrationStep, ...]  # step-by-step audit trail
    final_state: OrchestrationState  # SUCCESS, FAILED, TIMEOUT, UNKNOWN
    control-plane_trace: RuntimeTrace  # control-plane-level trace
```

### What Traces Record

Per step:
- Step name
- Result (success, error message, content)
- Duration
- Timestamp
- Trace ID (unique to that step execution)

Aggregate:
- All steps (even failed ones)
- Execution order (preserves sequence)
- Final state (how did control-plane end?)
- Orchestration duration (total time for all steps)

### What Traces Never Do

- ❌ Merge or collapse data (each step remains auditable)
- ❌ Summarize or abstract (exact step results preserved)
- ❌ Include application state (runtime-level facts only)
- ❌ Support mutation (immutable after creation)

---

## Validation Semantics

Validation is observational only. Never mutating.

### Observational Validation

Good:
- Schema validation (does structure match expected schema?)
- Constraint checking (are required fields present?)
- Type checking (do fields have expected types?)
- Bounds checking (are values within acceptable ranges?)

Bad:
- Normalization (transform values to canonical form)
- Enrichment (add computed fields)
- Coercion (convert types automatically)
- Correction (fix invalid values)

### Why Observational Only

Mutation semantics create hidden control-plane state:
- If validation corrects a value, what was the original? (loss of auditability)
- If validation enriches data, what is the authoritative version? (state confusion)
- If validation coerces types, what was actually read? (loss of determinism)

### Validation Result

Validation returns:
- `success: bool` (does structure match schema?)
- `error: str | None` (what constraint failed?)

Validation never returns:
- Corrected data
- Enriched data
- Transformed data
- Inferred data

---

## Rollback Semantics

**Phase 1 Orchestration: No rollback support.**

### Why Not Rollback

Rollback introduces compensation logic:
- If step 3 fails, undo step 2
- If step 2 fails, undo step 1
- If undo fails, what now?

This complexity:
- Destroys deterministic replay (undo paths are conditional)
- Creates state recovery problems (what state is authoritative after partial undo?)
- Introduces compensation semantics (operation X is undone by operation Y)
- Requires transaction-like guarantees (consistency across steps)

### Current Correct Semantics

Fail-fast deterministic execution:
- If any step fails, halt immediately
- Do NOT undo previous steps
- Return complete control-plane result (all steps, including failures)
- Operator can inspect results and decide manual recovery

This preserves:
- Deterministic replay
- Auditability (all steps remain visible)
- Simplicity (no compensation logic)
- Correctness (failed control-planes are idempotent to re-run)

### Future Consideration (Not Phase 1)

Once control-plane determinism is proven stable, future phases may introduce:
- Explicit idempotence semantics (same control-plane can be safely re-run)
- State snapshot/restore (operator can manually save/restore state)
- Compensating operations (explicit, not implicit)

For now: fail-fast, halt-on-error, preserve all traces.

---

## No Runtime Boundary Bypass

Orchestration never:
- Touches filesystem directly
- Executes subprocesses directly
- Mutates traces directly
- Modifies runtime state

Orchestration only:
- Calls FilesystemManager API (read, list, exists)
- Calls ShellExecutor API (execute with timeout)
- Coordinates results (composition, not integration)
- Aggregates traces (immutable composition only)

### Why

Runtime boundaries exist to enforce:
- Path containment (resolver guarantees)
- Resource limits (policy guarantees)
- Exception safety (domain exceptions only)
- Trace immutability (frozen dataclass)

If control-plane bypasses these boundaries:
- Containment escapes become possible
- Resource exhaustion becomes possible
- OS exceptions can leak
- Traces become mutable

Orchestration is only trustworthy if it inherits runtime trustworthiness.

---

## Orchestration Result States

```python
@enum
class OrchestrationState:
    SUCCESS = "success"  # all steps executed successfully
    FAILED = "failed"    # a step failed, control-plane halted
    TIMEOUT = "timeout"  # a step timed out, control-plane halted
    UNKNOWN = "unknown"  # control-plane could not complete (internal error)
```

### State Transitions

```
START
  ↓
Read Configuration
  ├─ success → Validate Structure
  └─ failure → FAILED (don't proceed)
      ↓
  Validate Structure
    ├─ success → Execute Command
    └─ failure → FAILED (don't proceed)
        ↓
    Execute Command
      ├─ success → Capture Result
      ├─ failure → FAILED (don't proceed)
      └─ timeout → TIMEOUT (don't proceed)
          ↓
      Capture Result
        ├─ success → SUCCESS
        └─ failure → UNKNOWN (should not happen)
```

### What Results Mean

- **SUCCESS**: Read succeeded, validation passed, command executed, result captured
- **FAILED**: A step failed (read, validate, or execute); subsequent steps not executed
- **TIMEOUT**: Command timed out; subsequent steps (capture) not executed
- **UNKNOWN**: Orchestration could not complete (internal error in control-plane itself)

---

## Determinism Properties

### Deterministic Replay Invariant

Given identical inputs and control-plane definition:

```
determinism_check(orch1_result, orch2_result):
  assert orch1_result.steps == orch2_result.steps
  assert orch1_result.final_state == orch2_result.final_state
  assert orch1_result.traces are in same order
  assert orch1_result.traces have same error_type (if failed)
```

### Non-Determinism Red Flags

If ANY of these vary across runs:
- ❌ Step execution order changes
- ❌ Same input produces different error at same step
- ❌ Traces are in different order
- ❌ Final state differs

→ **Orchestration is not deterministic. Investigation required.**

### Guaranteeing Determinism

1. **Fixed execution order** (no dynamic reordering)
2. **Fail-fast halt** (no conditional continuation)
3. **Explicit state** (no hidden configuration)
4. **Immutable traces** (no post-execution mutation)
5. **No external signals** (no event-driven logic)

If all five properties hold, determinism is preserved.

---

## Summary: Orchestration Trust Model

Orchestration is trustworthy if:

1. **It preserves runtime trust** - Never bypasses resolver/policy/manager boundaries
2. **It remains deterministic** - Same input always produces identical execution
3. **It stays auditable** - Every step is individually traceable
4. **It fails explicitly** - Failures are clear, not hidden
5. **It avoids complexity** - No branching, no retries, no compensation

Trust + Determinism + Auditability = Safe Composition

If any one property breaks, the system begins collapsing into entropy.

---

## Appendix: Forbidden Orchestration Patterns

Do NOT implement:

- Workflow engines (DAG execution, dynamic scheduling)
- Plugin architectures (runtime operation loading)
- Event buses (event-driven control-plane)
- Middleware stacks (request/response pipelines)
- Generic operation registries (dynamic dispatch)
- Retryable control-plane (automatic failure recovery)
- Conditional branching (dynamic path selection)
- Transactional semantics (rollback/undo logic)
- Async execution (parallel step execution)
- State machines (implicit state transitions)

Reason: All destroy determinism, auditability, or explicitness.

---

## Next: Composite Operation Specification

See `COMPOSITE_OPERATION_SPEC.md` for the concrete first control-plane slice.
