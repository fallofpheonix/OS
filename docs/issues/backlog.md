# GitHub Issue Backlog

## P0: Critical Infrastructure & Kernel Integration

### Issue #1: [CI/Kernel] Automated BPF Compilation Pipeline
**Body:**
## Problem
The eBPF kernel objects (`egress_policy.c`, `xdp_ingress.c`) cannot be compiled in the current Darwin/macOS environment, blocking validation of the fast-path network logic.

## Current State
C-source exists in `phoenix_os/warden/ebpf/src/`. CI workflow created but requires a configured Linux runner environment.

## Required Work
- [ ] Finalize `.github/workflows/bpf-build.yml` with proper Linux runner configuration.
- [ ] Validate bytecode generation on Linux environment.
- [ ] Automate ELF artifact publication.

## Risk
High (Deployment/Verification Blocker)

---

### Issue #2: [Infrastructure] LinuxKit Scaffolding for Immutable OS
**Body:**
## Problem
PhoenixOS currently runs as a Go binary on macOS. To move to a production-ready, zero-trust cloud OS, it must boot as an immutable, container-optimized LinuxKit image.

## Current State
Conceptual design exists in `phoenix_os/infrastructure/linuxkit/phoenix_os.yml`.

## Required Work
- [ ] Configure `init` to trigger the Warden binary on boot.
- [ ] Define the immutable root filesystem configuration.
- [ ] Implement ISO generation script.

## Risk
Medium

---

## P1: Telemetry, Observability & Forensics

### Issue #3: [Telemetry/Forensics] Implement Ledger Query & Replay API
**Body:**
## Problem
While the `Ledger` is implemented in-memory, there is no interface to query the chain to reconstruct system causality or validate the Warden's decisions after a containment event.

## Required Work
- [ ] Implement `GetEntryByActionID(actionID string)`.
- [ ] Implement `ReconstructCausalChain(actionID string)` to walk the `PrevHash` chain.
- [ ] Create a CLI tool: `phoenix-ctl ledger show --action <id>`.

## Risk
Medium

---

### Issue #4: [Security/Control] RFC-011: Telemetry Confidence Model Integration
**Body:**
## Problem
The `monitor.SlidingWindow` (TCS engine) is implemented, but the Warden/Arbiter loop is not fully constrained by the confidence threshold to prevent self-induced DoS.

## Required Work
- [ ] Bridge `TCSEngine.Evaluate()` to Warden enforcement logic.
- [ ] Implement 'Degraded Mode' FSM transition logic.
- [ ] Create unit tests for Arbiter behavior under $T_C < 0.85$.

## Risk
Medium

---

## Future Roadmap (Phase 2+)

1. **Distributed Control Plane (Phase 2):** Implement Nexus consensus protocols.
2. **AI Reasoning Layer (Phase 3):** Integrate local LLM for correlation and planning.
3. **Game Theory Engine (Phase 4):** Model Stackelberg defense strategies.
4. **Hardware Root of Trust (Phase 5):** TPM 2.0 integration for boot-time measurement.
