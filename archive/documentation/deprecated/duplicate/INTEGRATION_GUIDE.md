# Multi-LLM Cognition Engine: Integration Guide

## System Overview

You now have a complete multi-LLM control-plane system that:

1. **Routes tasks to specialized models** based on type
2. **Manages separate memory subsystems** by purpose
3. **Decomposes complex requests** using phi3:mini planner
4. **Executes in a read-only sandbox** initially
5. **Integrates with repo-analyzer** for code intelligence

## Architecture Layers

### Layer 1: Models (Unified Interface)
```python
from models import ModelRegistry, ModelType

adapter = ModelRegistry.get_adapter(ModelType.QWEN)
result = adapter.generate(prompt="...", context={})
```

All models expose the same interface:
- `generate(prompt, context, **kwargs) → GenerationResult`
- Consistent error handling
- Performance measurement (tokens, latency)

### Layer 2: Orchestration
```python
from orchestrator.control_plane import ControlPlane

engine = ControlPlane()
result = engine.process_request(
    user_request="Analyze for duplicated modules",
    repo_context={...}
)
```

The engine handles:
- Task type inference
- Model routing (TaskRouter)
- Task decomposition (Planner)
- Session management
- Task graph execution

### Layer 3: Memory (Separated by Type)
```python
from memory import MemorySystem
from contracts.models import MemoryType

memory = MemorySystem()

# Store semantic knowledge
memory.semantic.store(memory_record)

# Record failures
memory.failures.record_failure(error_type, stack_trace, context)

# Record architecture decisions
memory.architecture.record_decision(title, context, decision, consequences)

# Active task context
memory.session.set_context("repo", repo_path)
```

Memory types:
- **Semantic**: Long-term code/repo knowledge (searchable via embeddings)
- **Session**: Active task context (FIFO, cleared per session)
- **Failure**: Debugging history (for future reference)
- **Architecture**: ADRs and design decisions

### Layer 4: Execution
```python
from runtime.mutation_sandbox import MutationSandbox

sandbox = MutationSandbox(repo_path, journal=journal)
result = sandbox.apply_mutation(diff_plan)
```

The execution layer provides:
- **Mutation Safety**: Every change is journaled and hash-chained.
- **Transactional Integrity**: Multi-file updates are staged and verified.
- **Granular Rollback**: State recovery at the filesystem level.
- **Command Governance**: Risk-based gating and auditing.

## Task Routing

The router automatically assigns tasks:

| Task Type | Primary Model | Why |
|-----------|--------------|-----|
| Code Generation | Qwen | Best for authoring |
| Debugging | DeepSeek | Specializes in analysis |
| Documentation | Mistral | Optimized for clarity |
| Architecture | Qwen | Design expertise |
| Extraction | Qwen | Code understanding |
| Planning | Phi3 | Task decomposition |
| Synthesis | CodeLlama | Alternative generation |

## Integration with repo-analyzer

The astraeus-core complements repo-analyzer:

**repo-analyzer**: Machine-readable code structure
- File scanning and parsing
- AST analysis (with tree-sitter)
- Semantic embeddings
- Code search

**astraeus-core**: AI-driven analysis
- Task routing and planning
- Multi-model control-plane
- Memory and learning
- Architecture decisions

Combined use:
```python
from repo_analyzer.analyzers.repository_analyzer import RepositoryAnalyzer
from cognition_engine import ControlPlane

# Step 1: Analyze repo structure
analyzer = RepositoryAnalyzer()
repo_analysis = analyzer.analyze(repo_path, "my-project")

# Step 2: Ask AI engine to analyze
engine = ControlPlane()
result = engine.process_request(
    user_request="Suggest module extraction based on this repo",
    repo_context={
        "analysis": repo_analysis,
        "path": repo_path
    }
)
```

## Execution Flow Example

**User Request**: "Analyze this repo and suggest reusable modules"

1. **Planner (phi3:mini)** decomposes:
   - Step 1: Scan repository structure
   - Step 2: Build semantic index
   - Step 3: Analyze architecture
   - Step 4: Generate extraction recommendations

2. **Router** assigns:
   - Steps 1-3 → Qwen (code analysis)
   - Step 4 → Mistral (documentation generation)

3. **Session Memory** tracks:
   - Repo path
   - Intermediate results
   - Working context

4. **Execution Sandbox** runs:
   - Directory traversal
   - File content reading
   - Git history inspection (read-only)

5. **Memory System** stores:
   - Semantic embeddings of analyzed code
   - Architecture decisions made
   - Any failures encountered

## Implementation Roadmap & Status

The project has transitioned from its prototype phase to a completion-oriented execution structure. Implementation is now governed by the **Master TODO Hierarchy**.

