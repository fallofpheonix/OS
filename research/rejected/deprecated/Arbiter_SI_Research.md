# Research: Importance Scoring (SI) for Phoenix Arbiter

## Objective
Develop a mathematical framework to dynamically calculate the Importance Score ($S_I$) to differentiate between routine system noise and adversarial lateral movement.

## Proposed Formula Framework: The Multiplicative Model
$$ S_I = \min\left(1.0, \left( C_N \times G_S \right) \cdot e^{k \cdot F_A} \right) $$

- **Baseline Threat ($C_N \times G_S$):** The inherent risk is the severity of the event multiplied by the value of the node.
- **Compounding Factor ($e^{k \cdot F_A}$):** $F_A$ (Historical Frequency) acts as an exponential growth modifier, where $k$ is a tuning constant derived from the Phoenix Sentinel's thermodynamic entropy state.

## Variable Mapping Table

| Variable | Definition | Source in Phoenix OS |
| --- | --- | --- |
| **$C_N$** | Node Criticality (0.1 to 1.0) | Static YAML config injected via Cloud Control Plane. |
| **$G_S$** | Cgroup Severity (0.1 to 1.0) | eBPF cgroup/connect4 hook payload + PID context. |
| **$F_A$** | Historical Frequency (0.0 to 1.0) | Phoenix Monitor in-memory time-series window (e.g., last 60s). |

## Next Steps
1. Map out data sources for $C_N, F_A, G_S$ within the existing Telemetry/Trace framework.
2. Define the normalization functions for each variable (0.0 to 1.0).
3. Draft a simulation plan to test the $S_I$ sensitivity to synthetic attack patterns vs. normal operation.
