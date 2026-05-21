# Stage 30: Security Games in Pheonix

## 1. Core Objective
To formulate and analyze strategic, sequential security interactions using Stackelberg Security Games (SSG). By designating Pheonix as the leader committing to randomized monitoring and containment policies, and attackers as rational followers who observe and adapt, Pheonix optimizes defensive allocations against intelligent adversaries.

---

## 2. Stackelberg Security Game Formulation

### 2.1. Basic Structure
A Stackelberg Security Game on host operations is defined by:
*   **Leader (Pheonix):** Commits to a randomized strategy profile (mixed strategy) $\mathbf{x} \in \Delta(\mathcal{C})$ representing the coverage probabilities of various telemetry and enforcement resources.
*   **Follower (Attacker):** Observes the leader's commit probability distribution (e.g. by measuring latency or tracking honeypot states) and selects a utility-maximizing action $a \in \mathcal{A}$.

### 2.2. Mathematical Model
Let:
*   $\mathcal{T}$: Set of security targets (critical processes, system files, connection sockets).
*   $C$: Set of defender coverage profiles (combinations of eBPF monitoring rate, CPU cgroups throttling, sandbox constraints).
*   $x_t \in [0, 1]$: Probability that target $t$ is covered by security controls.
*   $U_D^c(t)$ / $U_D^u(t)$: Defender's utility if target $t$ is attacked when covered ($c$) vs. uncovered ($u$).
*   $U_A^c(t)$ / $U_A^u(t)$: Attacker's utility if target $t$ is attacked when covered ($c$) vs. uncovered ($u$).

The defender's expected utility when target $t$ is attacked is:
$$U_D(t, x_t) = x_t U_D^c(t) + (1 - x_t) U_D^u(t)$$
The attacker's expected utility when attacking target $t$ is:
$$U_A(t, x_t) = x_t U_A^c(t) + (1 - x_t) U_A^u(t)$$

#### Optimization Formulation:
The defender computes the Strong Stackelberg Equilibrium (SSE):
$$\max_{\mathbf{x}, t^*} \quad U_D(t^*, x_{t^*})$$
$$\text{subject to} \quad t^* \in \arg\max_{t \in \mathcal{T}} U_A(t, x_t)$$
$$\sum_{t \in \mathcal{T}} x_t \le k \quad (\text{resource budget restriction})$$
Where $k$ is the maximum allowed monitoring overhead (CPU limit).

---

## 3. Subsystem Mapping & Implementation
*   **Subsystem Directory:** `07_security/game/stackelberg/`
*   **Components:**
    *   `leader/`: mixed-strategy allocator mapping eBPF sampling densities.
    *   `follower/`: adversary simulator predicting attacker response curves.

---

## 4. Experiment Backlog

### Experiment R030: Ransomware Stackelberg Damping Simulation
*   **Objective:** Simulate sequential ransomware adaptation where the process adjusts encryption velocity based on observed file read/write throttling. Compute the optimal SSE monitoring schedule.
*   **Telemetry Source:** system write rates, solver outputs.
*   **Metrics:**
    *   Attacker speed reduction $\ge 70\%$.
    *   Defensive utility improvement vs static rules $\ge 40\%$.
*   **Integration Target:** `07_security/game/stackelberg`.
