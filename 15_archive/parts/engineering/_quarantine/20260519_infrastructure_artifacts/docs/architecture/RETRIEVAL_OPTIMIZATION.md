# Retrieval Optimization Strategy

## 1. Graph-Aware & Topology-Aware Retrieval
To reduce context explosion and limit irrelevant retrieval, the `RetrievalLayer` must move beyond pure vector similarity (which hallucinates context) to Graph-RAG over the `ArchitectureGraph`.
- **Implementation:** Query embeddings will first match an `ArchitecturalNode`. Instead of retrieving raw string chunks, the system will walk the `SemanticTopologyEngine` edges (dependencies, usages, boundaries) to collect deterministic contexts.
- **Safety:** Prevents hallucinated contexts because only statically verified edges (via the AST) are traversed. `BoundaryViolation` entities will actively block retrieval across untrusted boundaries.

## 2. Temporal Retrieval
Codebase intent is heavily tied to recent changes. `SemanticMemoryStore` must adopt a time-decay factor.
- **Implementation:** Boost embeddings of recently modified nodes (leveraging `ArchitecturalTemperature`). Older, stable modules are only retrieved if explicitly referenced by exact symbol matches.
- **Execution:** When `GraphInvalidationEngine` flags a subtree, the retrieval layer pins that subtree's context into active memory for the duration of the mutation session.

## 3. Symbol-Aware Retrieval
Vector searches often fail on strict deterministic symbols (e.g., `ClassDefinition` exact matches).
- **Implementation:** Hybrid search architecture. Exact symbol matching via SQLite (`SQLiteMemoryStore`) takes precedence over semantic similarity. If a prompt mentions `TaskRouter`, the system deterministically retrieves the symbol's AST scope before querying the vector database.
- **Impact:** Guarantees repository grounding. Semantic search is demoted to a fallback mechanism for natural language concepts lacking explicit symbols.

## Summary of Reductions
- **Context Explosion:** Solved by enforcing max-depth walks on the topology graph.
- **Irrelevant Retrieval:** Mitigated by prioritizing explicit symbol matching.
- **Hallucinated Context:** Prevented by anchoring all semantic matches to a physically validated `ArchitecturalNode` in the AST.
