# Cognition Engine Roadmap: Advisory AI Layer

This roadmap details the granular execution plan for the Advisory AI Layer (PhoenixOS Stage 9 / Phase F7), based on the May 2026 status report.

## Phase B: Real Self-Repair (Weeks 2-3)
*   **B1: Failure Taxonomy**: Expand `failure_types.py` to include `CIRCULAR_IMPORT`, `DEPENDENCY_CONFLICT`, etc.
*   **B2: Repair Evaluator**: Build `repair/evaluator.py` to classify outcomes (SUCCESS, REGRESSED, FAILED).
*   **B3: Differential Validation**: Implement output hashing in `orchestrator/dag.py` to skip redundant downstream reruns.
*   **B4: Rollback Risk Assessment**: Log risk levels (LOW, MEDIUM, DANGEROUS) to the event bus before mutations.
*   **B5: Semantic Repair Memory**: Query SQLite for past successful repair strategies for similar failure contexts.

## Phase C: Repository Intelligence (Weeks 4-5)
*   **C1: AST Parsing**: Complete `repo_indexer/ast_parser.py` to extract symbols and build semantic embeddings.
*   **C2: Architecture Invariants**: Implement a checker in `repo_indexer/architecture_rules.py` to prevent import cycles.
*   **C3: ProjectContext Injection**: Inject real repository context (signatures, constraints) into model prompts.
*   **C4: Semantic Diffing**: Compute behavioral diffs (e.g., signature changes) after code generation tasks.

## Phase D: Safe Autonomous Mutation (Month 2, Weeks 1-2)
*   **D1: Sandbox Hardening**: Implement resource limits (CPU/Memory/FD) and path allowlisting in `sandbox/`.
*   **D2: Command Risk Classifier**: Classify shell commands (SAFE to DESTRUCTIVE) before execution.
*   **D3: Filesystem Journaling**: Log all file operations to a SQLite journal for precise rollback.
*   **D4: Approval Simulation**: Dry-run DANGEROUS commands and show impact summaries for human approval.

## Phase E: Long-Running Sessions (Month 2, Weeks 3-4)
*   **E1: Crash Recovery**: Implement session resumption that resets "running" tasks to "pending" on restart.
*   **E2: Learning Accumulator**: Extract successful repair strategies and store in architecture memory.
*   **E3: Context Compression**: Summarize or archive old task history to keep model prompts within limits.
*   **E4: Guidance Injection**: Allow mid-run human guidance to be injected into the session state.

## Phase F: Polish & Production Readiness (Month 3)
*   **F1: UI Verification**: Verify all 6 WebSocket message types in the frontend console.
*   **F2: Keyboard Shortcuts**: Implement listed shortcuts for execution control.
*   **F3: Metrics Dashboard**: Wire the WebSocket to emit real-time task performance data.
*   **F4: Strict Type Safety**: Complete full MyPy strict mode coverage.
*   **F5: Stress Testing**: Execute deployment checklist (100+ runs, 8+ hour stability, failure injection).

## Metrics Targets
| Phase | Task Completion Rate | Repair Success Rate | Avg Latency |
| :--- | :--- | :--- | :--- |
| **A (Current)** | >70% | untested | <3000ms |
| **B** | >75% | >50% | <3000ms |
| **C** | >80% | >60% | <3500ms |
| **D** | >85% | >65% | <4000ms |
| **E** | >90% | >70% | <4000ms |
| **F** | >90% | >70% | <3000ms |
