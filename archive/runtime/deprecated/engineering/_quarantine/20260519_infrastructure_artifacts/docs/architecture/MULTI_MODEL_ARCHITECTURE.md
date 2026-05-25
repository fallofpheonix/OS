# Multi-Model Architecture Design

## 1. Model Routing & Execution Constraints
Astraeus enforces a hard constraint: one active model at a time due to hardware bounds (M3 8GB) and deterministic execution requirements. The `TaskRouter` must dynamically allocate models sequentially without overlapping context windows.

### Task Allocation
- **Parsing (Deterministic extraction):** Local high-speed, low-parameter model (e.g., Llama-3-8B-instruct or specialized coder). Handles AST synthesis and syntax validation.
- **Planning (DAG generation):** High-reasoning model. Responsible for generating the `TaskGraph`, topological sorting, and dependency mapping.
- **Repair (Mutation loop):** Deep reasoning/coding model (e.g., DeepSeek-Coder). Must generate high-confidence semantic diffs.
- **Architecture Reasoning (Topology mapping):** Large context model. Evaluates `SideEffectType` and `Criticality`.
- **Memory Synthesis (Compression):** Fast, summarization-focused model. Compresses session data into `SessionMemory` and `ArchitectureMemory` before TTL expiration.

## 2. Planner / Verifier Loops
- **Loop:** The `Planner` model generates a mutation, which is instantly parsed and passed to the `Validator` model.
- **Verification:** The Verifier acts entirely independently. It uses Pydantic strict mode validation to verify the schema and runs static analysis tools. No code is executed until the Verifier yields a `RepairStatus.SUCCESS`.

## 3. Repair Validators & Uncertainty Estimators
- **Uncertainty Estimators:** The Router evaluates log-probabilities (or pseudo-confidence proxies) of the Repair model. If mutation confidence is low, it triggers a "Help Request" to the user or scales back the mutation size to a smaller subtree.
- **Repair Validators:** Uses `evaluator.py` to deterministically score the outcome. Fallbacks are restricted to prevent infinite hallucination loops, capping at `max_retries=2`.

## 4. Grounding & Safety
- **No AGI Abstractions:** The routing is purely rule-based and state-machine driven via `engine.py`. There are no autonomous agent loops that can self-modify the routing logic.
- **Observability:** Every transition between models is logged in the append-only SQLite DB.
