# Distributed Cognition Platform - Definitive Engineering Roadmap

Date: 2026-05-14
Status: canonical execution plan

## 0. Definition

Build a local-first distributed cognition platform for engineering workflows.

Not:

- AGI.
- General assistant.
- Voice companion.
- Chatbot wrapper.
- Autonomous OS.

Core loop:

```text
goal -> decomposition -> control-plane -> execution -> validation -> recovery -> memory -> delivery
```

Initial domain:

```text
software engineering automation
```

Primary invariant:

```text
deterministic infrastructure around probabilistic models
```

## 1. Architectural Principles

1. Orchestration beats model intelligence.
2. Constrained cognition beats unrestricted autonomy.
3. Every action is observable, replayable, and recoverable.
4. Safety precedes autonomy.
5. Single-node reliability precedes distribution.
6. Raw LLM output is never trusted.
7. Memory is scoped retrieval, not accumulation.
8. Every tool action has an explicit permission level.
9. Research forks are pattern sources, not platform dependencies.

## 2. Initial Repository Shape

Start as a monorepo.

```text
cognition-core/
├── planner/
│   ├── decomposer.py
│   ├── schemas.py
│   └── prompts/
├── orchestrator/
│   ├── router.py
│   ├── queue.py
│   ├── dag.py
│   └── retry.py
├── validator/
│   ├── syntax.py
│   ├── critic.py
│   └── contracts.py
├── memory/
│   ├── store.py
│   ├── retrieval.py
│   └── schema.sql
├── sandbox/
│   ├── runtime.py
│   ├── policies.py
│   └── dockerfiles/
├── shared_context/
│   ├── state.py
│   └── artifacts.py
├── events/
│   ├── event_bus.py
│   └── schema.py
├── tools/
│   ├── registry.py
│   ├── permissions.py
│   └── specs.py
├── models/
│   ├── ollama.py
│   └── protocol.py
├── cli/
│   └── main.py
├── artifacts/
├── tests/
├── pyproject.toml
├── Taskfile.yml
└── README.md
```

Do not split repos until:

- Interfaces are stable.
- At least two consumers exist.
- Test boundaries are clear.
- A module can be versioned independently.

Future repos, only after split pressure is proven:

```text
cognition-core
tool-governor
memory-engine
context-compaction
sandbox-runtime
repo-indexer
research-agents
frontend-console
vscode-extension
plugin-runtime
observability-stack
```

## 3. Tech Stack

Phase 1:

```yaml
language:
  - Python 3.12

runtime:
  - asyncio

models:
  - Ollama
  - mistral
  - deepseek-coder
  - phi

storage:
  - SQLite
  - local filesystem

sandbox:
  - Docker

quality:
  - Ruff
  - MyPy strict
  - Pytest
```

Phase 3:

```yaml
api:
  - FastAPI
  - WebSocket streaming

frontend:
  - Next.js
```

Later only after measured need:

```yaml
advanced:
  - Tauri
  - Temporal
  - Rust scheduler
  - plugin runtime
```

Forbidden in MVP:

- Kubernetes.
- Kafka.
- Redis clusters.
- Temporal.
- Microservices.
- Browser automation as core runtime.
- Autonomous network access.

## 4. MVP Scope

The MVP is only:

```text
User goal
  -> task decomposer
  -> execution DAG
  -> model router
  -> parallel execution
  -> validation
  -> structured artifacts
```

Anything outside this is backlog.

## 5. Phase M1 - Vertical Slice

Duration: weeks 1-3.

Goal:

```text
prove control-plane works
```

### M1.1 Local Inference

Tasks:

- Install Ollama.
- Pull `mistral`, `deepseek-coder`, and `phi`.
- Verify HTTP inference.
- Add a minimal model protocol wrapper in `models/protocol.py`.
- Add Ollama adapter in `models/ollama.py`.

Output:

```text
local model call returns deterministic structured text under timeout
```

### M1.2 Planner Engine

File:

```text
planner/decomposer.py
```

Responsibilities:

- Convert goal into atomic tasks.
- Emit dependencies.
- Assign task type.
- Emit validation requirements.
- Emit model hint.

Required output:

```json
{
  "tasks": [
    {
      "id": "task_1",
      "type": "code_generation",
      "goal": "implement auth middleware",
      "depends_on": [],
      "validation": ["syntax"],
      "assigned_model": "deepseek-coder"
    }
  ]
}
```

Rules:

- JSON only.
- Validate against schema.
- Retry malformed planner output once.
- If still invalid, fail the run.

### M1.3 Shared Context

File:

```text
shared_context/state.py
```

State owns:

- goals.
- tasks.
- DAG.
- artifacts.
- failures.
- model outputs.
- validation results.

Invariant:

```text
shared_context is the single source of truth for one run
```

### M1.4 Model Router

File:

```text
orchestrator/router.py
```

Initial policy:

```yaml
planning: mistral
coding: deepseek-coder
summarization: phi
critique: mistral
```

Responsibilities:

- Map task type to model.
- Enforce timeout.
- Return structured result.
- Surface malformed output as typed failure.

### M1.5 Execution Queue

Files:

