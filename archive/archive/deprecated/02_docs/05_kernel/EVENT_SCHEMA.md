# PhoenixOS Kernel Event Schema

This document defines the raw telemetry format for events ingested from the eBPF sensors (L2).

## Raw Kernel Event (JSON)

Every event emitted by the kernel probes MUST adhere to the following structure to ensure deterministic normalization.

```json
{
  "event_id": "string (UUID or sequence)",
  "pid": "integer (Process ID)",
  "ppid": "integer (Parent Process ID)",
  "process": "string (Binary name/path)",
  "syscall": "string (Name of the triggered syscall)",
  "timestamp": "integer (Nanoseconds since boot)",
  "cpu": "integer (Core ID where event originated)"
}
```

## Field Constraints
1. **Timestamp:** Must use `bpf_ktime_get_ns()` for monotonic timing across cores.
2. **Process Name:** Limited to 16 characters (TASK_COMM_LEN) unless extended.
3. **CPU ID:** Critical for identifying potential race conditions or core-affinity anomalies.
