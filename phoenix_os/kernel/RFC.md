# RFC: Phoenix Kernel

## 1. Description
The Phoenix Kernel module handles in-kernel event capture (L2) and policy enforcement (L1). It acts as the primary data source for the Phoenix Bus.

## 2. Specification
- **Probes:** `tracepoints` for syscalls, `kprobes` for internal kernel functions.
- **Enforcement:** `LSM` hooks (Linux Security Modules) for mandatory access control.
- **Communication:** `BPF_MAP_TYPE_RINGBUF` for high-speed event export.

## 3. Security
All BPF programs are verified by the kernel verifier to ensure:
- No infinite loops.
- No invalid memory access.
- Bounded execution time.
