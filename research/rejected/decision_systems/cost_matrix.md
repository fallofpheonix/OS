# Phoenix Arbiter: Cost Matrix (L5.5)

This document defines the deterministic cost model for the Phoenix Arbiter Strategic Policy layer. It maps attack severity (Attack Cost) against containment actions (Containment Cost) to enable objective, game-theoretic decision making.

## 1. Attack Cost (AC)
Attack cost represents the calculated damage or risk to the system if an anomaly is allowed to persist.

| Severity Level | Base Cost | Description |
| :--- | :--- | :--- |
| **CRITICAL** | 1000.0 | System compromise, data exfiltration, or kernel instability. |
| **HIGH** | 500.0 | Unauthorized access to sensitive services, potential lateral movement. |
| **MEDIUM** | 100.0 | Resource exhaustion, unauthorized configuration changes. |
| **LOW** | 10.0 | Policy violations, unusual but non-destructive behavior. |
| **INFO** | 1.0 | Baseline drift, diagnostic interest. |

**Formula Modifier**: `AC_final = BaseCost * (1 + SystemLoad)`
*Where SystemLoad is a normalized value [0.0 - 1.0].*

## 2. Containment Cost (CC)
Containment cost represents the operational impact, service disruption, and resource overhead of a mitigation action.

| Action Class | Base Cost | Description |
| :--- | :--- | :--- |
| **ClassObserve** | 0.0 | No impact. Passive monitoring. |
| **ClassLog** | 1.0 | Minimal impact. Log rotation and storage overhead. |
| **ClassThrottle** | 50.0 | Moderate impact. Performance degradation for the target. |
| **ClassLocalIsolate** | 200.0 | High impact. Single process/container isolation. |
| **ClassClusterIsolate** | 800.0 | Severe impact. Service-wide isolation. |
| **ClassKernelEmergency**| 2000.0 | Existential impact. Node-wide lockdown. |

## 3. Decision Logic
The Arbiter recommends an action if:
`ContainmentCost < AttackCost`

**Optimization Goal**: Minimize `TotalCost = (ContainmentCost + ResidualAttackCost)` where `ResidualAttackCost` is the remaining risk after action.

## 4. Example Scenarios
- **Scenario A**: High threat on `nginx` (AC=500).
    - `StopService` (CC=800) -> Rejected (Cost too high).
    - `ContainProcess` (CC=200) -> **Recommended** (200 < 500).
- **Scenario B**: Medium threat on `sshd` (AC=100).
    - `ClassLocalIsolate` (CC=200) -> Rejected.
    - `ClassThrottle` (CC=50) -> **Recommended** (50 < 100).
