Risk Register Update — Phoenix
================================

Seeded risks detected during repo audit:

- Premature kernel work
  - Likelihood: High
  - Impact: High
  - Mitigation: Freeze kernel prototypes until experiments validated

- Missing telemetry schema enforcement
  - Likelihood: High
  - Impact: High
  - Mitigation: Prioritize RFC-001 and normalizer implementation

- Unconnected theory islands
  - Likelihood: Medium
  - Impact: Medium
  - Mitigation: Map each theory to telemetry and experiments (THEORY_TO_OS_MAP.md)
# Risk Register Update (Architecture Constraints)

| ID | Risk | Severity | Mitigation |
| :--- | :--- | :--- | :--- |
| **ARCH-01** | Mathematical computations (Entropy/Games) take >1ms, crashing the event pipeline. | CRITICAL | Move compute off the main eBPF ring-buffer reader thread. Implement C/Rust mathematical backend instead of Python. |
| **ARCH-02** | PID Controller overshoot starves benign system processes. | HIGH | Implement strict bounds (Max Throttle) and fail-open policies on PID actuators. Validate via R031. |
| **ARCH-03** | AI layers implemented before deterministic physics layers are proven. | CRITICAL | Enforce dependency lock: AI subsystems (`06_ai`) cannot be integrated into runtime until L4 (Graphs) and L6 (Physics) pass validation gates. |