```text
orchestrator/dag.py
orchestrator/queue.py
```

Responsibilities:

- Topologically schedule tasks.
- Run dependency-free tasks in parallel.
- Use `asyncio.gather()`.
- Stop downstream tasks if dependency fails.
- Persist task status transitions.

Task states:

```text
pending -> ready -> running -> succeeded
pending -> ready -> running -> failed
failed -> retrying -> running
blocked
```

### M1.6 Artifact System

File:

```text
shared_context/artifacts.py
```

Structure:

```text
artifacts/
└── run_<timestamp>/
    ├── run.json
    ├── events.jsonl
    └── tasks/
        └── task_001/
            ├── input.json
            ├── output.md
            ├── validation.json
            └── logs.txt
```

Rules:

- Never overwrite run artifacts.
- Every task stores raw model output and parsed output.
- Every validation failure stores exact reason.

### M1.7 CLI Runner

File:

```text
cli/main.py
```

Usage:

```bash
python -m cli.main "build a REST API skeleton with tests"
```

M1 gate:

```text
One engineering goal decomposes into a DAG, routes to multiple local models, executes with dependency order, validates basic outputs, and writes useful artifacts.
```

If gate fails:

```text
stop all downstream work and fix decomposition
```

## 6. Phase M2 - Memory and Safety

Duration: weeks 4-6.

Goal:

```text
make execution observable, recoverable, and safe
```

### M2.1 Event Bus

File:

```text
events/event_bus.py
```

Event schema:

```json
{
  "timestamp": "2026-05-14T00:00:00Z",
  "run_id": "run_001",
  "task_id": "task_001",
  "action": "task_started",
  "result": "ok",
  "latency_ms": 12,
  "metadata": {}
}
```

Required events:

- run_started.
- run_finished.
- task_created.
- task_ready.
- task_started.
- model_called.
- model_failed.
- validation_failed.
- task_succeeded.
- task_failed.
- retry_scheduled.
- approval_required.
- sandbox_started.
- sandbox_finished.

### M2.2 SQLite Memory

File:

```text
memory/store.py
```

Store only:

- goals.
- decompositions.
- final artifacts.
- validation outcomes.
- failure lessons.
- architecture summaries.

Do not store:

- every intermediate token.
- full chat transcripts by default.
- untrusted remote content as memory.
- secrets.

### M2.3 Retrieval

File:

```text
memory/retrieval.py
```

Initial retrieval:

- keyword search over SQLite FTS.
- top-k successful prior solutions.
- top-k failure lessons.

Semantic embeddings are optional after keyword retrieval works.

Retrieval invariant:

```text
retrieval suggests context; it does not inject blindly
```

### M2.4 Critic and Validation

Files:

```text
validator/syntax.py
validator/critic.py
```

Pipeline:

```text
generator -> syntax validator -> critic model -> retry
```

Max retries:

```text
2
```

Required validators:

- JSON schema validation.
- Python syntax validation.
- Markdown artifact presence.
- Basic command dry-run where sandboxed.

### M2.5 Sandbox Runtime

File:

```text
sandbox/runtime.py
```

Docker constraints:

```bash
--network none
--memory 512m
--cpus 1
--read-only
```

Mutable paths must be explicit mounted temp dirs.

### M2.6 Approval System

Files:

```text
tools/permissions.py
sandbox/policies.py
```

Permission levels:

```text
none
read_only
write
execute
network
dangerous
```

Require approval for:

- deletion.
- installs.
- shell execution outside sandbox.
- network access.
- writes outside artifact directory.
- credential access.

M2 gate:

```text
The system survives malformed model output, bad generated code, failed validation, and retry exhaustion without uncontrolled side effects.
```

## 7. Phase M3 - Interface

Duration: weeks 7-9.

Goal:

```text
make control-plane inspectable and usable
```

### M3.1 FastAPI Backend

Endpoints:

```yaml
POST /run:
  starts run

GET /status/{run_id}:
  returns DAG and task states

GET /artifact/{run_id}/{task_id}:
  returns task artifact

WS /stream/{run_id}:
  streams events
```

### M3.2 Frontend Console

Layout:

```text
left: goal input and run summary
right: DAG, task status, approvals, artifact viewer
```

Required views:

- current run graph.
- task details.
- approval queue.
- artifact browser.
- event log.

### M3.3 Live Execution Graph

Display states:

- waiting.
- running.
- succeeded.
- failed.
- retrying.
- blocked.
- approval_required.

### M3.4 Repo Indexer

Files:

```text
repo_indexer/
├── scan.py
├── symbols.py
└── summary.py
```

Purpose:

- Index codebase structure.
- Generate architecture summaries.
- Provide project-aware generation context.

M3 gate:

```text
A user can run a project-aware engineering task through the UI, inspect execution, approve dangerous steps, and retrieve artifacts.
```

## 8. Phase M4 - Developer Workflow Integration

Duration: months 3-4.

Goal:

```text
integrate into actual engineering workflows
```

Features:

- Multi-file edit planning.
- Test generation and execution.
- CI failure diagnosis.
- Architecture map generation.
- VSCode extension.
- Sandboxed terminal integration.

