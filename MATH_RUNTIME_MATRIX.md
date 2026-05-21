# Mathematics Runtime Matrix

### 1. Shannon Entropy & KL Divergence
*   **Formula:** $H(X) = -\sum P(x_i) \log_2 P(x_i)$ and $D_{KL}(P \parallel Q)$
*   **Meaning:** Measures randomness of file writes to detect encryption.
*   **Telemetry Source:** eBPF `vfs_write`, `sys_write`
*   **Target Module:** `09_telemetry/entropy_engine`
*   **Acceptance Criteria:** 4KB block calculated in < 5 us.

### 2. Kalman & Wavelet Filters
*   **Formula:** Standard discrete-time Kalman filtering and Discrete Wavelet Transform.
*   **Meaning:** Smooths CPU jitter; isolates periodic tasks from anomalies.
*   **Telemetry Source:** Disk write timers, CPU queue stats.
*   **Target Module:** `09_telemetry/math_filters`
*   **Acceptance Criteria:** Smoothing latency < 100 us.

### 3. VCG Auction Economics
*   **Formula:** $p_i = \max \sum_{j \neq i} \hat{v}_j(y) - \sum_{j \neq i} \hat{v}_j(x^*)$
*   **Meaning:** Allocates resources (CPU/Network) transparently based on bid curves to prevent starvation attacks by strategic, greedy processes.
*   **Telemetry Source:** Process priority/cgroup allocations.
*   **Target Module:** `10_kernel/scheduler`
*   **Acceptance Criteria:** Allocation efficiency >= 92%.