# Strict Implementation Gates

### Gate 1: Telemetry Saturation
*   **Constraint:** No graph processing or game theory implementation can begin until `09_telemetry/ebpf` captures >100k events/sec with <1% CPU overhead.
*   **Enforcement:** R001 Benchmark must pass and be merged.

### Gate 2: Graph Extraction Latency
*   **Constraint:** No physical modeling (SDI) or PID control can begin until process lineage DAGs (L4) can be extracted and queried in <1ms.
*   **Enforcement:** R003 and R022 Benchmarks must pass.

### Gate 3: Game Theory Solver Performance
*   **Constraint:** No PID loop or kernel scheduling can use Stackelberg outputs until the solver consistently returns policies in <1ms.
*   **Enforcement:** R027 Matrix Allocator must pass.

### Gate 4: Kernel Space Lock
*   **Constraint:** **ABSOLUTELY NO KERNEL MODIFICATIONS (`10_kernel/*`)** until userspace PID control (Phase D) is completely validated under adversarial conditions in `14_experiments`.
*   **Enforcement:** R031 must pass.

### Gate 5: AI Evidence Layer
*   **Constraint:** AI/ML (`06_ai`) modules are barred from taking autonomous actions. They remain strictly advisory (RAG, assistant) or experimental until Game Theory (L5.5) deterministic bounding is proven.
*   **Enforcement:** Architectural review required before AI can trigger containment.