### 12-Domain Execution Roadmap
- **Phase A-B**: Foundation, Sanitization, Runtime, and Safety Substrate.
- **Phase C-D**: Event Sourcing, Replay, and Repository Cognition.
- **Phase E-F**: Semantic Verification and Temporal Cognition.
- **Phase G-H**: Concurrency, Distributed Governance, and Zero Trust.
- **Phase I-J**: Observability, Operations, and CI/CD.
- **Phase K-L**: Adaptive Repair, Intelligence, and Productionization.

For the full detailed hierarchy and status, see **[TODO.md](./TODO.md)** and **[IMPLEMENTATION_STATUS.md](./IMPLEMENTATION_STATUS.md)**.

## Design Principles

1. **No Monolithic Agent**
   - Each model has a specific role (not catch-all)
   - Router decides which model for each task
   - Planner orchestrates using phi3:mini (lightweight)

2. **Memory by Purpose**
   - Not one giant vector DB
   - Semantic: code/repo knowledge
   - Session: active task context
   - Failure: debugging reference
   - Architecture: system decisions

3. **Deterministic & Grounded**
   - Every operation must be lineage-traceable
   - All cognition is repository-grounded
   - All replay must be deterministic

4. **Separation of Concerns**
   - Models: inference only
   - Orchestrator: routing + planning
   - Memory: storage + retrieval
   - Runtime: safe execution

5. **M3 8GB Optimizations**
   - Single active inference model
   - Lazy loading of models
   - Batch embeddings
   - Per-session memory clearing

## File Structure

```
astraeus-core/
├── contracts/
│   ├── __init__.py
│   └── models.py (TaskType, MemoryType, Task, etc.)
├── models/
│   ├── __init__.py (ModelRegistry)
│   ├── base_adapter.py
│   ├── qwen_adapter.py
│   ├── deepseek_adapter.py
│   ├── mistral_adapter.py
│   ├── codellama_adapter.py
│   └── phi_adapter.py
├── orchestrator/
│   ├── __init__.py
│   ├── engine.py (ControlPlane - main entry point)
│   ├── router.py (TaskRouter)
│   ├── planner.py (Planner - task decomposition)
│   ├── session_manager.py (SessionManager)
│   └── task_graph.py (TaskGraph - DAG)
├── memory/
│   ├── __init__.py (MemorySystem)
│   ├── semantic_store.py (ChromaDB integration)
│   ├── session_memory.py (Active context)
│   ├── failure_memory.py (Debugging history)
│   └── architecture_memory.py (ADRs)
├── runtime/
│   ├── __init__.py
│   └── sandbox.py (ExecutionSandbox - read-only)
├── agents/
│   └── __init__.py (Scaffolding for v2)
├── __init__.py
├── README.md
├── pyproject.toml
└── INTEGRATION_GUIDE.md (this file)
```

## Running the System

### Initialization
```python
from cognition_engine import ControlPlane

engine = ControlPlane()
```

### Single Request
```python
result = engine.process_request(
    user_request="Generate a REST API for user authentication",
    repo_context={"path": "/my/repo"}
)

print(result["response"])
print(f"Model: {result['assigned_model']}")
print(f"Tokens: {result['tokens_used']}")
```

### Multi-Step Workflow
```python
# The engine handles decomposition internally
result = engine.process_request(
    user_request="Analyze this repo for duplicated utilities and suggest extraction",
    repo_context={...}
)
```

### Memory Queries
```python
from contracts.models import MemoryType

# Find similar failures
failures = engine.memory.failures.find_similar_failures(
    error_type="AssertionError",
    stack_trace="..."
)

# Get architecture decisions
decisions = engine.memory.architecture.list_decisions()

# Record new decision
adr_id = engine.memory.architecture.record_decision(
    title="Use event-driven architecture",
    context="Growing complexity...",
    decision="Implement CQRS pattern",
    consequences="Eventual consistency required"
)
```

## Constraints & Guarantees

✓ **Single active model**: Only one model running at a time  
✓ **Lazy loading**: Models loaded on first use  
✓ **Memory isolation**: Separate subsystems by type  
✓ **Read-only execution**: No file mutations initially  
✓ **Session scope**: Memory cleared per session (except semantic/architecture)  
✓ **M3-optimized**: Designed for 8GB RAM machine  

## What's Next

This is your foundation for:
1. **Self-analyzing engineering infrastructure**
2. **Accumulated engineering wisdom** (via memory)
3. **Reusable cognition modules** (via extraction)
4. **Self-improving architecture** (via learning)

Not a chatbot. A **real engineering cognition system**.
