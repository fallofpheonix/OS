# Phoenix Kernel (L1/L2)

The foundational telemetry and enforcement layer of PhoenixOS.

## Purpose
Provides low-overhead system visibility and security enforcement using eBPF and LSM hooks.

## Performance Budget
- **Trace Latency:** < 500 ns per event.
- **CPU Overhead:** < 1% total system usage.

## Validation Gates
- [ ] eBPF verification success.
- [ ] Atomic map updates.
- [ ] Zero packet/event loss in ring buffer.
