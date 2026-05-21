# Game Theory Runtime Matrix

### 1. Strong Stackelberg Equilibrium (SSE)
*   **Formula:** $\max U_D(t^*, x_{t^*})$ subject to attacker best response.
*   **Meaning:** Pheonix (Leader) commits to a monitoring distribution over targets. Attacker (Follower) selects a target.
*   **Telemetry Source:** eBPF sampling allocations, critical file paths.
*   **Target Module:** `07_security/game`
*   **Acceptance Criteria:** Solver execution time <= 1 ms.

### 2. Bayesian Intent Classifier
*   **Meaning:** Continually updates belief $P(\text{Ransomware} \mid \text{Writes})$ using Bayesian updates.
*   **Telemetry Source:** Write rates + entropy scores.
*   **Target Module:** `07_security/game` / `06_ai/classifiers`
*   **Acceptance Criteria:** > 97% accuracy within 5 telemetry samples.

### 3. Multi-Agent Reinforcement Learning (MARL)
*   **Meaning:** Decentralized node coordination for swarm containment.
*   **Telemetry Source:** Inter-node telemetry.
*   **Target Module:** `06_ai/agents`
*   **Acceptance Criteria:** Swarm consensus < 100ms.