---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Composite Operation Specification: Read-Validate-Execute-Capture

**Purpose**: Define the first control-plane slice that proves trust survives composition.

**Scope**: Single deterministic execution chain. No branching, no dynamic dispatch, no retry logic.

---

## Operation Overview

```
Step 1: Read Configuration File
         ↓ (deterministic, not conditional)
Step 2: Validate Configuration Structure
         ↓ (deterministic, not conditional)
Step 3: Execute Deployment Command
         ↓ (deterministic, not conditional)
Step 4: Capture Execution Result
         ↓
Aggregate Traces & Return Orchestration Result
```

---

## Step 1: Read Configuration File

**Responsibility**: Read a YAML or JSON configuration file from the workspace.

**Input**:
- `config_path: str` - Relative path within workspace (e.g., `config/deploy.yaml`)

**Process**:
1. Call `FilesystemManager.read_file(config_path)`
2. Receive `FileOperationResult` with content or error

**Output**:
- `OrchestrationStep(step_name="read", result=result, trace=result.trace)`

**Success Condition**:
- `result.success == True`
- `result.content` is non-empty string

**Failure Conditions**:
- File not found → `result.success == False`, error contains "does not exist"
- File outside workspace → `result.success == False`, error contains "boundary violation"
- File too large (>1 MB) → `result.success == False`, error contains "exceeds"
- Binary file → `result.success == False`, error contains "binary"
- Invalid UTF-8 → `result.success == False`, error contains "unsupported encoding"

**If Fails**: HALT. Do NOT proceed to Step 2. Return control-plane result with `final_state=FAILED`.

**Trace Captured**: `FileOperationResult.trace` (immutable, contains operation details)

---

## Step 2: Validate Configuration Structure

**Responsibility**: Validate that the read configuration has required structure and content.

**Input**:
- `config_content: str` - Raw string from Step 1
- `schema: dict` - Expected structure (for Step 1 control-plane: simple schema check)

