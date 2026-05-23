# Repair Intelligence Research

## 1. Repair Ranking
To optimize the repair loop managed by `RepairPlanner`, the system must introduce historical Repair Ranking.
- **Mechanism:** Query `FailureMemory` for past attempts at mutating similar `ArchitecturalNode` patterns. Rank proposed repairs based on the historical success rate of their respective code structures.
- **Safety:** Deterministic sorting ensures the highest probability fix is executed first, preventing random hallucination loops.

## 2. Failure Memory & Patch Confidence Estimation
- **Failure Memory:** Expands `FailureMemory` to store exact semantic diffs of failed patches alongside their error traces.
- **Patch Confidence Estimation:** Before applying a patch, the Verifier model cross-references the proposed diff against the `FailureMemory`. If a similar semantic patch has failed before, the confidence score drops, and the patch is rejected pre-execution.
- **Mutation Confidence Scoring:** Assign a heuristic confidence score based on test coverage density of the affected nodes and historical instability (`ArchitecturalTemperature`).

## 3. Semantic Diff Validation
- **Current Issue:** LLMs often hallucinate surrounding context when generating diffs.
- **Improvement:** Introduce a strict AST-diff validator. The patch is applied to an isolated, in-memory representation of the AST. If the patch breaks syntactical boundaries (e.g., mismatched brackets, unresolved imports) without even running tests, it is instantly rejected.

## 4. Topology-Aware Repair
- **Concept:** Repairs cannot be isolated to a single file if `SideEffectType` indicates a cross-boundary dependency.
- **Implementation:** When `RepairPlanner` encounters a failure in node A, it queries `SemanticTopologyEngine` for all nodes that depend on A. The repair context explicitly includes the interfaces of dependent nodes to ensure the repair does not cascade failures into adjacent systems.
