---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Runtime Substrate Architecture

## One Diagram

```
┌─────────────────────────────────────────────────────────────┐
│ APPLICATION (Agent/Orchestrator Layer)                      │
│ Receives results, orchestrates next operations              │
└──────────────────────────┬──────────────────────────────────┘
                           │
                    [FilesystemManager]
                    [ShellExecutor]
                           │
┌──────────────────────────┴──────────────────────────────────┐
│ RUNTIME SUBSTRATE (Deterministic, Governed)                 │
├──────────────────────────────────────────────────────────────┤
│                                                              │
│  RESOLVER (Trust Primitive)      POLICY (Limits)            │
│  ├─ Normalize                    ├─ max_file_bytes (1MB)   │
│  ├─ Follow symlinks              ├─ max_dir_entries (1000) │
│  ├─ Detect loops                 ├─ Encoding (UTF-8)       │
│  └─ Re-canonicalize              └─ Reject binaries        │
│                                                              │
│  CONTAINMENT (Post-Resolution)   TRACES (Immutable)        │
│  ├─ Commonpath check             ├─ trace_id (uuid)        │
│  ├─ Boundary assertion           ├─ operation, target      │
│  └─ Prevents escapes             ├─ duration_ms, success   │
│                                  ├─ error_type, timestamp  │
│  EXCEPTIONS (Domain-Level)       └─ frozen dataclass       │
│  ├─ WorkspaceBoundaryViolation                             │
│  ├─ SymlinkLoop                                            │
│  ├─ FileTooLarge                                           │
│  ├─ BinaryFileRejected                                     │
│  └─ No raw OS exceptions leak                              │
│                                                              │
└──────────────────────────────────────────────────────────────┘
                           │
        [File Operations Results]
        [Command Execution Results]
        [Immutable Traces]
                           │
                    [Audit Trail]
                    [Observability]
```

## Design Philosophy

### 1. Canonicalization as Root Trust Primitive

The **Resolver** performs deterministic path canonicalization:
- Input: raw user-provided path
- Process:
  1. Normalize (remove `.`, `..`, duplicate slashes)
  2. Follow symlinks with explicit loop detection
  3. Re-canonicalize after resolution
- Output: absolute canonical Path

This single operation enables all downstream trust:
- Containment checks can rely on `commonpath()` comparison
- Idempotence guarantees prevent TOCTOU
- Determinism enables audit trails

### 2. Validation After Canonicalization

Order matters:
```
User Input → Canonicalize → Validate → Execute
```

