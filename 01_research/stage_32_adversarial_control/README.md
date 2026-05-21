# Stage 32: Adversarial Control in Phoenix

## 1. Core Objective
To mathematically integrate continuous feedback control systems (PID, state-space models) with game-theoretic utility payoffs. By synthesizing control systems with adversarial game strategies, Phoenix ensures that physical system dynamics remain stable and optimal even under hostile manipulation and active evasion attempts.

---

## 2. Mathematical Synthesis

### 2.1. Game-Controlled Feedback Loop
In standard control systems, the controller parameters (e.g. PID gains $K_p, K_i, K_d$) are static. In Phoenix, these parameters are dynamically adjusted by the Phoenix Arbiter based on the active player strategy.

```text
               +----------------------------------+
               | Phoenix Arbiter / Payoff Calculation |
               +----------------------------------+
                                | Adjust Gains (Kp, Ki, Kd)
                                v
  r(t)  e(t)   +-----------------+  u(t)   +-------------------+
----->( - )--->| PID Controller  |-------->| Containment       |
       ^       +-----------------+         | Throttling Actuator|
       |                                   +-------------------+
       |                                             |
       +----------------- Host State x(t) <----------+
```

### 2.2. State-Space Game Model
The host dynamics are modeled in discrete-time state-space form under dual inputs:
$$\mathbf{x}_{k+1} = \mathbf{A}\mathbf{x}_k + \mathbf{B}_D \mathbf{u}_k + \mathbf{B}_A \mathbf{v}_k + \mathbf{w}_k$$

Where:
*   $\mathbf{x}_k \in \mathbb{R}^n$: State vector of the host (resource usages, queue sizes).
*   $\mathbf{u}_k \in \mathbb{R}^m$: Control input chosen by Phoenix (cgroups CPU quota, bandwidth).
*   $\mathbf{v}_k \in \mathbb{R}^p$: Disturbance input chosen by the Attacker (encryption rate, process spawn speed).
*   $\mathbf{w}_k$: Environmental process noise.

#### Linear-Quadratic Dynamic Game (LQ Game):
Phoenix minimizes system cost while the attacker maximizes damage:
$$J_D = \sum_{k=0}^{\infty} \left( \mathbf{x}_k^T \mathbf{Q}_D \mathbf{x}_k + \mathbf{u}_k^T \mathbf{R}_D \mathbf{u}_k \right)$$
$$J_A = \sum_{k=0}^{\infty} \left( \mathbf{x}_k^T \mathbf{Q}_A \mathbf{x}_k + \mathbf{v}_k^T \mathbf{R}_A \mathbf{v}_k \right)$$

Phoenix solves the coupled Riccati equations to find the optimal feedback policy $\mathbf{u}_k = -\mathbf{K}_D \mathbf{x}_k$ that is robust against the worst-case strategic attack input $\mathbf{v}_k$.

---

## 3. Subsystem Mapping
*   **Subsystem Directory:** `07_security/control/`
    *   `pid/`: PID engine with dynamic gain adjustments.
    *   `adaptive/`: Parameter gain schedulers driven by game-theory decisions.

---

## 4. Experiment Backlog

### Experiment R032: Adversarial PID Damping
*   **Objective:** Inject a simulated evasive workload that oscillates to trigger control system overshoot. Run the LQ game stabilizer to adjust PID gains and damp the oscillation.
*   **Metrics:**
    *   Overshoot reduction $\ge 50\%$.
    *   Settling time $\le 2.0$ seconds.
*   **Integration Target:** `07_security/control/adaptive`.
