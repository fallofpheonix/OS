# PhoenixOS Security Boundaries & Trust Model

## Philosophy
This document explicitly defines the trust boundaries between PhoenixOS subsystems. It dictates which components are authoritative, which are advisory, and the assumptions regarding adversary capabilities.

## Trust Boundaries

### 1. Telemetry Source (Guard)
- **Status**: **Trusted** (Kernel eBPF) / **Semi-Trusted** (Userspace Adapter)
- **Assumption**: Events retrieved from kernel space (eBPF) are assumed to be a true representation of system execution state. 
- **Forgery Resistance**: If an adversary gains ring-0 execution, the telemetry stream cannot be trusted. Phoenix relies on the Linux kernel's internal boundaries for primary event fidelity.

### 2. The Internal Event Bus (phoenix-bus)
- **Status**: **Trusted Boundary**
- **Assumption**: Data on the internal bus is sequentially monotonic. Memory access to the bus is restricted to the Phoenix runtime.

### 3. Trace Storage (phoenix-trace)
- **Status**: **Trusted but Verifiable**
- **Assumption**: The SQLite WAL database is stored securely. However, the integrity of the data is verified continuously via the `hash` and `prev_hash` chain.
- **Replay Attack Resistance**: The `monotonic_ns` field paired with the SHA-256 hash-chain ensures that an adversary cannot reorder, drop, or inject historical events into the SQLite store without invalidating the chain.

### 4. Monitor Anomaly Engine (phoenix-monitor)
- **Status**: **Advisory Only**
- **Assumption**: The Monitor's Kalman Filter and EWMA scores are purely mathematical evaluations. It has **no authority** to mutate system state, block processes, or write to the primary Trace store. It merely emits `telemetry.scored` messages to the Warden.

### 5. Control Layer (phoenix-warden)
- **Status**: **Authoritative**
- **Assumption**: The Warden FSM is the single source of truth for mitigation states (`NORMAL`, `SUSPICIOUS`, `CONTAINED`). Only the Warden can issue actions back to the Guard.

## Future Mitigations
- Implementation of signed telemetry via userspace enclave.
- Tamper-evident network replication of the hash chain.
