# Threat Model: PhoenixOS

## 1. Assets & Attack Surface

| Asset | Importance | Primary Entry Point | Impact of Compromise |
| :--- | :--- | :--- | :--- |
| **Phoenix Ledger** | Critical | Kernel tampering | Loss of forensic integrity. |
| **eBPF Probes** | High | Syscall injection | Blindness to system activity. |
| **Warden Controller** | High | Policy manipulation | Denial of service / Illegal actuation. |
| **Telemetry Bus** | Medium | Memory corruption | Signal noise / Spoofing. |

## 2. Attacker Profiles
- **Automated Malware:** Rapid propagation, targeted at L1/L2 bypass.
- **Advanced Persistent Threat (APT):** Stealthy, long-term persistence in WARM/COLD storage.
- **Insider Threat:** Authorized access to L5.5 policy configuration.

## 3. Defense Table

| Threat | Mitigation | Layer |
| :--- | :--- | :--- |
| **Probe Tampering** | Verifiable Ledger entries for probe load/unload. | L1/L2 |
| **State Injection** | HMAC-signed state transitions in Warden. | L5 |
| **Signal Jamming** | Entropy-based anomaly detection on telemetry bus. | L3 |
| **Policy Override** | Multi-agent consensus required for CRITICAL transitions. | L7 |
