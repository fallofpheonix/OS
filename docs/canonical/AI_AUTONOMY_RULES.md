---
Status: Draft
Implementation: 0%
Confidence: High
---
# AI Autonomy Rules

> This document defines the limits and safety constraints for AI-driven logic and self-evolution.

## 1. Actuation Gating

- **Intent-Only:** AI models (Cognition) can only generate `Intent` objects. They have no direct access to system syscalls or actuators.
- **Warden Enforcement:** Every intent is validated against the AI's current `CapabilityToken` by the Warden.
- **Shadow Mode Default:** In production-intent deployments, **Shadow Mode is ENABLED by default**. Any AI-driven actuation that passes validation will be recorded in the Ledger but will NOT have physical substrate effect until Shadow Mode is explicitly disabled. This ensures a "forensics-first" validation period for all new autonomy logic.

## 2. Self-Evolution Gating

- **No Direct Modification:** AI cannot modify the code of the Governance or Assurance layers.
- **Verification Pipeline:** Proposed code changes must pass the full verification suite (TLA+, Chaos, Invariants) before being ledgered for promotion.

--- 
*Refer to [docs/MASTER_INVARIANTS.md](../MASTER_INVARIANTS.md) for global autonomy laws (INV-003).*
