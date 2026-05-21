# Stage 31: Multi-Agent Security in SentinelOS

## 1. Core Objective
To formulate cooperative, multi-agent defense architectures for host clusters and containers. By modeling individual security agents (telemetry, graph, physics, containment) as players in a cooperative game, SentinelOS optimizes joint resource consumption (Shapley value) and achieves decentralized consensus regarding active threats.

---

## 2. Mathematical Formulations

### 2.1. Cooperative Games & Shapley Value Allocation
When multiple system daemons share resources (e.g. CPU cache, disk bandwidth) and security agents enforce constraints, we assign overhead credits and threat costs fairly using the Shapley value.

Let $v(S)$ be the characteristic function mapping a coalition of agents $S \subseteq \mathcal{N}$ to their joint utility (e.g. threat detection rate or overhead reduction).
The marginal contribution of agent $i$ is calculated as:
$$\phi_i(v) = \sum_{S \subseteq \mathcal{N} \setminus \{i\}} \frac{|S|! (|\mathcal{N}| - |S| - 1)!}{|\mathcal{N}|!} \left[ v(S \cup \{i\}) - v(S) \right]$$

Where:
*   $\phi_i(v)$: Shapley value of agent $i$.
*   $\mathcal{N}$: The set of all system agents.
*   *Application:* Assigns fair performance costs to container groups under containment throttling, preventing arbitrary starvation of benign processes.

---

### 2.2. Multi-Agent Reinforcement Learning (MARL) Swarms
Hosts cooperate in a decentralized network to isolate lateral threat spread:
*   **State Space $\mathcal{S}$:** Global network threat state.
*   **Action Space $\mathcal{A}_i$:** Local containment choices of host $i$.
*   **Reward Function $\mathcal{R}$:** Unified system uptime minus propagation loss.
*   Decentralized agents coordinate policy weights using consensus algorithms (e.g. Raft or Gossip protocols) to arrive at a global optimal threshold without central coordination.

---

## 3. Subsystem Mapping
*   **Subsystem Directory:** `06_ai/game/`
    *   `evolution/`: Swarm coordination models.
    *   `adaptive/`: Collaborative multi-agent learning controllers.

---

## 4. Experiment Backlog

### Experiment R031: Swarm Consensus Under Attack
*   **Objective:** Deploy a cluster of 5 nodes running local cooperative game models. Inject a network threat and verify consensus propagation speed.
*   **Metrics:**
    *   Consensus convergence latency $\le 100$ ms.
    *   Containment path accuracy $\ge 95\%$.
*   **Integration Target:** `06_ai/game/evolution`.