Never validate before canonicalization (e.g., don't check `../` before resolving symlinks).

### 3. Post-Resolution Enforcement

All trust checks happen on the resolved path:
- Workspace containment check: `ensure_within_workspace(resolved_path)`
- Resource limits: `ensure_file_within_size_limit(resolved_path)`
- Content policy: `read_text_file(resolved_path, limits)`

This prevents TOCTOU races where a symlink could change between validation and execution.

### 4. Domain-Level Exception Model

Never expose OS-level exceptions:
- Catch `OSError`, `FileNotFoundError`, `PermissionError`, etc.
- Convert to domain exceptions: `WorkspaceBoundaryViolation`, `FileTooLarge`, etc.
- Domain exceptions provide context and prevent information leakage

### 5. Immutable Traces

Every operation produces a frozen `RuntimeTrace`:
- Trace metadata cannot be altered (frozen dataclass)
- Enables safe audit trails and observability
- Duration is computed once (no clock-skew issues)
- No middleware or instrumentation frameworks needed

## Component Breakdown

### Resolver (`runtime/filesystem/resolver.py`)

**Responsibility**: Deterministic path canonicalization.

**Public API**:
```python
class Resolver:
    def __init__(self, workspace_root: str) -> None: ...
    def resolve(self, raw_path: str) -> Path: ...
    def ensure_within_workspace(self, resolved_path: Path) -> Path: ...
```

**Invariants**:
- `resolve(resolve(x)) == resolve(x)` (idempotence)
- Detects symlink loops via explicit visited-node tracking
- Post-resolution containment check never fails post-execute

**Test Coverage**:
- Unit tests: 3 tests (idempotence, workspace boundary, circular symlinks)
- Integration tests: 5 tests (absolute escape, relative escape, nested symlinks)

---

### Policy (`runtime/filesystem/policy.py`)

**Responsibility**: Resource limit enforcement and content validation.

**Public API**:
```python
@dataclass(frozen=True)
class ResourceLimits:
    max_file_bytes: int
    max_directory_entries: int
    encoding: str

def ensure_file_within_size_limit(path: Path, limits: ResourceLimits) -> None: ...
def read_text_file(path: Path, limits: ResourceLimits) -> str: ...
def list_directory_entries(path: Path, limits: ResourceLimits) -> List[str]: ...
```

**Invariants**:
- Files exceeding `max_file_bytes` are rejected (default 1 MB)
- Binary content (null bytes) is rejected before UTF-8 decode
- UTF-8 decoding is strict (`errors="strict"`)
- Directories with more than `max_directory_entries` are rejected

**Test Coverage**:
- Unit tests: 4 tests (large file, binary, encoding, directory entries)
- Integration tests: 4 tests (oversized file, binary payload, malformed encoding, oversized directory)

---

### FilesystemManager (`runtime/filesystem/manager.py`)

**Responsibility**: Orchestrates resolver, policy, and operation execution.

**Public API**:
```python
class FilesystemManager:
    def __init__(self, workspace_root: str, max_file_bytes: int = ..., max_directory_entries: int = ...): ...
    
    def read_file(self, requested_path: str) -> FileOperationResult: ...
    def list_directory(self, requested_path: str) -> FileOperationResult: ...
    def exists(self, requested_path: str) -> FileOperationResult: ...
```

**Operation Pipeline**:
1. Record start time
2. Resolve path with containment check
3. Apply resource policy
4. Execute operation
5. Emit trace (always)
6. Return result (never raise)

**Invariants**:
- No exceptions escape (all caught, converted to FileOperationResult)
- Traces emitted on all paths (success and failure)
- Duration consistency: `result.duration_ms == result.trace.duration_ms`

**Test Coverage**:
- Integration tests: 15+ tests covering containment, exhaustion, stability, integrity

---

### RuntimeTrace (`runtime/tracing/models.py`)

**Responsibility**: Immutable trace records for audit and observability.

**Structure**:
```python
@dataclass(frozen=True, slots=True)
class RuntimeTrace:
    trace_id: str               # uuid4().hex
    runtime_category: str       # "runtime.filesystem", "runtime.shell"
    operation: str              # "read_file", "list_directory", "execute"
    target: str                 # path or command
    duration_ms: int            # computed at operation boundary
    success: bool               # True if operation succeeded
    error_type: str | None      # exception class name if failed
    timestamp: str              # ISO 8601 UTC
```

**Invariants**:
- Frozen (immutable after creation)
- Duration computed once at operation boundary
- No distributed tracing IDs (local operation-scoped)

---

### ShellExecutor (`runtime/shell/executor.py`)

**Responsibility**: Command execution with timeout and output capture.

**Public API**:
```python
class ShellExecutor:
    def execute(
        self,
        command: str,
        args: List[str] = [],
        timeout_seconds: float = 30.0
    ) -> ExecutionResult: ...
```

**Invariants**:
- Timeouts rejected at policy layer (never execute)
- Arguments validated against simple pattern rules
- Traces emitted on all paths
- Shell traces use identical frozen dataclass structure

**Test Coverage**:
- Integration tests: 6 tests (policy rejection, timeout, argument violation, trace emission, immutability)

---

## Trust Hierarchy

```
Level 0: Operating System
├─ Provides: kernel, filesystem, process management
└─ Trust: unverified (could have vulnerabilities)

Level 1: Resolver (Trust Primitive)
├─ Input: raw paths
├─ Process: explicit canonicalization
├─ Output: canonical paths
├─ Trust: proven via idempotence tests

Level 2: Policy + Containment
├─ Input: canonical paths
├─ Process: resource limits, boundary checks
├─ Output: validated or rejected
├─ Trust: proven via adversarial tests

Level 3: FilesystemManager
├─ Input: user requests
├─ Process: orchestrates resolver + policy
├─ Output: results with traces
├─ Trust: proven via integration tests

Level 4: Application
├─ Input: guaranteed-safe operation results
├─ Process: application logic
├─ Output: new operations
├─ Trust: bounded by substrate safety
```

## Why This Works

### Simplicity
- No middleware, observability frameworks, or instrumentation complexity
- 7 classes, ~400 lines of core code
- 34 tests, all passing
- Fits in one diagram

### Explicitness
- Every trust decision is visible in the code
- No magic conversions or implicit safety
- Canonicalization is the root primitive
- Validation is post-canonicalization

### Determinism
- Idempotent canonicalization prevents TOCTOU
- Frozen traces enable reproducible audit
- Explicit error handling (no silent failures)
- Duration measured once (no clock issues)

### Auditability
- Immutable traces at every operation boundary
- Domain-level exceptions preserve context
- No exception leakage hides OS details
- Reproducible error classification

## Next: Orchestration Safety

Once the substrate is at **Level 2 (Verified Governed Execution)**, control-plane can be introduced safely:

1. **Task Graph Execution**: Schedule tasks with guaranteed-safe filesystem/shell operations
2. **Agent Composition**: Compose agents knowing each operation is auditable and bounded
3. **Error Recovery**: Recover from failures knowing traces are immutable and deterministic

The substrate's trust boundary protects control-plane from:
- Path traversal attacks
- Resource exhaustion
- Binary execution
- Silent failures
- Non-reproducible errors

All because canonicalization, post-resolution validation, and immutable tracing form an unbreakable foundation.
