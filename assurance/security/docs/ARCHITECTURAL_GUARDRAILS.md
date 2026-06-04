---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Architectural Guardrails: Protecting Against Framework Drift

## ⚠️ Most Important Warning

**The biggest risk to this system is NOT technical failure.**

**The biggest risk is psychological and architectural.**

Specifically: The temptation to generalize, abstract, and add flexibility.

Each small convenience abstraction seems harmless individually.

Collectively, they destroy:
- Cognitive compressibility
- Deterministic semantics
- Auditability
- Explicit behavior

**Most systems collapse at this exact point.**

**Core Defense**: Do not abstract until pressure justifies it (3+ types, measured patterns).

## Current System State

**Proven Property**: Deterministic trust semantics survive first-order composition

**Current Architecture Level**: Bounded execution proof (NOT a workflow platform)

**One Orchestration Chain**: Read → Validate → Execute → Capture

**Cognitive Compressibility**: Intact (still fits in one mental model)

---

## The Exact Point of Danger

Most systems collapse into framework engineering at this stage:

```
Stage 1 (Current): ✅ One control-plane chain
                     - Explicit, bounded, deterministic
                     - Fits in one mental model
                     - No abstraction pressure yet

Stage 2 (Dangerous): ❌ "Let's make it reusable"
                       - Step registry (plugins)
                       - Configurable ordering
                       - Middleware hooks
                       - Event notifications
                       - Orchestration templates

Result: Architecture entropy. Cognitive compressibility lost.
```

## Guardrails: What is Explicitly Forbidden Right Now

### Forbidden Abstraction Pattern #1: Pluggable Steps

❌ **FORBIDDEN** (Do not do this):
```python
class OrchestrationStep:
    """Generic step handler"""
    
class StepRegistry:
    """Register step handlers"""
    registry = {}
    
    @staticmethod
    def register(name, handler):
        registry[name] = handler
```

✅ **CORRECT** (What we have):
```python
def _step_read_configuration(self, config_path: str) -> OrchestrationStep:
    # Hard-coded, explicit, deterministic
    
def _step_validate_configuration(self, content: str) -> OrchestrationStep:
    # Hard-coded, explicit, deterministic
```

**Justification**: One control-plane chain. No registry needed.

---

### Forbidden Abstraction Pattern #2: Configurable Step Order

❌ **FORBIDDEN**:
```python
class Orchestration:
    def __init__(self, step_sequence: list[str]):
        self.steps = step_sequence  # Configurable order
```

✅ **CORRECT**:
```python
def execute(self, config_path: str) -> OrchestrationResult:
    # Fixed order
    step1 = self._step_read_configuration(config_path)
    step2 = self._step_validate_configuration(...)
    step3 = self._step_execute_command(...)
    step4 = self._step_capture_result(...)
```

**Justification**: Determinism requires fixed ordering. No dynamic reordering.

---

### Forbidden Abstraction Pattern #3: Step Middleware

❌ **FORBIDDEN**:
```python
class Orchestration:
    def __init__(self):
        self.before_hooks = []
        self.after_hooks = []
    
    def execute(self, config_path: str):
        for hook in self.before_hooks:
            hook(config_path)  # Hidden execution
        # ... main logic ...
        for hook in self.after_hooks:
            hook(result)  # Hidden mutation
```

✅ **CORRECT**:
```python
def execute(self, config_path: str) -> OrchestrationResult:
    # All steps explicit, no hidden hooks
    steps = []
    steps.append(self._step_read_configuration(config_path))
    steps.append(self._step_validate_configuration(...))
    # etc.
```

**Justification**: Explicitness enables auditability.

---

### Forbidden Abstraction Pattern #4: Conditional Execution

❌ **FORBIDDEN**:
```python
def execute(self, config_path: str) -> OrchestrationResult:
    if config.get("skip_validation"):
        # Conditional step skipping
        return self._execute_without_validation(config_path)
    else:
        # Normal path
        return self._execute_with_validation(config_path)
```

✅ **CORRECT**:
```python
def execute(self, config_path: str) -> OrchestrationResult:
    # Always read
    step1 = self._step_read_configuration(config_path)
    if not step1.result.success:
        return self._build_control-plane_result([step1], OrchestrationState.FAILED, ...)
    
    # Always validate
    step2 = self._step_validate_configuration(...)
    if not step2.result.success:
        return self._build_control-plane_result([step1, step2], OrchestrationState.FAILED, ...)
    
    # etc.
```

**Justification**: Determinism requires fixed sequence.

---

### Forbidden Abstraction Pattern #5: Orchestration State Persistence

❌ **FORBIDDEN**:
```python
class Orchestration:
    def __init__(self):
        self.execution_history = []  # Persistent state
        self.last_result = None       # Caching
        self.context = {}             # Execution context
```

✅ **CORRECT**:
```python
def execute(self, config_path: str) -> OrchestrationResult:
    # No self.* state modifications
    # All state in local variables and return value
    # No implicit context
```

**Justification**: Explicit state enables determinism.

---

