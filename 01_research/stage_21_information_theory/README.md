# Stage 21: Information Theory in Pheonix

## 1. Core Objective
To leverage information theory as a formal, low-overhead detection and optimization layer. By measuring the uncertainty, information content, and statistical drift of system resources and event streams, Pheonix can detect malicious activity (such as ransomware encryption or stealthy persistence) that evades traditional threshold-based systems, while optimizing telemetry pipelines via entropy-based adaptive sampling.

---

## 2. Mathematical Formulations & Subsystem Mapping

### 2.1. Shannon Entropy ($H$) for High-Entropy Write Detection
Ransomware attacks perform massive encryption of user files, which manifests as high-entropy data writes to the filesystem. Pheonix calculates the Shannon entropy of write buffers to distinguish plain text/structured data from encrypted or compressed payloads.

#### Formula:
$$H(X) = -\sum_{i=0}^{255} P(x_i) \log_2 P(x_i)$$

Where:
*   $X$ represents a write buffer payload block (typically 4096 bytes).
*   $x_i$ is a byte value ($0 \le i \le 255$).
*   $P(x_i) = \frac{count(x_i)}{N}$ is the empirical probability of byte $x_i$ appearing in a block of size $N$.
*   $H(X)$ is measured in bits. A pure random/encrypted block yields $H(X) \approx 8.0$, whereas plain text English typically yields $3.5 \le H(X) \le 5.0$.

#### Subsystem Mapping:
*   **Collector Hook:** eBPF `kprobe/vfs_write` (or LSM hook `security_file_write`).
*   **Adaptive Sampling Optimization:** To avoid the CPU overhead of running $O(N)$ byte counts on every write, the eBPF agent applies a Bernoulli sampling filter:
    $$S \sim \text{Bernoulli}(p), \quad p = 0.05$$
    Only 5% of 4KB blocks are evaluated. If a process exhibits $H(X) \ge 7.2$ bits for $M \ge 3$ consecutive sampled blocks, the collector escalates sampling to $p = 1.0$ for that PID and notifies the **Security Layer**.

---

### 2.2. Kullback-Leibler (KL) Divergence ($D_{KL}$) for System Drift Analysis
Intruders attempting lateral movement or establishing persistence alter the system call signature distribution of compromised services. Pheonix calculates the relative entropy (KL Divergence) between the active system call profile and the baseline profile.

#### Formula:
$$D_{KL}(P \parallel Q) = \sum_{k \in \mathcal{S}} P(s_k) \log_2 \frac{P(s_k)}{Q(s_k)}$$

Where:
*   $\mathcal{S}$ is the set of all monitored system calls (e.g., `sys_enter_write`, `sys_enter_openat`, etc.).
*   $Q(s_k)$ is the baseline probability of syscall $s_k$ calculated over a historical window (e.g., last 24 hours).
*   $P(s_k)$ is the observed probability of syscall $s_k$ in the current sliding window (e.g., last 5 minutes).
*   $D_{KL}(P \parallel Q)$ is the information loss (in bits) when approximating the true distribution $P$ with baseline $Q$.

#### Subsystem Mapping:
*   **Engine Component:** Pheonix `AI Correlation Engine` (`06_ai/`).
*   **Actionable Threshold:** If $D_{KL}(P \parallel Q) > 0.5 \text{ bits}$ for a system service daemon (e.g., `sshd` or `httpd`), Pheonix triggers an alert for anomalous behavior drift, catching slow-build stealth attacks that keep execution rates below traditional volume thresholds.

---

### 2.3. Mutual Information ($I$) for Process Interaction Correlation
Attackers often use local inter-process communication (IPC) or stealthy piping to bridge phishing attachments to credential harvesting tools. Pheonix measures Mutual Information between processes to uncover hidden channels.

#### Formula:
$$I(X; Y) = \sum_{y \in \mathcal{Y}} \sum_{x \in \mathcal{X}} P(x, y) \log_2 \left( \frac{P(x, y)}{P(x) P(y)} \right)$$

Where:
*   $X$ and $Y$ are random variables representing the system calls executed by Process $A$ and Process $B$ in a synchronized temporal window.
*   $P(x, y)$ is the joint probability distribution of syscall sequences occurring concurrently.
*   $P(x), P(y)$ are marginal probability distributions.
*   $I(X; Y) = 0$ implies independent operation. High mutual information $I(X; Y) > \theta_{ipc}$ without explicit socket or pipe edges indicates covert process synchronization.

#### Subsystem Mapping:
*   **Engine Component:** `Process Lineage DAG` normalizer and correlator (`09_telemetry/bus/`).

---

## 3. Telemetry Ingestion Optimization (Adaptive Compression)
Using entropy values, the Event Bus Normalizer optimizes telemetry transport bandwidth:
1.  **Low-Entropy Event Elimination:** Telemetry events with Shannon information content below a dynamic threshold $H_{event} < 0.1 \text{ bits}$ (e.g., repeated duplicate read events from idle processes) are aggregated at the collector level rather than transmitted.
2.  **Telemetry Rate Adaptation:** Under CPU/network congestion, the event normalizer drops low-priority events, prioritizing events with high surprise value (high self-information: $I(x) = -\log_2 P(x)$).

---

## 4. Verification Gate & Validation Metrics
*   **Validation Test:** Run the `telemetry_replay` experiment. Insert simulated ransomware writes (high-entropy, pseudorandom) and benign writes (structured ASCII).
*   **Target Metric:** 
    *   Ransomware detection delay $\le 10$ file writes.
    *   Bernoulli sampling overhead $\le 0.5\%$ CPU overhead at $50,000 \text{ events/sec}$.