**Process**:
1. Parse content as JSON or YAML (Python's JSON library for JSON, standard dict for simple structures)
2. Check required keys exist
3. Check required values are non-empty
4. Return validation result

**Output**:
- `OrchestrationStep(step_name="validate", result=validation_result, trace=validation_trace)`

**Success Condition**:
- Parsed structure matches expected schema
- All required keys present
- All required values non-empty

**Failure Conditions**:
- JSON/YAML parse error → validation fails, error describes parse failure
- Required key missing → validation fails, error names the key
- Required value empty → validation fails, error names the field

**Important**: Validation is OBSERVATIONAL only.
- ❌ Do NOT transform the configuration
- ❌ Do NOT normalize values
- ❌ Do NOT enrich with defaults
- ✅ Only check structure matches expected schema

**If Fails**: HALT. Do NOT proceed to Step 3. Return control-plane result with `final_state=FAILED`.

**Trace Captured**: Custom `ValidationTrace` (simple dataclass with step_name="validate", success, error_type, duration)

### Validation Schema for Phase 1

Simple structure:
```python
{
    "command": str,              # required, non-empty
    "args": list[str],           # required, may be empty
    "description": str           # optional
}
```

Example valid configuration:
```json
{
    "command": "deploy.sh",
    "args": ["--env", "production"],
    "description": "Deploy to production"
}
```

Example invalid configuration:
```json
{
    "command": "",  # FAIL: empty required field
    "args": ["--env"]
}
```

---

## Step 3: Execute Deployment Command

**Responsibility**: Execute the command specified in the validated configuration.

**Input**:
- `config: dict` - Validated configuration from Step 2
- `timeout_seconds: float` - Command timeout (default 30.0)

**Process**:
1. Extract `command` and `args` from validated config
2. Call `ShellExecutor.execute(command, args, timeout_seconds)`
3. Receive `ExecutionResult` with exit code, stdout, stderr, or error

**Output**:
- `OrchestrationStep(step_name="execute", result=result, trace=result.trace)`

**Success Condition**:
- Command completed without timeout
- Exit code can be any value (0 or non-zero)
- `result.success == True` means execution completed (not necessarily exit code 0)

**Failure Conditions**:
- Command not found → `result.success == False`, error contains "not found"
- Command timed out → `result.success == False`, error contains "timeout"
- Shell execution error → `result.success == False`, error describes the error

**Important Distinction**:
- `result.success == True` means: "command executed and produced output"
- `exit_code != 0` means: "command executed but returned non-zero exit code"
- These are NOT the same as control-plane failure

**If Fails (timeout or not found)**: HALT. Do NOT proceed to Step 4. Return control-plane result with `final_state=FAILED` or `final_state=TIMEOUT`.

**If Succeeds (completed)**: Continue to Step 4 even if `exit_code != 0`.

**Trace Captured**: `ExecutionResult.trace` (immutable, contains command, exit code, duration)

---

## Step 4: Capture Execution Result

**Responsibility**: Capture and structure the execution result for the control-plane result.

**Input**:
- `execution_result: ExecutionResult` - Result from Step 3

**Process**:
1. Extract stdout, stderr, exit_code from execution result
2. Create `CaptureResult` dataclass with this information
3. Return step result

**Output**:
- `OrchestrationStep(step_name="capture", result=capture_result, trace=execution_result.trace)`

**Success Condition**:
- `result.success == True` (execution completed)

**Failure Conditions**:
- Step 3 failed → capture is not executed (early halt due to Step 3 failure)

**Important**: Capture does NOT interpret the results.
- ❌ Do NOT check if exit code indicates success
- ❌ Do NOT parse command output
- ❌ Do NOT decide if deployment "worked"
- ✅ Only capture stdout, stderr, exit code as facts

**Trace Captured**: Reuse `ExecutionResult.trace` from Step 3

---

## Orchestration Result Model

```python
@dataclass(frozen=True)
class OrchestrationStep:
    step_name: str  # "read", "validate", "execute", "capture"
    result: object  # FileOperationResult, ValidationResult, ExecutionResult, CaptureResult
    trace: RuntimeTrace  # immutable trace

@dataclass(frozen=True)
class OrchestrationResult:
    steps: tuple[OrchestrationStep, ...]  # all steps executed (even failures)
    final_state: OrchestrationState  # SUCCESS, FAILED, TIMEOUT, UNKNOWN
    control-plane_trace: RuntimeTrace  # aggregated trace
```

### Final State Determination

```python
def determine_final_state(steps: List[OrchestrationStep]) -> OrchestrationState:
    for step in steps:
        if step.step_name == "read" and not step.result.success:
            return OrchestrationState.FAILED  # read failed
        if step.step_name == "validate" and not step.result.success:
            return OrchestrationState.FAILED  # validation failed
        if step.step_name == "execute":
            if "timeout" in step.result.error.lower():
                return OrchestrationState.TIMEOUT  # timeout
            if not step.result.success:
                return OrchestrationState.FAILED  # execution failed
        if step.step_name == "capture":
            if step.result.success:
                return OrchestrationState.SUCCESS  # all steps succeeded
    
    return OrchestrationState.UNKNOWN  # should not reach here
```

---

## Execution Flow Examples

### Example 1: Success Path

```
Input: config_path="config/deploy.yaml"

Step 1: Read
  result.success = True
  result.content = "command: deploy.sh\nargs: [--prod]"
  ✓ Proceed to Step 2

Step 2: Validate
  parsed = {"command": "deploy.sh", "args": ["--prod"]}
  schema check = True (all required fields present)
  ✓ Proceed to Step 3

Step 3: Execute
  command = "deploy.sh"
  args = ["--prod"]
  exit_code = 0
  result.success = True
  ✓ Proceed to Step 4

Step 4: Capture
  stdout = "Deployment successful"
  stderr = ""
  exit_code = 0
  result.success = True
  ✓ HALT

Final State: SUCCESS
Traces: [read_trace, validate_trace, execute_trace, execute_trace (reused)]
```

### Example 2: Read Fails

```
Input: config_path="config/missing.yaml"

Step 1: Read
  result.success = False
  result.error = "file does not exist: config/missing.yaml"
  ✗ HALT (don't proceed to Step 2)

Step 2-4: NOT EXECUTED

Final State: FAILED
Traces: [read_trace]
```

### Example 3: Validate Fails

```
Input: config_path="config/invalid.yaml"

Step 1: Read
  result.success = True
  result.content = "command: ''"  # empty command
  ✓ Proceed to Step 2

Step 2: Validate
  parsed = {"command": ""}
  schema check = False (empty required field)
  result.success = False
  result.error = "required field 'command' is empty"
  ✗ HALT (don't proceed to Step 3)

Step 3-4: NOT EXECUTED

Final State: FAILED
Traces: [read_trace, validate_trace]
```

### Example 4: Execute Times Out

```
Input: config_path="config/deploy.yaml"

Step 1: Read
  result.success = True
  ✓ Proceed to Step 2

Step 2: Validate
  result.success = True
  ✓ Proceed to Step 3

Step 3: Execute
  command = "sleep 100"  # will timeout
  timeout_seconds = 5
  result.success = False
  result.error = "execution timed out after 5 seconds"
  ✗ HALT (don't proceed to Step 4)

Step 4: NOT EXECUTED

Final State: TIMEOUT
Traces: [read_trace, validate_trace, execute_trace]
```

---

## Determinism Contract

This control-plane satisfies deterministic replay:

```
control-plane_result_1 = execute_composite_operation(config_path, timeout)
control-plane_result_2 = execute_composite_operation(config_path, timeout)

assert control-plane_result_1.steps == control-plane_result_2.steps
assert control-plane_result_1.final_state == control-plane_result_2.final_state
assert len(control-plane_result_1.steps) == len(control-plane_result_2.steps)
assert all(s1.trace.error_type == s2.trace.error_type for s1, s2 in zip(...))
```

Same input, same runtime state → identical control-plane result.

---

## Failure Propagation Contract

```
if step N fails:
    do NOT execute steps N+1, N+2, ...
    halt immediately
    return control-plane_result with:
        - all steps N and before
        - final_state = FAILED or TIMEOUT
        - no subsequent steps
```

Every failure is deterministic. Same failure always halts at same step.

---

## Trace Aggregation Contract

```
control-plane_result.steps contains:
    - step_name (identifier)
    - result (operation result)
    - trace (immutable runtime trace)

All traces are immutable.
All traces are in execution order.
All traces reference the exact runtime operation.
No traces are merged, summarized, or reconstructed.
```

Complete audit trail. Step-level granularity. Immutable composition.

---

## Step-Identity Preservation Contract

```
control-plane_result.steps = [
    OrchestrationStep(step_name="read", ...),
    OrchestrationStep(step_name="validate", ...),
    OrchestrationStep(step_name="execute", ...),
]

Each step is individually auditable.
Each step's trace is independent.
Each step's result is preserved exactly.
No aggregation collapses or summarizes step information.
```

Every step remains visible. Partial failure traces are complete.

---

## Next: Implementation

The implementation must:

1. **Call runtime APIs only** (FilesystemManager, ShellExecutor)
2. **Never bypass runtime boundaries**
3. **Preserve step identity** (each step is auditable)
4. **Aggregate immutably** (traces are frozen)
5. **Fail-fast deterministically** (same failure, same halt point)
6. **Support deterministic replay** (same input, identical result)

See `runtime/control-plane/composite_operation.py` for implementation.
