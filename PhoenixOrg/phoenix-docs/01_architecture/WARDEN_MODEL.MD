# Phoenix Warden: Actuation & Control Model (L5)

## 1. Objective
The Warden provides the deterministic actuation layer for PhoenixOS. It enforces system state transitions and applies resource constraints based on strategic policy.

## 2. Finite State Machine (FSM)
The Warden operates a 5-state discrete FSM:
- **SAFE**: Normal operation. No constraints.
- **WATCH**: Elevated monitoring. No active constraints.
- **SUSPICIOUS**: Initial containment. Soft resource limits applied.
- **CRITICAL**: Hard containment. Process isolation and significant throttling.
- **COMPROMISED**: Full isolation. Node-level lockdown.

## 3. Stabilization Constraints
To prevent rapid oscillation and "containment storms":
- **Actuation Cooldown**: 10 logical ticks post-actuation where no further transitions are allowed (except Class 3+ overrides).
- **State Dwell Limits**: Minimum of 30 logical ticks in elevated states before de-escalation.
- **Recovery Budgeting**: Maximum of 3 automatic de-escalations per session.

## 4. Actuation Classes
- **CLASS 0 (Observe)**
- **CLASS 1 (Log/Alert)**
- **CLASS 2 (Throttle)**: Resource limits (cgroups).
- **CLASS 3 (Local Isolate)**: Container/Process isolation.
- **CLASS 4 (Cluster Isolate)**: Node isolation.
- **CLASS 5 (Kernel Emergency)**: Panic or hardware-level stop.

## 5. Implementation Details
- **Controller**: PID (Proportional-Integral-Derivative) for throttling.
- **Actuators**: eBPF egress filters, cgroups CPU limits, `SIGSTOP` signals.
- **Veto Power**: The Warden can veto AI-suggested actions if they violate safety budgets or dwell limits.
