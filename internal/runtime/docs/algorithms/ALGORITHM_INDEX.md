---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Runtime — Algorithm Index

> Last verified: 2026-06-04

Detailed review of algorithmic implementation inside runtime.

## 1. Byzantine Fault Tolerant Consensus
- **Core Loop**: Round-robin leader election, proposal broadcast, double-round voting (prepare/commit).
- **Time Complexity**: $O(N^2)$ communication complexity where $N$ is the number of active validators.

## 2. Linear History Fork Detection
- **Function**: `ForkDetector.Detect()`
- **Logic**: Iterates over matching logical time sequences. If parent hashes differ for the same sequence index, a fork is signaled.
- **Time Complexity**: $O(1)$ lookup via indexed parent-hash indexes.

## 3. Invariant Engine Validation
- **Loop**: Evaluates registered invariant functions sequentially.
- **Time Complexity**: $O(I)$ where $I$ is the number of invariants registered.
