# Future ML Integration Strategy

## 1. Grounding ML in Determinism
Astraeus explicitly rejects AGI claims and fake abstractions. All future ML integrations must be physically verifiable, observable, and strictly bounded by the DAG orchestrator.

## 2. Symbolic + Neural Hybrids
- **Direction:** Pure neural reasoning (LLMs) is prone to hallucination. Pure symbolic reasoning (AST/Static Analysis) is rigid. 
- **Implementation:** Create a hybrid pipeline. The neural model proposes a repair or plan, but it is immediately mapped to a symbolic representation (a proposed AST diff). The symbolic engine (Python `ast`, `MyPy`) verifies the structural integrity of the neural output *before* it is ever written to disk or executed in the sandbox.

## 3. Self-Verification Systems
- **Automated Repair Evaluation:** Instead of asking an LLM "does this look right?", the system compiles the code, runs the test suite, and extracts the stack trace. The LLM is only utilized to interpret the deterministic stack trace, not to guess the execution outcome.
- **State Snapshots (Cognition State):** Before invoking any large model, the `SessionManager` creates a lightweight snapshot of the current `TaskGraph` state. If the model fails or times out, the orchestrator rolls back to the snapshot deterministically, bypassing the need for neural recovery.

## 4. Code Embeddings & Adaptive Retrieval
- **Future Integration:** Moving beyond standard sentence-transformers to specialized AST-aware code embedding models. These models will natively understand `SideEffectType` and variable scoping, ensuring that retrieval pulls in not just lexically similar text, but structurally equivalent logic patterns.
- **Control Plane Evolution:** The orchestrator will remain deterministic. ML is strictly relegated to the worker nodes (parsing, planning, generating code), while the control plane (DAG, queuing, transaction safety) remains 100% hard-coded Python.
