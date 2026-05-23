# Threat Model: PhoenixOS

This threat model defines our assets, attacker vectors, entry points, impacts, and defenses under a Zero-Trust architecture.

## Asset Catalog
1. **Evidence Ledger:** Cryptographic proof chain documenting process actions and FSM state transitions.
2. **Telemetry Stream:** Live syscall records driving entropy monitoring.
3. **Warden FSM State:** Current system containment level (NORMAL -> CONTAINED).
4. **Policy Configuration:** Gating thresholds and de-escalation budgets.

## Attacker Profile & Scenarios

### Threat 1: Denial-of-Service via Telemetry Noise (Mimicry Attack)
- **Attacker:** Compromised low-privilege userspace process generating massive normal-severity traffic.
- **Entry Point:** Syscall interface (L1/L2).
- **Impact:** Ring-buffer saturation leading to dropping of actual critical threat alerts.
- **Defense:** *Evidence Reserve Lane* blocks low-severity events at 85% queue pressure. *Priority Pre-emption Shield* drops oldest low-severity events to accommodate critical events at 100% capacity.

### Threat 2: Time Warp State Poisoning
- **Attacker:** Root process attempting to inject fake historical events to trigger rollback or reset de-escalation history.
- **Entry Point:** Replay injection endpoint.
- **Impact:** FSM escapes containment or resets de-escalation budget without approval.
- **Defense:** Bounded sequence window verification in [tcs.go](file:///Users/fallofpheonix/os/phoenix_os/tcs/tcs.go) rejects future/past anomalies, and Ledger validation hashes enforce state progression integrity sequentially.

### Threat 3: Concurrent State Manipulation (SOC Operator Bypass)
- **Attacker:** Malicious script invoking SOC `/game/action` API concurrently during high-load telemetry runs.
- **Entry Point:** Game Server HTTP API.
- **Impact:** Warden state corruption, double-decrementing budgets.
- **Defense:** Enforced Mutex locks on `Warden` and RWMutex on `Ledger` ensure atomic, serialized state transitions.
