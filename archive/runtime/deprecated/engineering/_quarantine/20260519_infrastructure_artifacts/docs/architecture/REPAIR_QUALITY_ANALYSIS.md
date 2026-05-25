# Repair Quality Analysis & Confidence Model

## 1. Repair Scoring
Repair quality is no longer binary (`SUCCESS`/`FAILED`). It is scored based on the **Repair Effectiveness Index (REI)**:
- **Reductivity:** Does the repair reduce the complexity (LOC, cyclomatic complexity) of the failing block?
- **Stability:** Does the repair affect high `ArchitecturalTemperature` modules?
- **Precision:** What is the ratio of "Intended AST Changes" to "Unintended Side Effects"?

## 2. Failure Memory & Patch Ranking
Astraeus maintains a "Graveyard of Failed Patches." 
- **Patch Ranking:** When a failure occurs, the `RepairPlanner` generates N candidate patches. Each candidate is ranked by the `Semantic Diff Engine` against the Graveyard.
- **Divergence Requirement:** Candidates that are semantically similar to failed patches are ranked lower. Candidates that introduce novel, verifiable patterns are ranked higher.

## 3. Repair Confidence Estimation
Before a repair is executed, the system calculates a **Repair Confidence Score (RCS)**:
- **Scope Confidence:** Is the failure localized to a single module? (High Confidence)
- **Dependency Confidence:** Does the module have few downstream consumers?
- **Historical Confidence:** Have repairs in this module succeeded historically?

## 4. Topology-Aware Verification
- **Global Impact Analysis:** A repair in `orchestrator` is verified against the `runtime` and `api` layers. If the repair changes the "Semantic Topology" (the way information flows between these layers), the RCS is penalized by 50%.
- **Invariant Strengthening:** A high-quality repair doesn't just "fix the bug"—it introduces a new `InvariantDefinition` to the `invariants.yaml` to ensure the bug never returns.

## 5. Mutation Confidence Model (Summary)
The Confidence Model acts as the final arbiter for the `CognitionEngine`. It aggregates data from the Verification Pipelines and Repair Analysis to decide if the engine is "hallucinating" or "reasoning."

| Score | Meaning | Action |
|-------|---------|--------|
| 0.9 - 1.0 | High Confidence | Immediate execution. |
| 0.7 - 0.9 | Moderate Confidence | Execution with Snapshot + Revert capability enabled. |
| 0.4 - 0.7 | Low Confidence | Mandatory human intervention (Help Request). |
| < 0.4 | Hallucination Detected | Immediate abort and failure memory recording. |
