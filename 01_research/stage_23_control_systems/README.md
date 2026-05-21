# Stage 23: Control Systems in Pheonix

## 1. Core Objective
To treat security containment and telemetry rate control as a formal, closed-loop feedback control system. By modeling system behaviors using state-space variables and adjusting parameters via Proportional-Integral-Derivative (PID) and adaptive controllers, Pheonix stabilizes telemetry pipelines during event surges and enforces soft, multi-tier cgroups throttling on anomalous processes to damp attack velocity before hard termination.

---

## 2. Mathematical Formulations & Subsystem Mapping

### 2.1. Closed-Loop Feedback Control Model
The security containment loop is modeled as a continuous feedback loop:

```text
       r(t)        e(t)     +------------+  u(t)  +-------------+
Setpoint --------> ( - )--->| Controller |------->| Containment |
(Allowed Limit)     ^       +------------+        | Enforcement |
                    |                             +-------------+
                    |                                    |
                    |                                    | State Change
                    |                                    v
                    |       +------------+        +-------------+
                    +-------| Estimator  |<-------| System Host |
                            +------------+        |    x(t)     |
                                                  +-------------+
```

Where:
*   $x(t) \in \mathbb{R}^n$: **System State** (CPU utilization, file write rate, entropy, network connection rate).
*   $u(t) \in \mathbb{R}^m$: **Control Input** (cgroups quota limits, socket rate limits, sampling rate).
*   $y(t) \in \mathbb{R}^p$: **Observed Output** (measured threat levels, queue occupancy).
*   $r(t)$: **Target Setpoint** (e.g., maximum allowed queue size, allowable anomalous threshold).
*   $e(t) = y(t) - r(t)$: **Tracking Error**.

---

### 2.2. Mathematical Formulations

#### PID Feedback Control for Telemetry Sampling Rate:
To stabilize telemetry event queues during high-throughput execution (e.g., compilation or port scanning):
$$u(t) = K_p e(t) + K_i \int_0^t e(\tau) \, d\tau + K_d \frac{de(t)}{dt}$$
Where the control output $u(t)$ adapts the Bernoulli event filter sampling probability $p(t) = \max(0.001, p_{nominal} - u(t))$.

#### cgroups Dynamic Quota Throttling:
Rather than immediately terminating a business-critical database or service daemon, Pheonix dampens its capability via continuous CPU quota adjustment:
$$Quota(t) = Quota_{base} \cdot \left(1 - \min\left(1, K_p e_{susp}(t) + K_i \int_{0}^t e_{susp}(\tau) \, d\tau\right)\right)$$
Where $e_{susp}(t) = \text{Threat Level}(t) - \Theta_{allow}$. As the estimated threat level increases, CPU cycles are progressively starved to buy time for deeper AI analysis or forensic snapshotting.

---

### 2.3. Containment Levels
Pheonix defines five formal containment enforcement levels (L1–L5):

*   **L1: Alert & Telemetry Escalation:** Passive notification. Increases telemetry sampling rate $p \to 1.0$ for the target process.
*   **L2: Rate Limiting & Throttling:** Enforces CPU cgroups throttling and network bandwidth shaping.
*   **L3: Process Freezing:** Suspends process execution using `SIGSTOP` or cgroup freezers (`cgroup.freeze`), freezing its state in memory.
*   **L4: Network Isolation:** Applies eBPF TC/XDP or iptables rules to drop all outgoing and incoming packets from the container/host except to the SIEM.
*   **L5: Forensic Snapshot & Termination:** Performs a memory dump/core snapshot of the process tree followed by immediate `SIGKILL` termination.

---

## 3. Subsystem Mapping
*   **Engine Directory:** `07_security/control/`
*   **Components:**
    *   `pid/`: Core implementation of PID loop calculations for queue size and throttle dynamics.
    *   `adaptive/`: Parameter self-tuning algorithms (e.g. gain scheduling) based on historical noise levels.
    *   `reinforcement/`: Experimental learning agent for fine-tuning containment policies under mixed system workloads.

---

## 4. Experiment Backlog

### Experiment R023: PID Containment & Damping Evaluation
*   **Objective:** Inject a simulated resource-exhaustion or high-rate attack payload, apply cgroup-based PID throttling, and measure overshoot, settling time, containment latency, and overall host stability.
*   **Telemetry Source:** eBPF `sched_process_exec`, cgroups statistics, CPU load metrics.
*   **Metrics:**
    *   Throttling overshoot $\le 10\%$.
    *   Stability convergence time $\le 3.0$ seconds.
    *   Response execution latency $\le 50$ ms.
*   **Integration Target:** `07_security/control/pid`.
