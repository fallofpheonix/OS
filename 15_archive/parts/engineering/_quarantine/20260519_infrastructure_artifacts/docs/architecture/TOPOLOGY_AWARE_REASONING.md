# Topology-Aware Reasoning

## 1. Introduction
Topology-aware reasoning elevates the LLM from operating on raw text strings to reasoning over the physically verified structure of the repository. It relies heavily on the `RepoIndexer` and `ArchitectureGraph`.

## 2. Repository Graph Embeddings
- **Strategy:** Instead of embedding raw code chunks, we embed the structural representation of the `ArchitecturalNode`. This includes its properties (`Mutability`, `Criticality`), its docstrings, and its explicit edges (dependencies, usages).
- **Benefit:** When reasoning about a change, the model is forced to acknowledge the physical boundaries of the code, preventing it from hallucinating connections that do not exist in the AST.

## 3. Impact on Planning
- **Topology-Aware Planning:** The `Planner` does not just output a sequence of text edits. It outputs a series of graph mutations.
- **Execution:** To edit file A, the planner must first request the `ArchitectureGraph` for file A. It verifies `BoundaryViolation` rules. If file A is locked or critical, the planner must dynamically adapt to modify an interface or create an adapter, rather than forcing a brute-force edit.

## 4. Reinforcement Learning from Repairs (RLFR)
- **Concept:** While true RL is complex, we can simulate RLFR deterministically.
- **Implementation:** Every time a topology-aware repair succeeds (validates and passes tests), the `SemanticTopologyEngine` records a positive weight on that specific graph transformation pattern. Future planning sessions prioritize this transformation pattern.
