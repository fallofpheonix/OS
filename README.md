# PhoenixOS

## Status: v0.1.0-alpha-hardened (STAGE A COMPLETE)

PhoenixOS is a deterministic security-control operating substrate. It implements a thermodynamic approach to system protection, "quenching" high-entropy anomalies into low-entropy, policy-enforced states.

---

## The Hardened Substrate (Stage A Genesis)

The core architecture is now verified and hardened. The system successfully demonstrates:
- **Mechanical Determinism:** Warden FSM enforces state transitions (SAFE -> WATCH -> SUSPICIOUS -> CRITICAL -> COMPROMISED).
- **Thermodynamic Guardrails:** Policy-gated actuation via `redlines.json`.
- **Causal Lineage:** Every event is captured in a Merkle-DAG via the Phoenix Ledger.
- **L7 Neural-Mechanical Bridge:** AI directives are filtered through a deterministic safety layer before actuation.

---

## System Layers (Phoenix Matrix)

| Layer | Component | Status |
| :--- | :--- | :--- |
| **L7** | Phoenix Nexus | **ACTIVE** (Oracle Bridge) |
| **L6** | Phoenix Sentinel | **ACTIVE** |
| **L5.5** | Phoenix Arbiter | **ACTIVE** |
| **L5** | Phoenix Warden | **ACTIVE** (Hardened FSM) |
| **L4** | Phoenix Trace | **ACTIVE** |
| **L3** | Phoenix Monitor | **ACTIVE** |
| **L2/L1**| Phoenix Guard | **ACTIVE** (<100ms Fast-Path) |

---

## Core Principles
1. **Determinism is sacred.** No race conditions. Monotonic logical clocks.
2. **Replay is authoritative.** System state can be reconstructed from the Ledger.
3. **AI is advisory.** The AI (G0DM0D3) proposes; the Substrate (Warden) disposes.

---

## Build & Audit

```bash
# Build the core services
go build ./...

# Run the Stage A Genesis Audit
go run tools/full_system_audit.go
```

## Validation

```bash
# Verify race safety
go test -race ./...

# Run fuzz testing on the Ledger
go test -fuzz ./...
```

---

## Project Roadmap
- **Stage A: Hardening** [COMPLETE]
- **Stage B: Formal Invariants** [PLANNED]
- **Stage C: OS Primitives** [PLANNED]
- **Stage D: Distributed Coordination** [PLANNED]

See `docs/PHASE_STATUS.md` and `STAGE_A_HARDENING_REPORT.md` for details.

