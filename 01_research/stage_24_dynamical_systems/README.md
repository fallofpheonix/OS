# Stage 24: Dynamical Systems in SentinelOS

## 1. Core Objective
To model the operating system host as a continuous-time, evolving dynamical system. By defining multi-dimensional host state trajectories, predicting resource trends, and modeling attack propagation (e.g., ransomware spread) using epidemiological differential equations, SentinelOS detects malicious divergences from normal execution trajectories and guarantees stable queueing behaviors under extreme loads.

---

## 2. Mathematical Formulations & Subsystem Mapping

### 2.1. Host State Trajectory & Differential Dynamics
The system's operational state at any instant is represented by a multi-dimensional state vector:
$$\mathbf{x}(t) = [x_1(t), x_2(t), x_3(t), x_4(t), x_5(t), x_6(t), x_7(t)]^T$$

Where:
*   $x_1(t)$: CPU utilization (load average).
*   $x_2(t)$: Physical memory RSS allocations.
*   $x_3(t)$: I/O read/write throughput (bytes/sec).
*   $x_4(t)$: Active process count.
*   $x_5(t)$: File modification rate.
*   $x_6(t)$: Network throughput and connection rate.
*   $x_7(t)$: Write buffer Shannon entropy.

The evolution of the host state is governed by a set of differential equations:
$$\frac{d\mathbf{x}}{dt} = \mathbf{f}(\mathbf{x}(t), \mathbf{u}(t)) + \mathbf{w}(t)$$

Where $\mathbf{u}(t)$ is the vector of control inputs and $\mathbf{w}(t)$ represents environmental noise.
*   **Anomalous Divergence:** SentinelOS maps the *expected trajectory* $\hat{\mathbf{x}}(t)$ based on historically learned vector flows. If the Euclidean distance between the *observed trajectory* $\mathbf{x}(t)$ and the expected trajectory exceeds a threshold, the system flags a trajectory anomaly:
    $$\|\mathbf{x}(t) - \hat{\mathbf{x}}(t)\| > \Theta_{dynamics}$$

---

### 2.2. Queueing Dynamics (Little's Law)
To configure the telemetry transport buffers without losing events:
$$L = \lambda W$$
Where:
*   $L$: Average number of telemetry events stored in the buffer (queue capacity).
*   $\lambda$: Average arrival rate of events from eBPF hook points (events/second).
*   $W$: Average processing latency (wait time) of the normalizer/ingest engine (seconds/event).

If telemetry rate surges due to an attack, the queue capacity must be sized such that $L_{size} \ge \lambda_{peak} W$ to guarantee zero event loss.

---

### 2.3. SIR Epidemic Model for Ransomware Containment
Ransomware encryption propagation behaves like an infectious disease outbreak within the filesystem lattice. We model this propagation mathematically:
$$\frac{dS}{dt} = -\beta \frac{S I}{N}$$
$$\frac{dI}{dt} = \beta \frac{S I}{N} - \gamma I$$
$$\frac{dR}{dt} = \gamma I$$

Where:
*   $S(t)$: Susceptible files (unencrypted, normal read/write files).
*   $I(t)$: Infected files (actively being written or locked by the ransomware thread).
*   $R(t)$: Recovered/Contained files (rendered read-only, isolated, or restored from backup).
*   $N$: Total file count.
*   $\beta$: Infection transmission coefficient (encryption velocity in files/sec).
*   $\gamma$: Recovery/Containment coefficient (rate of isolation and process damping).

#### The Basic Reproduction Number ($R_0$):
$$R_0 = \frac{\beta}{\gamma}$$
*   **Threshold condition:** To completely halt ransomware propagation, the containment rate must satisfy $R_0 < 1 \implies \gamma > \beta$.
*   If a ransomware process encrypts at $\beta = 100 \text{ files/second}$, and there is $k = 1$ active thread, the response latency $T_{resp} = \frac{1}{\gamma}$ to freeze the process must be $T_{resp} < 10 \text{ ms}$.

---

## 3. Subsystem Mapping
*   **Engine Directory:** `06_ai/dynamics/`
*   **Components:**
    *   `trajectory/`: Real-time state vector assembler collecting CPU, memory, IO, and write entropy.
    *   `forecast/`: Prediction engine calculating the expected future trajectory of the host state vector.
    *   `deviation/`: Comparator assessing divergence metrics between active and forecasted system paths.

---

## 4. Experiment Backlog

### Experiment R024: Trajectory Divergence Under Ransomware
*   **Objective:** Replay simulated ransomware and observe the state vector trajectory. Measure the deviation from the normal workload forecast and define the point of divergence.
*   **Telemetry Source:** eBPF process, disk, and entropy telemetry.
*   **Metrics:**
    *   Time-to-detect divergence $\le 3.0$ seconds.
    *   False positive rate under compilation workloads $\le 1\%$.
*   **Integration Target:** `06_ai/dynamics/trajectory`.
