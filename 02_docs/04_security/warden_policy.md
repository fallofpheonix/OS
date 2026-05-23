# PhoenixOS Warden Policy (L5)

This document defines the deterministic state transitions and safety constraints for the PhoenixOS containment engine.

## 1. System States

| State | Definition | Action Class |
| :--- | :--- | :--- |
| **SAFE** | Normal operating parameters. | Observe |
| **WATCH** | Minor anomalies detected (e.g., fork burst). | Enhanced Observation |
| **ALERT** | Significant disorder detected (e.g., mass write). | Forensic Snapshot |
| **CONTAIN** | Critical threat detected (e.g., reverse shell). | Full Isolation |
| **RECOVERY** | Operator-led state restoration. | Release / Restore |

## 2. Transition Triggers

- **PROCESS_FORK_BURST:** Triggers transition to `WATCH`.
- **NETWORK_BEACON:** Triggers transition to `ALERT`.
- **MASS_WRITE:** Triggers transition to `ALERT`.
- **REVERSE_SHELL:** Triggers transition to `CONTAIN` immediately.
- **HUMAN_OVERRIDE:** Triggers transition to `RECOVERY`.

## 3. Safety Constraints (Axioms)

1. **No AI Control Path:** AI (PhoenixMind) can advise on triggers but cannot execute the state transition. The Warden FSM is the sole deterministic actuator.
2. **No Auto-Kill:** PhoenixOS prioritizes **Containment** (freezing/isolation) over **Termination**. Deletion/Kill requires explicit human override or L5.5 strategic consensus.
3. **Determinism:** Given the same sequence of normalized events and logical ticks, every Warden node MUST arrive at the same state.
4. **Hysteresis (Dwell Limits):** De-escalation (moving to a lower-severity state) is blocked for a minimum of 30 logical ticks to prevent "flapping."
5. **Recovery Budget:** Automatic de-escalation is limited to 3 cycles per boot. Further recoveries require a manual `ResetBudget` action.
