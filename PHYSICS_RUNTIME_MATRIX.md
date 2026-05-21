# Physics Runtime Matrix

### 1. Ising Spin Lattice (Security Disorder Index)
*   **Formula:** $H(\sigma) = -J \sum \sigma_i \sigma_j - h \sum \sigma_i$
*   **Meaning:** Models the security state of N containers as a 2D lattice.
*   **Telemetry Source:** Container network connectivity, cross-process IPC.
*   **Target Module:** `07_security/physics`
*   **Acceptance Criteria:** Prediction rate of phase-transitions (cascading failures) >= 95%.

### 2. Arrhenius Activation Energy (Containment)
*   **Formula:** $k = A e^{-\frac{E_a}{k_B T}}$
*   **Meaning:** Calculates the probability of a threat escaping its sandbox based on the "threat temperature" ($T$) and sandbox restriction strictness ($E_a$).
*   **Telemetry Source:** Seccomp profiles, capability drops.
*   **Target Module:** `10_kernel/sandbox`
*   **Acceptance Criteria:** Escape probability accurately correlates with empirical threat models.