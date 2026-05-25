# Failure Domains

| Domain | Impact | Recovery | Rollback | Containment |
| :--- | :--- | :--- | :--- | :--- |
| **Memory** | Partial loss of context | Reload from snapshot | Revert to last commit | Halt ingestion |
| **Validator** | System paralysis | Restart validation service | Bypass to SAFE | Halt actuation |
| **Agent** | Task failure | Reroute task | Re-spawn agent | Sandbox isolate |
| **Router** | Misrouting / Load spill | Failover to default model | Revert policy | Throttle requests |
| **Runtime** | Execution stall | Re-sync kernel events | Replay trace | Cgroup freeze |
| **Training** | Model drift | Revert to checkpoint | Rollback dataset | Pipeline halt |
| **Telemetry** | Visibility loss | Re-probe eBPF | Re-calibrate clock | Fail-safe mode |
| **Security** | Integrity violation | System Lockdown | Re-verify hash chain | Hard isolation |
