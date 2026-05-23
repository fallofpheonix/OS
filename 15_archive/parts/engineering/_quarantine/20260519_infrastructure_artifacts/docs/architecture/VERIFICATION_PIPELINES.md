# Verification Pipelines

## 1. Multi-Stage Verification
Astraeus mutations must pass through a multi-stage verification pipeline before being committed to the codebase. This ensures that pass-by-syntax doesn't result in fail-by-logic.

### Stage 1: Deterministic Syntax & AST Validation
- **Goal:** Ensure code is runnable and semantically valid at the AST level.
- **Tools:** `ast.parse()`, `ruff`, `mypy`.
- **Invariants:** No syntax errors, no obvious type-safety violations in the mutated scope.

### Stage 2: Structural Invariant Validation
- **Goal:** Prevent architectural rot and boundary violations.
- **Tools:** `InvariantEngine`, `ArchitectureGraph`.
- **Invariants:** No new cyclic dependencies, no forbidden imports, no criticality-tier violations.

### Stage 3: Semantic Regression Detection
- **Goal:** Ensure the mutation did not break existing behavior.
- **Mechanism:** Run the `Semantic Diff Engine` against the `FailureMemory`. If the current mutation semantically matches a previously failed repair attempt, it is rejected.
- **Behavioral Diffing:** Verify that only the intended logic blocks were altered.

### Stage 4: Replay & State Validation
- **Goal:** Verify system-state consistency.
- **Mechanism:** Execute the task replay. Compare the `RunState` (session memory) before and after. If the mutation caused unexpected state divergence in non-targeted modules, it is flagged as a regression.

## 2. Dependency Validation
- Every mutation trigger a recursive dependency check via the `TaskGraph`. 
- If `Module A` is mutated, all downstream tasks (tasks that depend on A) are re-validated, even if they were previously marked as `SUCCEEDED`.

## 3. Mutation Confidence Scoring
Each stage provides a component score. The final **Mutation Confidence Score (MCS)** is calculated as:
`MCS = (Syntax_Pass * 0.2) + (Invariant_Pass * 0.3) + (Semantic_Diff_Score * 0.3) + (Test_Success * 0.2)`
- **MCS > 0.9:** Auto-commit.
- **0.7 < MCS < 0.9:** Warning - requires "Critic" review.
- **MCS < 0.7:** Reject & Trigger Repair.
