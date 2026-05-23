# Security Boundaries & Isolation Zones

PhoenixOS operates strict isolation zones to contain anomalies and protect platform integrity.

## 1. Fast-Path Isolation (L1 Guard)
- Syscalls that trigger high-confidence alerts (entropy > 7.9) immediately trigger containment in under 100ms via L1 Guard kernel hooks, bypassing Userspace, Strategic Policy, and AI layers.

## 2. Warden State Containment (L5)
- **CONTAINED State:** Suspicious PIDs are frozen using cgroups v2 freezer controller. Sockets are isolated via XDP eBPF filters.
- **Privilege Limits:** Warden runs with minimal caps (CAP_SYS_ADMIN for cgroups control and eBPF maps) and is locked to PID 1 in the final OS appliance.

## 3. Cryptographic Tamper-Proofing (Ledger)
- The evidence log is append-only.
- Every entry hash incorporates the previous block's hash.
- Gaps or clock tampering immediately cause `Ledger.Verify()` to fail, preventing node boot or peer gossip verification.
