# Runtime Graph: Causal Lineage & Process DAGs (L4)

## 1. Objective
Reconstruct parent-child process lineage to produce an incident DAG for timeline replay and causal analysis.

## 2. Graph Structure
- **Nodes**: Process events (fork, execve, exit).
- **Edges**: Causal parent-child relationships.
- **Identifiers**: Monotonic IDs combined with boot-time entropy to prevent PID reuse collisions.

## 3. Telemetry Ingestion
- **Syscalls**: `fork`, `clone`, `execve`, `setns`, `ptrace`.
- **Context**: Namespace and cgroup identifiers to track container escapes.

## 4. Integrity & Replay
The graph is stored in the 3-tier Trace Storage (HOT/WARM/COLD) and must be 100% reconstructible from the Truth Ledger evidence.

## 5. Security Applications
- **Drift Detection**: Detect when process behavior deviates from its historical graph signature.
- **Causal Analysis**: Trace the root cause of an alert back through the lineage DAG.
- **Isolation**: When a process is isolated, its entire sub-graph can be evaluated for containment.
