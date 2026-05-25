# Long-Context & Memory Optimization

## 1. Memory Compression & Semantic Summarization
Long-running sessions rapidly consume the LLM context window. 
- **Design:** Implement periodic state compression. As a session progresses, completed task subtrees in the `TaskGraph` are summarized by a low-parameter model into dense semantic blocks.
- **Integration:** The `SessionManager` removes raw execution logs of completed tasks from the active context window, replacing them with the semantic summary in `SessionMemory`.

## 2. Episodic Replay
- **Concept:** To recover from catastrophic failures or to analyze previous deep reasoning steps, the system must support episodic replay.
- **Implementation:** Using the append-only SQLite store (`SQLiteMemoryStore`), Astraeus can replay the exact state and context of any past task node deterministically, without relying on LLM recall.

## 3. Temporal Indexing & Adaptive Forgetting
- **Temporal Indexing:** Every entry in `ArchitectureMemory` is tagged with a temporal epoch corresponding to a git commit or transaction ID.
- **Adaptive Forgetting:** If an architectural pattern has been completely refactored or removed (determined via `GraphInvalidationEngine`), the associated memory vectors undergo "adaptive forgetting." They are either deleted from the vector index or heavily penalized during retrieval to prevent ghost-context retrieval.

## 4. Execution-Aware Memory
- **Design:** Memory is not just text; it includes the state of the execution environment. `FailureMemory` must map the failure to the exact hardware/sandbox state. This ensures the memory is grounded in deterministic execution reality, not just LLM string similarity.
