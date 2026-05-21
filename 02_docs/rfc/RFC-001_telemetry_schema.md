# RFC-001: Telemetry Schema

## Status
Approved

## 1. Purpose
This RFC standardizes the schema for all system-level telemetry captured within the Phoenix (PhoenixOS) ecosystem. The schema must support ingestion from both raw Linux-native eBPF collectors and simulated host-telemetry sources on non-Linux development machines (macOS).

## 2. Dependencies
*   **Compile-time:** Protobuf v3 / Go JSON parser
*   **Runtime:** Linux Kernel >= 5.15 (for production eBPF collectors) / Go 1.20+ (for mock collectors)

## 3. Data Schemas
Every event contains a **Common Header** and an **Event Payload**.

### 3.1 Common Header Schema
```json
{
  "timestamp": "string (RFC3339Nano)",
  "event_id": "string (UUIDv4)",
  "category": "string (process | syscall | filesystem | network | container | memory)",
  "event_type": "string",
  "host_id": "string",
  "pid": "uint32",
  "ppid": "uint32",
  "uid": "uint32",
  "gid": "uint32",
  "comm": "string",
  "exe_path": "string",
  "container_id": "string (optional)"
}
```

### 3.2 Category-Specific Payloads

#### 1. Process (`process`)
Captured on `sched_process_fork`, `sched_process_exec`, and `sched_process_exit`.
```json
{
  "args": ["string"],
  "env_vars": ["string"],
  "exit_code": "int32 (only for exit event)"
}
```

#### 2. System Call (`syscall`)
Captured on `sys_enter_*` and `sys_exit_*`.
```json
{
  "syscall_nr": "uint64",
  "args": ["uint64"],
  "retval": "int64"
}
```

#### 3. Filesystem (`filesystem`)
Captured on `vfs_open`, `vfs_write`, `vfs_read`, `vfs_unlink`.
```json
{
  "file_path": "string",
  "flags": "uint32",
  "mode": "uint32",
  "bytes_requested": "uint64",
  "bytes_transferred": "int64"
}
```

#### 4. Network (`network`)
Captured on `sys_enter_connect`, `sys_enter_accept`, `tcp_set_state`.
```json
{
  "saddr": "string",
  "daddr": "string",
  "sport": "uint16",
  "dport": "uint16",
  "protocol": "string (TCP | UDP | RAW)",
  "state_change": "string"
}
```

#### 5. Container (`container`)
Derived from cgroups namespaces (`mnt`, `pid`, `net`).
```json
{
  "namespace_pid": "uint32",
  "namespace_net": "uint32",
  "namespace_mnt": "uint32",
  "cgroup_path": "string"
}
```

#### 6. Memory (`memory`)
Captured on `sys_enter_mprotect`, `sys_enter_mmap`, `mm_page_alloc`.
```json
{
  "address": "uint64",
  "length": "uint64",
  "protection": "uint32 (PROT_READ | PROT_WRITE | PROT_EXEC)",
  "flags": "uint32"
}
```

---

## 4. Interfaces
The Telemetry schema will be implemented as:
1.  **JSON Schema:** [02_docs/schemas/telemetry_events.json](file:///Users/fallofPhoenix/os/02_docs/schemas/telemetry_events.json) for validating JSON over socket IPC.
2.  **Go Struct definitions** for the Broker and Phoenix Bus ingestion pipelines.

---

## 5. Threat Assumptions
*   **Log Tampering:** Attackers with root privileges might modify raw log files. *Mitigation:* Stream events in real-time to memory-buffered Unix domain sockets connected to the isolated Phoenix Bus broker.
*   **Sensor Blinding:** A malicious process may attempt to unload eBPF probes. *Mitigation:* Telemetry agent runs with immutable flags and monitors its own eBPF map status.
*   **Log Flooding (DoS):** Malware might generate millions of events to crash the correlator. *Mitigation:* Implement in-kernel eBPF rate-limiting maps and user-space ring buffers with oldest-dropped eviction policy.

---

## 6. Performance Expectations & Budget
*   **Throughput Budget:** Must handle up to **100,000 events/sec** per node.
*   **Latency Budget:** Time from kernel event hook to user-space agent serialization must be **< 50 microseconds**.
*   **Serialization Overhead:** Max **5 microseconds** per event.
*   **Payload Size Limit:** Serialized JSON payload must not exceed **512 bytes** for typical process/network events.

---

## 7. Failure Modes
1.  **Ring Buffer Loss:** The eBPF kernel-to-user ring buffer fills up, causing events to be dropped.
    *   *Indicator:* Kernel drops count incremented.
    *   *Action:* User-space agent dynamically adjusts sampling rates.
2.  **Parser Failures:** Input formatting changes or corrupted packets arrive at the broker.
    *   *Action:* Broker routes malformed events to a dead-letter queue (DLQ) and increments an alert counter.

---

## 8. Test Strategy
*   **Unit Tests:** JSON schema validator unit tests in Go.
*   **Fuzz Testing:** Fuzz the JSON parser with corrupted and truncated packets.
*   **Load Testing:** Feed mock telemetry stream at 150,000 events/sec and verify memory usage remains stable (< 50MB RAM for the agent).
