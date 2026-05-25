# PX-013 Live probe determinism

**Label**: kernel-determinism, F0-exit
**Status**: IN_PROGRESS

## Problem
Current kernel determinism tests rely on bus simulations. Empirical validation requires testing against real eBPF probes and live kernel event streams to verify stability under actual system load.

## Tasks
- [x] Implement kernel replay injection mechanism (Scaffolded)
- [x] Verify ring buffer overflow under live probe load (Scaffolded)
- [x] Perform CPU affinity replay tests (Scaffolded)
- [x] Execute cross-core event ordering validation (Scaffolded)
- [x] Perform clock skew stress tests (Scaffolded)
- [x] Generate `KERNEL_RUNTIME_REPORT.md`

## Verification
- Ring pressure must be observed and handled without non-deterministic drops.
- Causal ordering must hold across multi-core execution traces.
