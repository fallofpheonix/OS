# Project Evolution

## Current State
- **Phase**: F0 Stabilization
- **Substrate**: Verified Deterministic (Simulation)
- **Structure**: Sanitized Root, Categorized Externals
- **Next Primary Blockers**: Kernel Determinism (Hardware), External Adaptation

## Future State
- **F1 Maturity**: Kernel-level enforcement with <100ms Fast-Path.
- **F2 Formal**: TLA+ verified Warden FSM.

## Risk Register
- **R1**: Logic drift during external repo integration.
- **R2**: Kernel/Userspace logical clock synchronization on physical hardware.

## Dependency Map
- **D1**: `10_kernel` stress tests → F1 Runtime Unlock
- **D2**: `external/adapters` → astraeus/Noesis integration

## Unlock Conditions
- [ ] 100% Determinism match on multi-core kernel stress.
- [ ] Zero illegal dependencies in `external/dependent` modules.