### Forbidden Abstraction Pattern #6: Event-Driven Orchestration

❌ **FORBIDDEN**:
```python
class OrchestrationEventBus:
    def on_step_complete(self, listener):
        self.listeners.append(listener)
    
    def emit(self, event):
        for listener in self.listeners:
            listener(event)  # Hidden async-like behavior
```

✅ **CORRECT**:
```python
# No event bus, no listeners, no notifications
# Only explicit control-plane result returned
```

**Justification**: Events introduce non-deterministic ordering.

---

### Forbidden Abstraction Pattern #7: Retry Logic

❌ **FORBIDDEN**:
```python
def execute(self, config_path: str) -> OrchestrationResult:
    for attempt in range(self.max_retries):
        try:
            # Try to execute
            return self._do_execute(config_path)
        except Exception:
            if attempt < self.max_retries - 1:
                time.sleep(self.retry_delay)
```

✅ **CORRECT**:
```python
# Fail-fast semantics
# Error in step N → return failed result
# No retries, no recovery, no exponential backoff
```

**Justification**: Retries destroy determinism and observability.

---

### Forbidden Abstraction Pattern #8: Rollback/Compensation

❌ **FORBIDDEN**:
```python
class Orchestration:
    def execute(self, config_path: str) -> OrchestrationResult:
        try:
            result = self._do_execute(config_path)
            if result.final_state == OrchestrationState.FAILED:
                # Rollback executed steps
                self._rollback_step(3)  # Undo execute
                self._rollback_step(2)  # Undo validate
                self._rollback_step(1)  # Undo read
            return result
        except Exception:
            self._compensate()  # Compensating operations
```

✅ **CORRECT**:
```python
# No rollback, no compensation
# Fail-fast halt with deterministic error state
# Caller decides what to do with failed control-plane
```

**Justification**: Rollback introduces state recovery semantics (too early).

---

## When Abstraction Becomes Justified

**Current Condition**: 1 control-plane type = NO abstraction

**Future Condition** (only then):
- ✅ Implement 3+ distinct control-plane types
  - CompositeDeploymentOperation (current)
  - CompositeMonitoringOperation (hypothetical)
  - CompositeRolloutOperation (hypothetical)
- ✅ Observe 2+ shared patterns emerge naturally
  - All follow read → validate → execute → capture pattern
  - All have same failure semantics
  - All preserve step identity identically
- ✅ Cognitive compressibility still holds after 3+ types
  - System still fits in one mental model
  - No "framework thinking" required
  - Still bounded, still explicit

**ONLY THEN**:
- Consider extraction of shared base class (minimal)
- Consider generic result aggregation (if pattern is identical)
- Consider shared failure handling (if logic is identical)

**Critical Condition**:
- Extraction must NOT sacrifice step identity
- Extraction must NOT enable conditional execution
- Extraction must NOT introduce generic containers
- Extraction must NOT add framework capabilities

---

## Red Flag Checklist

If you notice ANY of these, you're near the collapse point:

- [ ] "This step handler should be pluggable"
- [ ] "Let's make the step order configurable"
- [ ] "We should support custom control-plane types"
- [ ] "Let's add a before/after hook system"
- [ ] "We need an control-plane registry"
- [ ] "Should this step be optional?"
- [ ] "Let's cache the last result"
- [ ] "We need control-plane context"
- [ ] "Should we retry on failure?"
- [ ] "We should emit events on step completion"
- [ ] "Let's add compensation logic"
- [ ] "Should we support middleware?"
- [ ] "We need an control-plane template system"
- [ ] "Let's make results generic with `Result<T>`"
- [ ] "We should pipeline operations together"

**If ANY of these thoughts appear**: STOP. Return to observation phase.

---

## Governance for Future Decisions

### Decision Point: "Add Feature X"

Ask these questions IN ORDER:

1. **Does it enable framework behavior?**
   - Plugins? Registries? Templates? Middleware? Hooks?
   - If YES → Forbidden. Stop.

2. **Does it require 3+ existing control-planes to justify?**
   - Currently have 1 control-plane
   - If NO → Not enough pressure. Stop.

3. **Does it sacrifice step identity?**
   - Generic containers? Type erasure? Collapsed traces?
   - If YES → Forbidden. Stop.

4. **Does it break cognitive compressibility?**
   - System still fits in one mental model?
   - If NO → Stop. Reevaluate architecture.

5. **Is it extracting pattern, not enabling it?**
   - Only pulling out shared code from 3+ existing types?
   - Not enabling new patterns?
   - If NO → Stop.

If all five pass, the feature is justified.

---

## Core Principle

**Right Now**: Explicit composition > Flexible control-plane

**Later** (after 3+ types + proven pressure): Minimal abstraction may be justified

**But**: Cognitive compressibility must survive. Always.

---

## Most Important Guardian

The person reviewing future PRs for this system should ask:

> "If a new engineer reads this code without knowing the system, will they still understand it fits in one mental model?"

If answer is NO, the change is too complex.

Protect cognitive compressibility aggressively.
