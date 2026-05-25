# ASTRAEUS - MASTER SPECIFICATION FOR EXECUTION SUBSTRATE

**Version**: 1.0  
**Status**: Critical Audit & Hardening Phase  
**Last Updated**: May 2026  
**Target**: Deterministic Repository-Grounded Engineering Execution Substrate  

---

## TABLE OF CONTENTS

1. [EXECUTIVE OVERVIEW](#executive-overview)
2. [CURRENT STATE ASSESSMENT](#current-state-assessment)
3. [PROJECT MISSION](#project-mission)
4. [HARD CONSTRAINTS](#hard-constraints)
5. [ARCHITECTURE OVERVIEW](#architecture-overview)
6. [IMPLEMENTATION SPECIFICATION](#implementation-specification)
7. [UI/DOCUMENT SPECIFICATION](#uiinput-specification)
8. [EXECUTION RULES](#execution-rules)
9. [VERIFICATION & MONITORING](#verification--monitoring)
10. [LONG-RUNNING OPERATION](#long-running-operation)
11. [TROUBLESHOOTING GUIDE](#troubleshooting-guide)
12. [CODEX INSTRUCTIONS](#codex-instructions)

---

## EXECUTIVE OVERVIEW

### What You're Building
A **deterministic engineering execution substrate** that:
- Decomposes goals into verifiable task DAGs
- Routes tasks to specialized models using a strict control-plane
- Executes in a policy-governed mutation sandbox
- Enforces filesystem-level transactional integrity
- Maintains an append-only, immutable event record
- Provides reconstructible state through temporal replay
- Asks for help when invariants are violated

### What This Is NOT
- An AI coding agent or chatbot
- Unrestricted or unobserved autonomy
- A generic agent framework
- A non-grounded cognition system

### Current Technical Reality
🔴 **PROTOTYPE BOTTLE_NECKS**:
- **Event Bus**: Single-file `fcntl` locking bottleneck (Non-distributed)
- **Snapshots**: Slow $O(N)$ `shutil.copy2` file-copy mechanism
- **Replay**: Semantic-only log hydration (Environment non-deterministic)
- **Context**: Hardcoded 6000-char truncation in orchestration
- **Security**: Naive path-matching in Risk Engine

**Your job**: Strip the prototype abstractions and harden the core substrate into a production-grade reliable engineering runtime.

---

## CRITICAL STATE ASSESSMENT

### Current Logic Path
```
User Goal
    ↓
Planner → Decompose (Untrusted)
    ↓
Router → Assign Models (Heuristic)
    ↓
DAG → Topological Sort (Graph-bound)
    ↓
Queue → Execute Sequentially (Local Process)
    ↓
Validator → Check Output (Regex/Critic)
    ↓
Mutation Sandbox → Apply (Staged Filesystem)
    ↓
Event Bus → Emit (Blocking I/O)
```

### Critical Vulnerabilities
1. **Mutation Collision**: `shutil.copy2` is not atomic; power loss during snapshot shreds state.
2. **Lock Contention**: `fcntl.flock` blocks scaling; cannot distribute workers.
3. **Context Poisoning**: Blind 6000-char slicing corrupts upstream dependency data.
4. **Environment Drift**: Replay ignores OS-level dependencies (`pip`, `npm`, `env`).
5. **Autonomy Escape**: Risk engine uses fragile filename-matching.

### Current Limitations
- Deterministic fallbacks hide real failures (disable in production)
- Model adapters timeout after 180s (DeepSeek)
- Repair taxonomy still small (adding incrementally)
- No active UI (use CLI + artifact viewer)

---

## SYSTEM MISSION

Astraeus exists to provide a **deterministic engineering execution substrate**. Its mission is to bridge the gap between hallucinated LLM outputs and verifiable repository state, ensuring every mutation is attributable, reversible, and grounded in architectural truth.
Week 5-6:  Repo intelligence
Week 7-8:  Transaction safety
Month 2:   Long-session operations
Month 3+:  Optional: UI, plugins, distributed
```

---

## HARD CONSTRAINTS

### MANDATORY EXECUTION INVARIANTS

```yaml
mandatory:
  one_active_model_only: true
    reason: Hardware bounds + sequential causality
    enforcement: asyncio.Lock on model access
  
  deterministic_offline_first: true
    reason: Zero-trust network + replay stability
    enforcement: Local Ollama only; no external telemetry
  
  append_only_immutable_events: true
    reason: Lineage traceability + audit integrity
    enforcement: File-locked JSONL or ACID SQLite
  
  bounded_retry_budgets: true
    reason: Resource exhaustion prevention
    enforcement: max_retries=2, repair_retries=1
  
  strict_semantic_validation: true
    reason: Hallucination containment
    enforcement: Invariant-engine check on all artifacts
  
  filesystem_transaction_safety: true
    reason: Atomic state transitions
    enforcement: Snapshot/Journal before ANY mutation
  
  lineage_artifact_persistence: true
    reason: Replayability + cognitive grounding
    enforcement: Store every attempt, prompt, and validation
  
  causal_determinism: true
    reason: Distributed consistency
    enforcement: Sequence-ordered event folding
```

### Hardware Requirements
```yaml
minimum:
  cpu: 4 cores
  ram: 8GB (M3 compatible)
  storage: 50GB free (artifacts + models)
  network: offline-capable

optimal:
  cpu: 8 cores
  ram: 16GB
  storage: 100GB free
```

### Model Requirements
```yaml
models:
  planning:
    name: "phi3:mini"
    size: "3.8B"
    purpose: "Task decomposition"
    memory: "2GB"
  
  coding:
    name: "qwen2.5-coder"
    size: "7B"
    purpose: "Code generation"
    memory: "4GB"
  
  debugging:
    name: "deepseek-coder"
    size: "6.7B"
    purpose: "Analysis + repair"
    memory: "4GB"
  
  documentation:
    name: "mistral"
    size: "7B"
    purpose: "Clarity + guides"
    memory: "4GB"
  
  synthesis:
    name: "codellama"
    size: "7B"
    purpose: "Alternative generation"
    memory: "4GB"

note: "Only ONE active at a time. Ollama handles model loading."
```

---

## ARCHITECTURE OVERVIEW

### Layer 1: Goal Decomposition (Untrusted)
The system treats user goals as fuzzy inputs that must be transformed into a verifiable Directed Acyclic Graph (DAG) of discrete engineering tasks.

**Files**:
- `planner/decomposer.py`: Transforms goals into `TaskNode` objects.
- `contracts/models.py`: Defines the strict schema for `TaskType` and `TaskNode`.
- `validator/dag_validator.py`: Ensures no circular dependencies and validates task coverage.

### Layer 2: Control-Plane & Routing
The control-plane manages the execution lifecycle, ensuring only one model is active and all state transitions are recorded.

**Files**:
- `orchestrator/router.py`: Heuristic-based routing of tasks to specialized model adapters.
- `orchestrator/dag.py`: Maintains the authoritative execution graph and manages task state (READY, RUNNING, COMPLETED, FAILED).
- `orchestrator/control_plane.py`: The central coordinator (ControlPlane) that wires the control-plane to the event bus and mutation sandbox.

### Layer 3: Safety Substrate (Mutation Sandbox)
Every operation that affects repository state MUST pass through the safety substrate. This is the boundary between cognition and reality.

**Files**:
- `runtime/mutation_sandbox.py`: Staged filesystem transaction manager.
- `runtime/risk_engine.py`: Classifies mutation danger levels and enforces approval gates.
- `runtime/snapshots.py`: Atomic state capturing (Warning: Current implementation uses slow file-copy).
- `transactions/journal.py`: Append-only hash-chained journal of all filesystem deltas.

### Layer 4: Verification & Repair
Failed tasks trigger an automated repair loop. Repaired artifacts are re-validated before being committed to the state.

**Files**:
- `validator/critic.py`: LLM-based semantic validation of artifacts against goal invariants.
- `repair/repair_planner.py`: Generates bounded repair DAGs for classified failure types.
- `orchestrator/control_plane.py`: Manages the recursion between task execution, failure capture, and repair planning.
    TaskType.CODE_GENERATION: ModelType.QWEN,
    TaskType.DEBUGGING: ModelType.DEEPSEEK,
    TaskType.DOCUMENTATION: ModelType.MISTRAL,
    TaskType.ARCHITECTURE: ModelType.QWEN,
    TaskType.EXTRACTION: ModelType.QWEN,
    TaskType.PLANNING: ModelType.PHI,
    TaskType.SYNTHESIS: ModelType.CODELLAMA,
}
```

### Layer 3: Sequential Execution
```
Ready Tasks (with dependencies met)
    ↓
Execution Queue (asyncio)
    ↓
[One model active at a time]
    ↓
Model Inference (via Ollama)
    ↓
Output Parsing
    ↓
Validation
    ↓
Store Result
    ↓
Update DAG State
    ↓
Compute Next Ready Tasks
```

**Files**:
- `orchestrator/queue.py` - ExecutionQueue
- `orchestrator/control_plane.py` - ControlPlane (main coordinator)
- `models/base_adapter.py` - BaseModelAdapter
- `models/{qwen,deepseek,mistral,codellama,phi}_adapter.py` - Model-specific adapters

**Execution Guarantees**:
```python
class ExecutionQueue:
    async def execute(self, dag: TaskGraph) -> dict[str, GenerationResult]:
        """
        Guarantees:
        1. Only one model runs at a time
        2. Tasks respect dependencies
        3. Every output is validated
        4. Results are persisted immediately
        5. Events are emitted for every state change
        """
```

### Layer 4: Validation & Failure Capture
```
Model Output
    ↓
Schema Validation (Pydantic)
    ↓
Syntax Check (for code)
    ↓
Type Check (for code)
    ↓
Content Validation (custom)
    ↓
[If Valid] → Continue
[If Invalid] → Create FailureRecord
```

**Files**:
- `validator/validator.py` - ValidationPipeline
- `validator/failure_types.py` - FailureType enum
- `validator/failure_record.py` - FailureRecord

**Failure Types** (continuously expanded):
```python
class FailureType(Enum):
    SYNTAX_ERROR = "syntax_error"
    IMPORT_ERROR = "import_error"
    DEPENDENCY_MISSING = "dependency_missing"
    TEST_FAILURE = "test_failure"
    RUNTIME_EXCEPTION = "runtime_exception"
    TIMEOUT = "timeout"
    HALLUCINATED_API = "hallucinated_api"
    PLANNER_INVALID = "planner_invalid"
    PERMISSION_DENIED = "permission_denied"
    INVALID_DIFF = "invalid_diff"
    UNCLASSIFIED = "unclassified"
```

### Layer 5: Repair & Localized Retry
```
FailureRecord
    ↓
Repair Planner (deepseek)
    ↓
Generate Repair Task DAG
    ↓
Execute Repair Tasks
    ↓
[If Success] → Compute affected subtree
           → Invalidate downstream artifacts
           → Rerun only affected tasks
[If Fail] → Create HelpRequest
```

**Files**:
- `repair/repair_planner.py` - RepairPlanner
- `orchestrator/dag.py` - compute_affected_subgraph()
- `help/help_request.py` - HelpRequest

**Repair Strategies** (Add incrementally):
```python
repair_strategies = {
    FailureType.IMPORT_ERROR: [
        "pip install {package}",
        "pip install -r requirements.txt"
    ],
    FailureType.SYNTAX_ERROR: [
        "fix_syntax(code, error_location)",
        "regenerate_from_scratch()"
    ],
    FailureType.TEST_FAILURE: [
        "isolate_failing_module",
        "regenerate_implementation",
        "adjust_test_assumptions"
    ],
}
```

### Layer 6: Memory Systems (Separated by Purpose)
```
Memory System
    ├── Semantic Store (ChromaDB)
    │   └── Repository code understanding
    │
    ├── Session Memory (In-memory dict)
    │   └── Active run context
    │
    ├── Failure Memory (SQLite)
    │   └── Debugging history + repair outcomes
    │
    └── Architecture Memory (SQLite)
        └── ADRs, invariants, design decisions
```

**Files**:
- `memory/semantic_store.py` - SemanticMemoryStore
- `memory/session_memory.py` - SessionMemory
- `memory/failure_memory.py` - FailureMemory
- `memory/architecture_memory.py` - ArchitectureMemory
- `memory/memory_system.py` - MemorySystem (coordinator)

**Critical Rule**: Do NOT use monolithic memory blobs. Each subsystem has specific purpose.

### Layer 7: Sandbox & Transactions
```
Code Output
    ↓
[If modifies project] → Create Snapshot
    ↓
Generate DiffPlan
    ↓
Apply to Staging Directory
    ↓
Run Validation (tests, lint, type-check)
    ↓
[If passes] → Atomic Commit to Live
[If fails] → Rollback + Repair
```

**Files**:
- `runtime/sandbox.py` - ExecutionSandbox
- `runtime/snapshots.py` - SnapshotEngine
- `transactions/diff_plan.py` - DiffPlan
- `transactions/runner.py` - TransactionRunner

**Staging Flow**:
```python
async def safe_file_edit(original_files: dict, new_content: dict) -> bool:
    """
    1. Create snapshot of original
    2. Apply edits to staging/
    3. Run validation on staging/
    4. If valid: copy staging/ → original
    5. If invalid: restore from snapshot
    """
```

### Layer 8: Long-Running State
```
Session Start
    ↓
Create Session ID
    ↓
Store SharedContext in SQLite
    ↓
Execute control-plane
    ↓
After each task:
    - Persist artifacts
    - Emit events
    - Update session state
    ↓
Process crashes
    ↓
Resume from checkpoint:
    - Restore SharedContext
    - Skip completed tasks
    - Resume pending tasks
```

**Files**:
- `shared_context/state.py` - SharedContext
- `runtime/snapshots.py` - Resume logic
- Database schema in `memory/store.py`

---

## IMPLEMENTATION SPECIFICATION (SUBSTRATE CORE)

### Core Substrate Directory Structure
```
astraeus-core/
│
├── contracts/
│   ├── models.py              # Authoritative Type System (TaskType, EventAction)
│   └── invariant_engine.py    # Policy Enforcement (Pre-flight validation)
│
├── planner/
│   └── decomposer.py          # Fuzzy Goal → Verifiable Task Graph
│
├── orchestrator/
│   ├── engine.py              # Control-Plane Coordinator (ControlPlane)
│   ├── dag.py                 # Causal Execution Graph (State machine)
│   └── queue.py               # Sequential Execution Dispatcher
│
├── runtime/
│   ├── mutation_sandbox.py    # Boundary between Cognition and Filesystem
│   ├── risk_engine.py         # Heuristic & Policy Danger Classification
│   ├── journal_manager.py     # Hash-chained Filesystem Delta Journaling
│   ├── snapshots.py           # State Capturing (Current: O(N) copy)
│   └── replay.py              # Temporal State Reconstruction
│
├── events/
│   ├── event_bus.py           # Immutable Append-Only Log (Current: fcntl-locked)
│   └── schema.py              # Lineage-Traceable Event Definitions
│
├── validator/
│   └── critic.py              # Semantic Invariant Verification
│
├── memory/
│   ├── store.py               # ACID Persistence (SQLite)
│   └── semantic_store.py      # Repository Grounding (Embeddings)
```
├── sandbox/
│   ├── __init__.py
│   ├── docker_runner.py       # Docker execution (future)
│   ├── process_limits.py      # CPU/memory/time limits
│   └── seccomp_rules.py       # Seccomp profiles (future)
│
├── memory/
│   ├── __init__.py            # MemorySystem (coordinator)
│   ├── store.py               # SQLite schema + queries
│   ├── semantic_store.py      # ChromaDB integration
│   ├── session_memory.py      # In-memory session state
│   ├── failure_memory.py      # Failure database
│   └── architecture_memory.py # ADR database
│
├── shared_context/
│   ├── __init__.py
│   ├── state.py               # SharedContext model
│   └── artifacts.py           # ArtifactStore
│
├── events/
│   ├── __init__.py
│   ├── schema.py              # RuntimeEvent model
│   ├── event_bus.py           # EventBus (pub-sub)
│   └── handlers.py            # Event handlers
│
├── metrics/
│   ├── __init__.py
│   ├── store.py               # MetricsStore
│   └── collector.py           # Metrics collection
│
├── tools/
│   ├── __init__.py
│   ├── permissions.py         # Permission system
│   └── approval_gate.py       # Approval requests
│
├── repo_indexer/
│   ├── __init__.py
│   ├── indexer.py             # RepoIndexer
│   ├── ast_parser.py          # AST extraction
│   ├── symbol_extractor.py    # Symbol extraction
│   ├── dependency_graph.py    # Dependency analysis
│   └── architecture_rules.py  # Architecture invariants
│
├── transactions/
│   ├── __init__.py
│   ├── diff_plan.py           # DiffPlan model
│   ├── runner.py              # TransactionRunner
│   └── validator.py           # Transaction validation
│
├── help/
│   ├── __init__.py
│   └── help_request.py        # HelpRequest model
│
├── cli/
│   ├── __init__.py
│   └── main.py                # CLI entry point
│
├── api/
│   ├── __init__.py
│   ├── main.py                # FastAPI server
│   ├── routes.py              # API endpoints
│   └── websocket.py           # WebSocket streaming
│
├── frontend-console/
│   ├── index.html             # Simple UI (see UI spec)
│   ├── main.js                # WebSocket + UI logic
│   └── styles.css             # Minimal styling
│
├── artifacts/
│   └── .gitkeep               # Run artifacts stored here
│
├── tests/
│   ├── __init__.py
│   ├── test_contracts.py
│   ├── test_planner.py
│   ├── test_router.py
│   ├── test_dag.py
│   ├── test_queue.py
│   ├── test_validator.py
│   ├── test_repair.py
│   ├── test_memory.py
│   ├── test_sandbox.py
│   ├── test_transactions.py
│   ├── test_e2e_control-plane.py
│   └── test_long_session.py
│
├── scripts/
│   ├── setup.sh               # Environment setup
│   ├── verify_ollama.py       # Check Ollama health
│   ├── lint.sh                # Ruff + MyPy
│   └── test.sh                # Run all tests
│
├── data/
│   └── .gitkeep               # Persistent data (SQLite, etc.)
│
├── pyproject.toml             # Dependencies
├── Makefile                   # Task automation
├── README.md                  # Quick start
├── ARCHITECTURE.md            # Design decisions
├── INTEGRATION_GUIDE.md       # How to integrate
├── ROADMAP.md                 # Future features
└── MASTER_SPEC.md             # This file
```

### Key Data Classes (contracts/models.py)

```python
from enum import Enum
from dataclasses import dataclass
from pydantic import BaseModel
from typing import Optional, List

# ─────────────────────────────────
# TASK MODELS
# ─────────────────────────────────

class TaskType(Enum):
    CODE_GENERATION = "code_generation"
    DEBUGGING = "debugging"
    DOCUMENTATION = "documentation"
    ARCHITECTURE = "architecture"
    EXTRACTION = "extraction"
    PLANNING = "planning"
    SYNTHESIS = "synthesis"
    TESTING = "testing"
    SCAFFOLDING = "scaffolding"

class ModelType(Enum):
    PHI = "phi3:mini"
    QWEN = "qwen2.5-coder"
    DEEPSEEK = "deepseek-coder"
    MISTRAL = "mistral"
    CODELLAMA = "codellama"

class TaskNode(BaseModel):
    id: str                          # Unique task ID
    run_id: str                      # Which control-plane run
    type: TaskType                   # Task category
    goal: str                        # What to accomplish
    depends_on: List[str] = []       # Task IDs this depends on
    assigned_model: ModelType        # Which model to use
    validation: List[str]            # How to validate output
    status: str = "pending"          # pending, ready, running, completed, failed
    retry_count: int = 0             # Current retry count
    repair_count: int = 0            # How many repairs attempted
    result: Optional[str] = None     # Output from model
    error: Optional[str] = None      # Error message if failed
    created_at: str = None           # Timestamp
    started_at: Optional[str] = None
    completed_at: Optional[str] = None

class GenerationResult(BaseModel):
    task_id: str
    output: str
    model_used: ModelType
    tokens_used: int
    latency_ms: float
    metadata: dict = {}

# ─────────────────────────────────
# FAILURE MODELS
# ─────────────────────────────────

class FailureType(Enum):
    SYNTAX_ERROR = "syntax_error"
    IMPORT_ERROR = "import_error"
    DEPENDENCY_MISSING = "dependency_missing"
    TEST_FAILURE = "test_failure"
    RUNTIME_EXCEPTION = "runtime_exception"
    TIMEOUT = "timeout"
    HALLUCINATED_API = "hallucinated_api"
    PLANNER_INVALID = "planner_invalid"
    PERMISSION_DENIED = "permission_denied"
    INVALID_DIFF = "invalid_diff"
    UNCLASSIFIED = "unclassified"

class FailureRecord(BaseModel):
    failure_id: str
    run_id: str
    task_id: str
    failure_type: FailureType
    raw_error: str
    structured_context: dict
    attempt: int
    timestamp: str
    resolution_status: str = "pending"  # pending, repaired, escalated, unresolved

# ─────────────────────────────────
# MEMORY MODELS
# ─────────────────────────────────

class MemoryType(Enum):
    SEMANTIC = "semantic"
    SESSION = "session"
    FAILURE = "failure"
    ARCHITECTURE = "architecture"

class MemoryRecord(BaseModel):
    memory_type: MemoryType
    key: str
    value: str
    embedding: Optional[List[float]] = None
    created_at: str
    last_accessed: str
    score: float = 1.0

# ─────────────────────────────────
# SESSION MODELS
# ─────────────────────────────────

class SharedContext(BaseModel):
    session_id: str
    run_id: str
    original_prompt: str
    task_dag: List[TaskNode]
    artifacts_dir: str
    snapshots: dict = {}
    approved_permissions: List[str] = []
    user_guidance: Optional[str] = None
    started_at: str
    last_updated_at: str

# ─────────────────────────────────
# EVENTS
# ─────────────────────────────────

class RuntimeEvent(BaseModel):
    timestamp: str
    run_id: str
    task_id: Optional[str]
    action: str
    status: str
    metadata: dict = {}
    # Actions: task_created, task_started, task_completed, 
    #          task_failed, repair_attempted, approval_requested, etc.

# ─────────────────────────────────
# HELP & ESCALATION
# ─────────────────────────────────

class HelpRequest(BaseModel):
    help_id: str
    run_id: str
    blocked_task_id: str
    problem_summary: str
    evidence: dict  # Failure context, attempted repairs, etc.
    suggestions: List[str]
    required_permission: Optional[str] = None
    created_at: str
```

---

## REPOSITORY & RUNTIME INTERFACE SPECIFICATION

### System Grounding
All interfaces MUST ground their state in the repository and the event bus. There is no "floating" UI state. If a task is not in the event log, it does not exist in the interface.

### Mandatory Visualizations
- **Causal Graph**: Real-time DAG state with lineage links to event sequence IDs.
- **Mutation Journal**: Live diff stream of the filesystem transaction log.
- **Temporal Replay Slider**: Interface for scrubbing through the event log to reconstruct past state.
- **Risk Gate**: High-visibility confirmation UI for operations classified as MEDIUM or DANGEROUS risk.

#### Substrate Dashboard Layout
```
┌────────────────────────────────────────────────────────┐
│               ASTRAEUS EXECUTION CONSOLE               │
├──────────────────┬──────────────────┬──────────────────┤
│                  │                  │                  │
│   CONTROL PLANE  │   CAUSAL GRAPH   │   MUTATION       │
│   (Left 25%)     │   (Center 50%)   │   JOURNAL        │
│                  │                  │   (Right 25%)    │
│                  │                  │                  │
│                  │                  │                  │
└──────────────────┴──────────────────┴──────────────────┘
│            EVENT BUS STREAM (Bottom 100%)              │
└────────────────────────────────────────────────────────┘
```
```

#### CENTER PANEL: Execution Visualization

```html
<div class="execution-panel">
  <h3>Task DAG</h3>
  <div id="dag-canvas"></div>
  
  <h3>Active Model</h3>
  <div class="active-model">
    <p>Model: <code id="active-model">-</code></p>
    <p>Progress: <progress id="progress" max="100" value="0"></progress></p>
    <p>Tokens: <span id="tokens">0</span> / Latency: <span id="latency">0ms</span></p>
  </div>
  
  <h3>Task Queue</h3>
  <div id="task-queue" class="task-list">
    <!-- Tasks appear here dynamically -->
  </div>
  
  <h3>Retries & Repairs</h3>
  <div id="retries" class="retry-log">
    <!-- Retry history -->
  </div>
</div>
```

#### RIGHT PANEL: Artifacts & Results

```html
<div class="artifacts-panel">
  <h3>Generated Artifacts</h3>
  <div id="artifacts" class="artifact-list">
    <!-- Artifact files appear here -->
  </div>
  
  <h3>Preview</h3>
  <div id="artifact-preview" class="preview-box">
    <!-- Selected artifact preview -->
  </div>
  
  <h3>Diff Viewer</h3>
  <div id="diff-viewer" class="diff-box">
    <!-- Show changes to files -->
  </div>
</div>
```

#### BOTTOM PANEL: Detailed Logs

```html
<div class="status-panel">
  <div class="tabs">
    <button class="tab-btn active" onclick="showTab('events')">Events</button>
    <button class="tab-btn" onclick="showTab('logs')">Logs</button>
    <button class="tab-btn" onclick="showTab('metrics')">Metrics</button>
    <button class="tab-btn" onclick="showTab('help')">Help</button>
  </div>
  
  <div id="events-tab" class="tab-content active">
    <!-- Event stream from event bus -->
    <div id="event-log" class="log-box"></div>
  </div>
  
  <div id="logs-tab" class="tab-content">
    <!-- Task logs and stderr -->
    <div id="task-logs" class="log-box"></div>
  </div>
  
  <div id="metrics-tab" class="tab-content">
    <!-- Metrics dashboard -->
    <div id="metrics-dashboard"></div>
  </div>
  
  <div id="help-tab" class="tab-content">
    <!-- Current help request -->
    <div id="help-display"></div>
  </div>
</div>
```

### WebSocket Communication Protocol

```javascript
// Frontend → Backend
const ws = new WebSocket("ws://localhost:8000/orchestrate");

// User submits prompt
ws.send(JSON.stringify({
  action: "decompose",
  payload: {
    prompt: "Build a REST API with JWT auth, tests, and Docker setup"
  }
}));

// Backend responds with DAG
// Event: task_created, task_ready, task_started, task_completed, etc.

// User approves dangerous action
ws.send(JSON.stringify({
  action: "approve",
  payload: {
    approval_id: "...",
    decision: "approve"
  }
}));

// Backend → Frontend
ws.onmessage = (event) => {
  const msg = JSON.parse(event.data);
  
  switch(msg.type) {
    case "dag_ready":
      displayDAG(msg.tasks);
      break;
    
    case "task_started":
      updateTaskState(msg.task_id, "running");
      displayActiveModel(msg.model);
      break;
    
    case "task_completed":
      updateTaskState(msg.task_id, "completed");
      displayArtifact(msg.artifact);
      break;
    
    case "task_failed":
      updateTaskState(msg.task_id, "failed");
      displayFailure(msg.failure);
      break;
    
    case "approval_required":
      displayApprovalRequest(msg.approval);
      break;
    
    case "help_requested":
      displayHelpRequest(msg.help);
      break;
    
    case "event":
      appendToEventLog(msg.event);
      break;
  }
};
```

### Key UI Interactions

#### 1. Input Mixed Prompt
```
User types:
  "Build a REST API service with:
   - JWT authentication
   - PostgreSQL integration
   - Docker containerization
   - Unit tests
   - API documentation
   - CI/CD pipeline"

System shows:
  "Decomposing... ⏳"

System displays:
  Task DAG with 8 nodes
  All tasks → Ready to Execute
```

#### 2. Execute & Monitor
```
User clicks "Execute"

System shows:
  Task 1: RUNNING (phi: decomposing plan)
  Active Model: phi3:mini
  Progress: ████░░░░ 40%

[Task 1 completes]

  Task 2: RUNNING (qwen: scaffolding project)
  Active Model: qwen2.5-coder
  Tokens: 1,250 / Latency: 2,450ms

[Task 2 completes]

  Task 3: RUNNING (deepseek: implementing auth)
  [...]
```

#### 3. Failure & Repair
```
User sees:
  Task 4: FAILED (ImportError: No module named 'jwt')

System shows:
  REPAIR ATTEMPT 1
  Suggested fix: pip install PyJWT
  Running repair...

  REPAIR ATTEMPT 2
  Re-running task 4...

  Task 4: COMPLETED ✅

System shows:
  Affected downstream tasks (5, 6, 7):
  ↻ Recalculating...
  All passed!
```

#### 4. Approval Gate
```
User sees:
  ⚠️ APPROVAL REQUIRED
  Action: Install system package: postgresql-dev
  Permission Level: 4 (install_and_configure)
  
  [Approve] [Deny] [Modify]

User clicks [Approve]

System logs:
  APPROVED_BY=user123
  TIMESTAMP=2026-05-15T14:30:45Z
  ACTION=install_postgresql-dev
```

#### 5. Help Request
```
User sees:
  ⛔ BLOCKED
  
  Task 7: Integration test failed
  Attempted repairs: 3
  Last error: "Circular import detected"
  
  NEED HUMAN HELP:
  - Review architecture
  - Suggest import restructuring
  - Modify task parameters
  
  [Submit Guidance] [Terminate] [Modify Goal]

User types guidance:
  "Refactor auth module to separate concerns"

System resumes with modified context
```

---

## EXECUTION RULES

### Rule 1: Sequential Causal Execution
The system MUST execute tasks in a strictly ordered sequence defined by the DAG. Parallel execution is prohibited to maintain causal determinism in the event log.

### Rule 2: Single-Instance Model Lock
A global `asyncio.Lock` MUST be held during all model inference calls. This ensures hardware stability on memory-constrained devices (M3 8GB) and prevents non-deterministic interleaving of model outputs.

### Rule 3: Atomic Mutation Commits
No changes to the repository are permitted outside of a `MutationSandbox` transaction. Every commit MUST be preceded by a filesystem snapshot and recorded in the hash-chained journal.

### Rule 4: Failure Taxonomy Enforcement
Failures MUST NOT be handled as generic exceptions. Every failure MUST be classified into the typed taxonomy (`FailureType`) and persisted as a `FailureRecord` to enable structured repair planning.

### Rule 5: Deterministic Replayability
Every event emitted to the bus MUST contain enough metadata (seeds, environment hashes, dependency versions) to allow the `ReplayEngine` to reconstruct the exact state of the system at that sequence ID.

### Rule 6: Zero-Trust Model Outputs
All model outputs are treated as untrusted and potentially malicious. Every artifact MUST undergo static validation (e.g., syntax checks) and semantic verification (critic review) before being applied to the repository.
        failure_type=classify_error(e),
        raw_error=str(e),
        structured_context=extract_context(e, task)
    )
    await memory.failures.store(failure)
    return FailedTaskResult(failure)
```

### Rule 5: Localized Retry Only
```python
# CORRECT
async def handle_task_failure(task_id: str):
    failure = get_failure(task_id)
    
    # Compute ONLY affected tasks
    affected = dag.compute_affected_subgraph(task_id)
    
    # Plan repair for this task
    repair_dag = repair_planner.plan(failure)
    
    # Execute repair
    repair_results = await execute_queue(repair_dag)
    
    # ONLY rerun affected downstream
    for affected_task in affected[1:]:  # Skip original
        if is_invalidated(affected_task, repair_results):
            await execute_single_task(affected_task)

# WRONG ❌
async def handle_task_failure(task_id: str):
    # This would waste computation
    dag = decompose_original_goal()  # Start over!
    await execute_queue(dag)  # Rerun everything!
```

### Rule 6: Snapshot Before Dangerous Actions
```python
async def execute_task_that_modifies_files(task: TaskNode):
    # BEFORE mutation
    snapshot = await snapshots.create_snapshot({
        "filesystem": compute_hashes(repo_path),
        "shared_context": shared_context.state,
        "task_state": dag.serialize(),
    })
    
    try:
        result = await model.generate(...)
        
        # Apply to staging
        await transactions.apply_to_staging(result)
        
        # Validate in staging
        validation_result = await validator.validate(staging_dir)
        
        if validation_result.passed:
            # Atomic commit
            await transactions.commit_to_live()
        else:
            # Rollback automatically
            await snapshots.restore(snapshot)
            return FailedTaskResult(validation_failure)
            
    except Exception as e:
        # Always restore on exception
        await snapshots.restore(snapshot)
        raise
```

### Rule 7: Validation Before Acceptance
```python
async def validate_task_output(task: TaskNode, output: str):
    """
    Every output MUST pass validation before acceptance.
    """
    
    # 1. Schema validation
    try:
        if task.type == TaskType.CODE_GENERATION:
            parsed = parse_code(output)
        elif task.type == TaskType.PLANNING:
            parsed = json.loads(output)
    except Exception as e:
        return ValidationFailure(FailureType.SCHEMA_VIOLATION, e)
    
    # 2. Content validation
    for validator_name in task.validation:
        result = await validators[validator_name](parsed)
        if not result.passed:
            return ValidationFailure(FailureType.CONTENT_INVALID, result.error)
    
    # 3. If code: syntax + types + tests
    if task.type == TaskType.CODE_GENERATION:
        syntax_result = check_syntax(parsed)
        if not syntax_result.passed:
            return ValidationFailure(FailureType.SYNTAX_ERROR, syntax_result.error)
        
        type_result = check_types(parsed)
        if not type_result.passed:
            return ValidationFailure(FailureType.TYPE_ERROR, type_result.error)
        
        test_result = run_tests(parsed)
        if not test_result.passed:
            return ValidationFailure(FailureType.TEST_FAILURE, test_result.error)
    
    return ValidationSuccess(parsed)
```

### Rule 8: Events for Everything
```python
# Every state change emits an event
events = [
    RuntimeEvent(action="session_created", run_id=run_id),
    RuntimeEvent(action="task_created", task_id=t1.id),
    RuntimeEvent(action="task_created", task_id=t2.id),
    RuntimeEvent(action="task_ready", task_id=t1.id),
    RuntimeEvent(action="task_started", task_id=t1.id, metadata={"model": "phi"}),
    RuntimeEvent(action="task_completed", task_id=t1.id, metadata={"tokens": 1250}),
    RuntimeEvent(action="task_ready", task_id=t2.id),
    RuntimeEvent(action="task_started", task_id=t2.id, metadata={"model": "qwen"}),
    # ... more events ...
]

# Append to event log (never delete)
for event in events:
    await event_bus.emit(event)
    await metrics.log_event(event)
```

### Rule 9: Retry Budgets
```python
task_retry_budget = {
    "task_retries": 2,        # Max 2 regular retries
    "repair_retries": 1,      # Max 1 repair attempt
    "total_limit": 3,         # Max 3 total attempts
}

async def handle_task_failure(task: TaskNode):
    if task.retry_count >= task_retry_budget["task_retries"]:
        if task.repair_count >= task_retry_budget["repair_retries"]:
            # Give up, request help
            await help_system.request_help(task)
            return
        else:
            # Try repair once
            await attempt_repair(task)
            return
    else:
        # Retry the task
        await execute_single_task(task)
```

### Rule 10: Permission Escalation
```python
async def execute_dangerous_command(command: str, permission_level: int):
    """
    Only execute dangerous commands with explicit approval.
    """
    
    if permission_level > current_permission_level:
        # Request approval
        approval = await approval_gate.request_approval({
            "action": command,
            "required_level": permission_level,
            "risks": analyze_risks(command),
            "rationale": "Install PostgreSQL for database layer",
        })
        
        if not approval.granted:
            raise PermissionDenied(command)
        
        # Log the approval
        await audit_log.log({
            "action": "approval_granted",
            "command": command,
            "approved_by": approval.user_id,
            "timestamp": now(),
        })
    
    # Execute with elevated permission
    return await sandbox.execute(command, level=permission_level)
```

---

## VERIFICATION, MONITORING & REPLAY

### Lineage Verification
Every mutation MUST be traceable to an event in the `EventBus` and a task in the `TaskGraph`. Inconsistencies between the filesystem journal and the event bus indicate a **CRITICAL SUBSTRATE FAILURE**.

### Runtime Health Metrics
- **Event Sequence Velocity**: Rate of append-only log growth.
- **Lock Contention Latency**: Time spent waiting on `fcntl.flock`.
- **Snapshot I/O Pressure**: Bytes copied per task execution.
- **Replay Divergence Count**: Number of times projected state differs from persisted run state.
- **Risk Gate Denial Rate**: Frequency of security-blocked operations.

### Mandatory Health Checks
```bash
# Verify system integrity
python scripts/verify_event_integrity.py  # Sequence continuity
python scripts/verify_journal_hashes.py   # Journal hash-chain
python scripts/verify_replay_determinism.py # State re-folding
python scripts/check_db_integrity.py      # ACID check
```
    events = await event_bus.get_events_for_run(run_id)
    assert len(events) > len(tasks)
    
    # 4. Replay test
    replay_result = await replay_engine.replay(run_id)
    # Compare replay events with original events
    assert are_event_sequences_equivalent(events, replay_result.events)
    
    # 5. Metrics sanity
    metrics = await metrics_store.get_for_run(run_id)
    assert metrics["task_success_rate"] > 0.3
    assert metrics["avg_retry_count"] < 3
```

### Key Metrics to Monitor

```python
class MetricsGroup:
    """
    Monitor these EVERY run
    """
    
    planner_metrics = {
        "task_count": int,               # Should be 3-20
        "max_depth": int,                # Should be <10
        "avg_dependencies": float,       # Should be <2
        "invalid_plans": int,            # Should be 0
    }
    
    execution_metrics = {
        "task_completion_rate": float,   # Should be >0.7
        "avg_retry_count": float,        # Should be <2
        "repair_success_rate": float,    # Should be >0.5
        "total_latency_ms": int,
        "avg_task_duration_ms": int,
    }
    
    model_metrics = {
        "qwen_usage_count": int,
        "deepseek_usage_count": int,
        "mistral_usage_count": int,
        "phi_usage_count": int,
        "avg_tokens_per_task": float,    # Should be <5000
        "timeout_count": int,            # Should be <2
    }
    
    failure_metrics = {
        "syntax_errors": int,
        "import_errors": int,
        "hallucinated_apis": int,
        "test_failures": int,
        "timeouts": int,
        "total_failures": int,
    }
    
    safety_metrics = {
        "approvals_requested": int,
        "approvals_granted": int,
        "approvals_denied": int,
        "dangerous_actions": int,
        "rollbacks": int,
    }
```

### Metrics Dashboard

```html
<div class="metrics-dashboard">
  <div class="metric-card">
    <h4>Task Completion Rate</h4>
    <div class="gauge" data-value="85"></div>
    <p>85% success | 15% failures | 0 escalations</p>
  </div>
  
  <div class="metric-card">
    <h4>Average Task Latency</h4>
    <p>2,350ms (Qwen avg: 3,100ms, DeepSeek: 2,100ms)</p>
  </div>
  
  <div class="metric-card">
    <h4>Repair Success Rate</h4>
    <div class="gauge" data-value="72"></div>
    <p>72% of failures auto-repaired</p>
  </div>
  
  <div class="metric-card">
    <h4>Failure Breakdown</h4>
    <ul>
      <li>SyntaxError: 3</li>
      <li>ImportError: 2</li>
      <li>TestFailure: 1</li>
      <li>HallucinatedAPI: 2</li>
    </ul>
  </div>
  
  <div class="metric-card">
    <h4>Safety Actions</h4>
    <p>Approvals: 8 granted, 0 denied | Rollbacks: 2</p>
  </div>
</div>
```

### Health Check Endpoints

```python
# GET /health
{
  "status": "healthy",
  "components": {
    "ollama": "connected",
    "database": "healthy",
    "memory": "healthy",
    "event_bus": "healthy",
    "sandbox": "ready"
  },
  "models": {
    "phi3:mini": "ready",
    "qwen2.5-coder": "ready",
    "deepseek-coder": "ready",
    "mistral": "ready",
    "codellama": "ready"
  },
  "metrics": {
    "uptime_hours": 72,
    "runs_completed": 42,
    "avg_run_duration_ms": 3500,
    "last_successful_run": "2026-05-15T14:30:00Z"
  }
}

# GET /health/detailed
{
  "current_run": {
    "run_id": "run_42",
    "started_at": "2026-05-15T14:00:00Z",
    "active_task": "t5",
    "active_model": "qwen",
    "progress": 0.65,
    "estimated_remaining_ms": 1200
  },
  "last_5_runs": [
    { "run_id": "run_42", "status": "executing", ... },
    { "run_id": "run_41", "status": "completed", ... },
    ...
  ]
}
```

---

## LONG-RUNNING OPERATION

### Session Management

#### Starting a Long-Running Session
```python
async def start_session(original_prompt: str) -> SharedContext:
    """
    Create resumable session
    """
    session_id = generate_session_id()
    run_id = generate_run_id()
    
    shared_context = SharedContext(
        session_id=session_id,
        run_id=run_id,
        original_prompt=original_prompt,
        task_dag=[],  # Will be filled after decomposition
        artifacts_dir=f"artifacts/{run_id}",
        started_at=now(),
    )
    
    # Persist to SQLite immediately
    await memory.store_context(shared_context)
    
    # Create artifacts directory
    os.makedirs(shared_context.artifacts_dir, exist_ok=True)
    
    return shared_context
```

#### Checkpointing State
```python
async def checkpoint_session(shared_context: SharedContext):
    """
    Save session state every N minutes
    """
    checkpoint = {
        "shared_context": shared_context.model_dump(),
        "task_states": [t.model_dump() for t in shared_context.task_dag],
        "completed_tasks": [t.id for t in shared_context.task_dag if t.status == "completed"],
        "checkpoint_time": now(),
    }
    
    # Save to SQLite
    await memory.save_checkpoint(shared_context.session_id, checkpoint)
    
    # Also save to JSON for easy inspection
    with open(f"{shared_context.artifacts_dir}/checkpoint.json", "w") as f:
        json.dump(checkpoint, f, indent=2)
```

#### Resuming After Crash
```python
async def resume_session(session_id: str) -> SharedContext:
    """
    Load session and skip completed tasks
    """
    # Restore from SQLite
    shared_context = await memory.load_context(session_id)
    
    if not shared_context:
        raise SessionNotFound(session_id)
    
    # Mark resuming
    event = RuntimeEvent(
        action="session_resumed",
        session_id=session_id,
        metadata={
            "previous_progress": len([t for t in shared_context.task_dag if t.status == "completed"]),
            "total_tasks": len(shared_context.task_dag),
        }
    )
    await event_bus.emit(event)
    
    return shared_context
```

#### Session Timeout Handling
```python
async def monitor_session_timeout(session_id: str, max_duration_hours: int = 24):
    """
    Monitor long-running sessions
    """
    while True:
        shared_context = await memory.load_context(session_id)
        
        if not shared_context:
            break
        
        elapsed = (now() - shared_context.started_at).total_seconds() / 3600
        
        if elapsed > max_duration_hours:
            # Graceful shutdown
            await help_system.request_help({
                "reason": "session_timeout",
                "session_id": session_id,
                "duration_hours": elapsed,
                "message": f"Session running for {elapsed:.1f} hours. Continue or reset?"
            })
            break
        
        await asyncio.sleep(300)  # Check every 5 minutes
```

### Memory Management for Long Sessions

```python
async def maintain_memory_during_session(shared_context: SharedContext):
    """
    Prevent memory bloat during long sessions
    """
    
    # Every hour:
    while True:
        # 1. Prune session memory (keep only recent 10 tasks)
        completed_tasks = [t for t in shared_context.task_dag if t.status == "completed"]
        if len(completed_tasks) > 10:
            old_tasks = completed_tasks[:-10]
            for task in old_tasks:
                # Archive to disk
                await artifacts_store.archive(task.id)
                # Remove from memory
                shared_context.task_dag.remove(task)
        
        # 2. Compact failure memory (consolidate similar failures)
        await memory.failures.compact()
        
        # 3. Compress architecture memory (summarize decisions)
        await memory.architecture.compress()
        
        # 4. Checkpoint
        await checkpoint_session(shared_context)
        
        await asyncio.sleep(3600)  # Every hour
```

### Data Retention Policy

```yaml
retention_policy:
  artifacts:
    completed_runs:
      keep_duration_days: 30
      keep_count_minimum: 10
    failed_runs:
      keep_duration_days: 90
      keep_count_minimum: 100
  
  events:
    keep_duration_days: 7
    keep_count_minimum: 100000
    older_events: archive_to_disk
  
  database:
    checkpoint_frequency: every_5_minutes
    backup_frequency: daily
    backup_retention: 7_days
  
  memory:
    session_memory_clear: per_session
    failure_memory_retain: permanent
    architecture_memory_retain: permanent
    semantic_memory_retain: permanent
```

---

## TROUBLESHOOTING GUIDE

### Problem: "Orchestration hangs after task 3"

**Diagnosis**:
```python
# Check if model is stuck
python scripts/check_model_health.py

# Check event log
sqlite3 data/cognition.db "SELECT * FROM events WHERE run_id='run_xyz' ORDER BY timestamp DESC LIMIT 20"

# Check if task 4 is ready
python scripts/check_dag_state.py run_xyz
```

**Solutions**:
1. **Model timeout**: Increase timeout from 180s to 300s in `models/ollama_client.py`
2. **Stuck semaphore**: Check for deadlocked asyncio.Lock - restart process
3. **Invalid dependency**: Verify DAG by replaying: `python scripts/replay_run.py run_xyz`

### Problem: "Task repeatedly fails with ImportError"

**Diagnosis**:
```
FailureRecord: {
  failure_type: "import_error",
  raw_error: "No module named 'jwt'",
  attempt: 2,
  resolution_status: "escalated"
}
```

**Solution Path**:
1. **First repair attempt**: Auto-install via pip
2. **Second repair attempt**: Modify import statement
3. **Third failure**: Escalate to help system

**Fix**:
```python
# Add repair strategy in repair/repair_planner.py
repair_strategies[FailureType.IMPORT_ERROR] = [
    "pip install {package_name}",
    "check_pyproject.toml_for_package",
    "modify_import_to_relative",
]
```

### Problem: "Artifacts are corrupted after rollback"

**Diagnosis**:
```python
# Check snapshot integrity
python scripts/verify_snapshots.py run_xyz

# Compare snapshots
diff artifacts/run_xyz/snapshot_before artifacts/run_xyz/snapshot_after
```

**Prevention**:
- Snapshots must be taken BEFORE every file edit
- Rollback must restore all files atomically
- Add checksum validation after restore

### Problem: "Memory grows unbounded in long sessions"

**Diagnosis**:
```bash
# Monitor memory usage
ps aux | grep astraeus-core
watch -n 5 'ps aux | grep python | grep cognition'

# Check what's in memory
python scripts/dump_memory_stats.py
```

**Solution**:
```python
# Implement aggressive pruning
async def cleanup_memory():
    # Clear session memory every hour
    session_memory.clear()
    
    # Prune artifacts older than 12 hours
    await artifacts_store.prune(max_age_hours=12)
    
    # Compact SQLite
    sqlite3.execute("VACUUM")
```

### Problem: "DAG topological sort fails"

**Diagnosis**:
```python
# Check for cycles
python scripts/detect_dag_cycles.py run_xyz

# Dump DAG structure
python scripts/visualize_dag.py run_xyz > dag.dot
```

**Fix**:
```python
# In orchestrator/dag.py, ensure:
def validate_dag_acyclic(tasks: list[TaskNode]):
    # Use DFS to detect cycles
    visited = set()
    rec_stack = set()
    
    def has_cycle(task_id):
        visited.add(task_id)
        rec_stack.add(task_id)
        
        for dep in get_dependencies(task_id):
            if dep not in visited:
                if has_cycle(dep):
                    return True
            elif dep in rec_stack:
                return True
        
        rec_stack.remove(task_id)
        return False
    
    for task in tasks:
        if has_cycle(task.id):
            raise DAGValidationError(f"Cycle detected: {task.id}")
```

### Problem: "Approval gate never responds"

**Diagnosis**:
```
WebSocket message timeout waiting for approval_response
User clicked [Approve] but system didn't receive it
```

**Check**:
```javascript
// In frontend console
console.log("WebSocket state:", ws.readyState);
// 0=CONNECTING, 1=OPEN, 2=CLOSING, 3=CLOSED

// Check backend logs
tail -f logs/api.log | grep approval
```

**Fix**:
```python
# Add timeout for approval requests
async def wait_for_approval(approval_id: str, timeout_seconds: int = 300):
    try:
        result = await asyncio.wait_for(
            approval_queue.get(approval_id),
            timeout=timeout_seconds
        )
        return result
    except asyncio.TimeoutError:
        # Automatically deny and escalate
        help_system.request_help(
            reason="approval_timeout",
            approval_id=approval_id
        )
```

### Problem: "Help requests never get resolved"

**Root Cause**: Help system has no SLA for human response

**Mitigation**:
```python
class HelpRequestSLA:
    critical: 30_minutes    # E.g., infinite loop
    high: 2_hours          # E.g., architectural blocker
    medium: 8_hours        # E.g., unusual error
    low: 24_hours          # E.g., documentation question
    
    async def monitor_help_sla():
        while True:
            for help_req in memory.get_unresolved_help_requests():
                elapsed = now() - help_req.created_at
                
                if elapsed > get_sla(help_req.priority):
                    # Escalate or auto-terminate
                    await escalate_help_request(help_req.id)
            
            await asyncio.sleep(60)  # Check every minute
```

### Problem: "Tests pass locally but fail in CI"

**Likely Cause**: Ollama models not installed in CI environment

**Solution**:
```bash
# Add to CI setup
ollama pull phi3:mini
ollama pull qwen2.5-coder
ollama pull deepseek-coder:6.7b
ollama pull mistral
ollama pull codellama

# Or use deterministic fallbacks for CI
export COGNITION_LIVE_OLLAMA=0  # Use mocked responses
pytest tests/
```

---

## CORE EXECUTION INSTRUCTIONS

### How to Use This Specification

#### 1. **Initial Setup**
```bash
cd /path/to/astraeus-core

# Read this entire master spec first
less COGNITION_ENGINE_MASTER_SPEC.md

# Review existing code
find . -name "*.py" -type f | head -20  # See structure
ls -la contracts/ orchestrator/ models/ memory/

# Run existing tests
python -m pytest tests/ -v

# Check Ollama is running
python scripts/verify_ollama.py
```

#### 2. **Your Primary Responsibilities**

**Week 1**: Orchestration Stability
- [ ] Remove deterministic fallbacks
- [ ] Test with live Ollama inference
- [ ] Harden timeout handling
- [ ] Verify DAG topological sort is bulletproof
- [ ] Build execution replay system

**Week 2**: Repair System Maturity
- [ ] Expand failure taxonomy from real failures
- [ ] Implement repair strategies incrementally
- [ ] Build repair success evaluator
- [ ] Verify localized retries work

**Week 3**: Repository Intelligence
- [ ] Complete AST parsing
- [ ] Implement architecture invariant checking
- [ ] Build hallucination detection
- [ ] Add ProjectContext to model prompts

**Week 4**: Transaction Safety
- [ ] Complete multi-file transaction system
- [ ] Add semantic diffing
- [ ] Implement staging + validation + commit
- [ ] Build rollback confidence system

**Month 2**: Long-Running Sessions
- [ ] Implement session persistence
- [ ] Add crash recovery
- [ ] Build session learning accumulator
- [ ] Verify 8+ hour sessions work

#### 3. **Daily Development Workflow**

```bash
# 1. Execute an engineering goal
python -m cli.main run "Goal" --repo /path/to/repo

# 2. Verify substrate integrity
python -m cli.main verify

# 3. Audit recent mutations
python -m cli.main audit --verbose

# 4. Replay a past run to debug
python -m cli.main replay run_ID

# 5. Continuous Testing
python -m pytest tests/
```
python scripts/check_db_integrity.py

# Implement feature
# Write tests
python -m pytest tests/test_*.py -v

# Lint and type-check
uv run ruff check .
uv run mypy cognition_engine/ --config-file=pyproject.toml

# Test end-to-end
python -m cli.main "simple test goal" \
  --artifacts /tmp/test-artifacts \
  --data /tmp/test-data

# Review metrics
sqlite3 data/cognition.db "SELECT * FROM metrics WHERE timestamp > datetime('now', '-1 hour')"

# Before commit
git add .
git commit -m "feat: [WHAT YOU DID]"

# Verify nothing broke
python -m pytest tests/
python scripts/replay_recent_run.py  # Ensure replay works
```

#### 4. **When to Refactor vs Add**

**DO REFACTOR** if:
- Code violates hard constraints (Rule 1-10 in Execution Rules)
- Test coverage drops below 80%
- Metrics degrade

**DON'T REFACTOR** if:
- System is currently stable
- Tests are passing
- You're adding a new feature
→ Add feature first, refactor later

#### 5. **How to Add a New Repair Strategy**

```python
# File: repair/repair_planner.py

# 1. Identify new failure type from real runs
failure_types_seen = [
    FailureType.CIRCULAR_IMPORT,  # New!
]

# 2. Add to failure_types.py
class FailureType(Enum):
    CIRCULAR_IMPORT = "circular_import"  # New!

# 3. Implement repair strategy
repair_strategies[FailureType.CIRCULAR_IMPORT] = [
    "separate_module_concerns",
    "introduce_mediator_pattern",
    "refactor_imports_topologically",
]

# 4. Add corresponding repair task generator
def generate_circular_import_repair(failure: FailureRecord) -> list[TaskNode]:
    return [
        TaskNode(
            id="repair_1",
            type=TaskType.ARCHITECTURE,
            goal="Analyze circular import: " + failure.structured_context["files"],
            assigned_model=ModelType.DEEPSEEK,
            validation=["no_circular_imports"]
        ),
        # ... more tasks ...
    ]

# 5. Test with real failure
# A circular import will occur → FailureRecord created
# RepairPlanner will invoke this strategy
# Monitor: did it work?
```

#### 6. **How to Debug a Failed Orchestration**

```bash
# 1. Find the run ID
sqlite3 data/cognition.db "SELECT run_id FROM runs ORDER BY started_at DESC LIMIT 1"
# Result: run_123

# 2. Replay the exact run
python scripts/replay_run.py run_123

# 3. Check which task failed
sqlite3 data/cognition.db "SELECT id, status, error FROM tasks WHERE run_id='run_123' ORDER BY completed_at DESC"

# 4. Inspect the failure
sqlite3 data/cognition.db "SELECT * FROM failures WHERE task_id='t5'"

# 5. View the artifact
cat artifacts/run_123/t5/output.txt

# 6. Check what model was used
sqlite3 data/cognition.db "SELECT assigned_model FROM tasks WHERE id='t5'"

# 7. See the prompt that was sent
cat artifacts/run_123/t5/prompt.md

# 8. Verify validation failed correctly
python scripts/validate_output.py artifacts/run_123/t5/output.txt

# 9. Check if repair was attempted
sqlite3 data/cognition.db "SELECT * FROM repairs WHERE failed_task_id='t5'"

# 10. Identify pattern
# If this is a new failure type:
#   → Add to failure_types.py
#   → Implement repair strategy
# If this is a known failure:
#   → Check if repair strategy exists
#   → Debug why repair failed
```

#### 7. **The Golden Rules for Codex**

```
✅ DO:
  - Make one change at a time
  - Test after every change
  - Add metrics for everything
  - Keep logs append-only
  - Always take snapshots before mutations
  - Fail fast with structured errors
  - Document architectural decisions
  - Verify replay works after changes

❌ DON'T:
  - Delete or modify events/logs
  - Run multiple models concurrently
  - Mutate shared state without locks
  - Skip validation
  - Create new memory systems (use the 4 existing)
  - Add complex features without tests
  - Ignore metrics regressions
  - Make async code more complex than needed
```

#### 8. **Communication Protocol for Handoffs**

When handing off to another AI or developer:

```markdown
## Cognition Engine Status Report

### Current State
- Running since: [timestamp]
- Total runs completed: N
- Success rate: X%
- Avg run duration: T minutes

### Latest Metrics
- Task completion rate: X%
- Repair success rate: Y%
- Most common failures: [list]
- Active session: [session_id]

### Recent Work
1. [What was done]
2. [What tests pass]
3. [What metrics changed]

### Next Priority
1. [Exact next task]
2. [How to verify it works]
3. [Acceptance criteria]

### Known Issues
- [Issue 1] (Severity: HIGH|MEDIUM|LOW)
- [Issue 2]

### Database Backups
- Last backup: [timestamp]
- Location: [path]

### Questions for Next Developer
- [?]
```

---

## DEPLOYMENT CHECKLIST

Before considering system "ready for production":

```yaml
checklist:
  control-plane:
    - [ ] 100+ successful runs without corruption
    - [ ] All 10 execution rules verified
    - [ ] DAG topological sort tested extensively
    - [ ] Timeouts handled gracefully
    - [ ] Replay works for all runs
  
  reliability:
    - [ ] 90%+ task completion rate
    - [ ] 70%+ repair success rate
    - [ ] All failures captured as FailureRecords
    - [ ] Localized retries working
    - [ ] Snapshots + rollback bulletproof
  
  safety:
    - [ ] No file mutations without approval
    - [ ] Audit log comprehensive
    - [ ] Permissions enforced
    - [ ] Approval gates working
    - [ ] No privilege escalation
  
  observability:
    - [ ] Event log complete and verifiable
    - [ ] Metrics dashboard functional
    - [ ] Health checks passing
    - [ ] All errors categorized
    - [ ] Debugging tools available
  
  long_running:
    - [ ] Sessions survive 8+ hours
    - [ ] Crashes recoverable
    - [ ] Memory stays bounded
    - [ ] Checkpoints working
    - [ ] No memory leaks detected
  
  testing:
    - [ ] 90%+ code coverage
    - [ ] All tests passing
    - [ ] Stress tests passing
    - [ ] Failure injection tests passing
    - [ ] Replay tests passing
  
  documentation:
    - [ ] Architecture documented
    - [ ] APIs documented
    - [ ] Troubleshooting guide written
    - [ ] Operations manual ready
    - [ ] Developer onboarding complete
```

---

## APPENDIX: Quick Reference

### Critical Files Reference
```
contracts/models.py          ← All data models live here
orchestrator/control_plane.py       ← Main entry point
models/base_adapter.py       ← How models work
memory/memory_system.py      ← Memory coordination
cli/main.py                  ← CLI interface
frontend-console/index.html  ← UI
scripts/verify_ollama.py     ← Health check
tests/test_e2e_control-plane.py ← Integration test
```

### Key Commands
```bash
# Start Ollama
ollama serve

# Start frontend
cd frontend-console && python -m http.server 8000

# Start API
python -m api.main

# Run CLI
python -m cli.main "your engineering goal here"

# Test everything
python -m pytest tests/ -v --cov

# Debug a run
python scripts/replay_run.py run_xyz
python scripts/visualize_dag.py run_xyz
```

### Common Tasks
```bash
# Add new repair strategy
# 1. Edit repair/repair_planner.py
# 2. Add FailureType to validator/failure_types.py
# 3. Write tests
# 4. Deploy

# Troubleshoot hangs
python scripts/check_model_health.py
ps aux | grep ollama

# Verify replay
python scripts/verify_all_replays.py

# Monitor long session
watch -n 5 'sqlite3 data/cognition.db "SELECT run_id, COUNT(*) FROM tasks WHERE status=\"completed\" GROUP BY run_id"'
```

---

## FINAL SUBSTRATE GOAL

The actual final system is a **deterministic, repository-grounded, temporally replayable, semantically verified, safely autonomous engineering execution substrate.** It is built to operate at the very core of autonomous engineering systems, providing the reliable foundation upon which higher-level cognition can execute.

---

## END OF MASTER SPECIFICATION

**Next Steps for Codex**:
1. Read this entire document
2. Review existing code structure
3. Run existing test suite
4. Start with "Orchestration Stability" week 1 tasks
5. Follow the verification checklist daily
6. Report metrics after every run

**Questions?** Review the Troubleshooting Guide or Architecture section above.

**Ready to build?** Start here: `python -m pytest tests/test_core_components.py -v`