Most important capability:

```text
repository-aware code generation with validation
```

M4 gate:

```text
The system successfully modifies a real repository, runs tests in sandbox, captures failures, retries once or twice, and emits a reviewable patch.
```

## 9. Phase M5 - Platformization

Duration: month 5+.

Only start if M4 has real usage.

Allowed:

- Plugin manifests.
- Tool-governor extraction.
- Memory-engine extraction.
- Context-compaction extraction.
- Observability stack.
- Rust scheduler.
- Temporal for durable workflows.

Still forbidden unless measured:

- Kubernetes.
- distributed workers.
- multi-tenant SaaS.
- always-on autonomous browsing.

M5 gate:

```text
At least three workflows use the same stable core, and one module has two independent consumers.
```

## 10. OpenHuman Extraction Boundary

OpenHuman fork is GPLv3 and must not be copied into permissive modules.

Allowed:

- architecture study.
- behavior analysis.
- clean-room reimplementation.
- fixture-inspired tests if legally acceptable.

Priority:

1. `tokenjuice/`: output compaction strategies, token budgets, pass-through safety.
2. `tools/`: permission model, tool schema, result caps.
3. `agent/harness/`: bounded tool-loop state machine, subagent definition schema.
4. `memory/tree/`: study only until there is a real memory consumer.

Forbidden:

- direct code transplantation into non-GPL modules.
- OpenHuman runtime adoption.
- UI, voice, mascot, billing, channel provider extraction.

## 11. Module Boundaries

Planner:

- Owns decomposition.
- Emits validated DAG.
- Does not execute.

Orchestrator:

- Owns scheduling and retries.
- Does not generate content.
- Does not mutate files directly.

Model router:

- Owns model selection and timeouts.
- Does not parse domain artifacts beyond schema validation.

Validator:

- Owns correctness checks.
- Can reject output.
- Cannot silently fix output without recording event.

Sandbox:

- Owns execution isolation.
- Cannot access network unless approved.
- Cannot write outside mounted dirs.

Memory:

- Owns durable lessons and retrieval.
- Does not blindly inject context.
- Does not store secrets.

Tools:

- Own capability metadata.
- Must declare permission level.
- Must return structured results.

Events:

- Own append-only run history.
- No component bypasses event emission for state transitions.

## 12. Learning Order

Required before implementation:

1. Python asyncio task scheduling.
2. JSON schema validation.
3. DAG topological scheduling.
4. SQLite schema design.
5. Docker sandboxing.
6. Ollama HTTP API.

Required before M3:

1. FastAPI streaming.
2. WebSocket event fanout.
3. Basic frontend state management.

Required before M5:

1. durable workflow semantics.
2. plugin capability design.
3. trace/observability design.
4. memory invalidation and compaction.

## 13. Anti-Failure Constraints

Overengineering:

- One phase at a time.
- No feature before current gate passes.

Bad decomposition:

- Planner output is the highest-priority quality surface.
- If decomposition is bad, improve planner prompts/schemas before adding execution features.

Hallucination loops:

- Max two retries.
- Validation failure is terminal after retry exhaustion.
- Every retry records cause.

Context explosion:

- Top-k retrieval only.
- Compaction before model history.
- Do not store everything.

Unsafe autonomy:

- Dangerous actions require approval.
- Shell/network default denied.
- Sandbox first.

Premature distribution:

- Single machine first.
- SQLite first.
- Filesystem artifacts first.

Uninspectable behavior:

- Event log required for every run.
- Raw and parsed model outputs stored.
- Validation reason stored.

## 14. Deployment Path

M1:

```text
local CLI only
```

M2:

```text
local CLI + Docker sandbox + SQLite
```

M3:

```text
local FastAPI + Next.js console
```

M4:

```text
developer workstation integration
```

M5:

```text
single-node service deployment
```

Only after M5:

```text
multi-worker or distributed deployment
```

## 15. Success Criteria

Capstone-grade success:

- Takes a software engineering goal.
- Produces dependency-aware task graph.
- Routes tasks to local models.
- Executes independent tasks in parallel.
- Validates outputs.
- Handles malformed outputs safely.
- Persists artifacts and events.
- Requires approval for risky actions.
- Demonstrates repository-aware patch generation.

Research-grade success:

- Replays runs.
- Compares model routing policies.
- Measures decomposition quality.
- Measures retry/failure rates.
- Quantifies context compaction impact.

Platform-grade success:

- Stable tool contract.
- Stable memory interface.
- Stable sandbox interface.
- Multiple workflows configured from the same core.

## 16. Final Rule Set

1. Never add a feature before the current gate passes.
2. Never scale unreliable execution.
3. Never optimize before measurement.
4. Never distribute before single-node stability.
5. Never trust raw LLM output.
6. Every autonomous action must be observable.
7. Every dangerous action must be approved.
8. Every extracted module must survive deletion of its research fork.

## 17. Endgame

The endgame is an engineering cognition substrate:

- reusable control-plane.
- reusable memory.
- reusable tooling.
- reusable sandbox execution.
- reusable validation infrastructure.

Future engineering workflows should become configuration over stable primitives, not new runtime reinvention.
