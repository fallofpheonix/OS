---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Security — Failure Modes

> Last verified: 2026-06-04

## Known Failure Vectors

| ID | Failure Mode | Mitigation |
|----|--------------|------------|
| S-01 | Unshare Privilege Denied | Sandbox requires process capability `CAP_SYS_ADMIN`. Fall back to standard container boundaries if missing. |
| S-02 | eBPF Probe Unload Failure | Kernel holds references. Implement automatic cleanups using driver close events. |
| S-03 | Quench Loop Failure | Killswitch fails to execute. The system executes hard halt via sysrq-trigger. |
