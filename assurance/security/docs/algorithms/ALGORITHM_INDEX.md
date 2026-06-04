---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Security — Algorithm Index

> Last verified: 2026-06-04

Detailed review of security math and isolation algorithms.

## 1. Stackelberg Defender Solver
- **Purpose**: Computes game-theoretic optimal defender response.
- **Algorithm**: Resolves quadratic optimization formulas modeling the probability of defense failure against cost constraints.
- **Time Complexity**: $O(K^2)$ where $K$ is the number of possible security policies.

## 2. State Disorder Index (SDI)
- **Purpose**: Measure anomaly levels through dynamic system disorder.
- **Formula**:
  $$\text{SDI} = - \sum_{i} P_i \ln P_i$$
  where $P_i$ is the probability distribution of event types over a rolling sliding-time window.

## 3. Sandboxing Mount Isolation
- **Purpose**: Creates unshared namespaces.
- **Steps**: Calls `syscall.Unshare(CLONE_NEWNS | CLONE_NEWPID | CLONE_NEWNET)`. Mounts read-only filesystems.
