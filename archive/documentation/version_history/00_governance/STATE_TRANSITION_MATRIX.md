# PhoenixOS: State Transition Matrix

This document defines the authorized state transition matrix for the containment FSM lifecycle in Warden and local subsystem controllers.

---

## 1. Transition Mapping Matrix

| Current State | Target State | Status | Validation Rule / Rationale |
| :--- | :--- | :--- | :--- |
| `SAFE` | `WATCH` | **✓ VALID** | Standard elevation during anomaly threshold trigger. |
| `SAFE` | `ALERT` | **✓ VALID** | Fast-path escalation on immediate critical severity indicator. |
| `SAFE` | `CONTAIN` | **✗ ILLEGAL** | Bypasses watch and alert analysis; direct containment blocked. |
| `WATCH` | `ALERT` | **✓ VALID** | Standard escalation from continuous monitoring. |
| `WATCH` | `CONTAIN` | **✓ VALID** | Allowed when cost of attack exceeds limits. |
| `WATCH` | `RECOVER` | **✗ ILLEGAL** | Cannot recover without active containment isolation first. |
| `ALERT` | `CONTAIN` | **✓ VALID** | Normal response pathway to quarantine threats. |
| `CONTAIN` | `RECOVER` | **✓ VALID** | Commences rollback restoration once system is isolated. |
| `RECOVER` | `SAFE` | **✓ VALID** | Returns system to initial clean state post-rollback validation. |

---

## 2. Dynamic Enforcement Code Integration
The rules above are programmatically enforced inside [policy.go](file:///Users/fallofpheonix/os/phoenix_os/containment/policy.go#L30-L45) via `isValidContainmentTransition`. Any illegal state jump attempted by the controllers is rejected with a structured transition error:
```go
func isValidContainmentTransition(current, next IsolationState) bool {
	switch current {
	case StateObserve:
		return next == StateWatch || next == StateThrottle
	case StateWatch:
		return next == StateThrottle || next == StateIsolate
	case StateThrottle:
		return next == StateIsolate
	case StateIsolate:
		return next == StateRecover
	case StateRecover:
		return next == StateObserve
	default:
		return false
	}
}
```
*Note: In code implementations, `StateObserve` represents `SAFE`, `StateThrottle` represents `ALERT`, and `StateIsolate` represents `CONTAIN`.*

---

## 3. Stabilization & Operational Rules

### Dwell Rules (Hysteresis)
- **Rule:** A minimum dwell duration of **30 logical clock ticks** is enforced in the `ALERT` and `CONTAIN` states before any de-escalation is allowed.
- **Purpose:** Prevents rapid state oscillations (flapping) during erratic telemetry inputs.

### Cooldown Periods
- **Rule:** A stabilization cooldown of **10 logical clock ticks** is enforced post-actuation (transitioning to `CONTAIN`).
- **Purpose:** Temporarily locks state transitions to let actuation effects stabilize. Critical class 3+ overrides bypass this cooldown.

### Rollback Policies
- **Rule:** Commencing `CONTAIN -> RECOVER` triggers the rollback engine to restore process, network, and file states to the last byte-for-byte verified snapshot.
- **Purpose:** Ensures deterministic state recovery with zero-drift telemetry verification.

### Emergency Paths (Phoenix Guard Fast-Path)
- **Rule:** Telemetry values with entropy > 7.9 trigger Phoenix Guard fast-path, executing a direct transition from `SAFE -> ALERT` under 100ms.
- **Purpose:** Bypasses intermediate strategic policy layers for immediate, low-latency critical path enforcement.

