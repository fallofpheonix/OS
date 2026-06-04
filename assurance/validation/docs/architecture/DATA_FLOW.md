---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Validation — Data Flow

> Last verified: 2026-06-04

This document defines the validation verification flows.

## Staged Verification Pipeline Flow

```mermaid
graph TD
    Trigger[Validation Run Trigger] -->|Run stage 1| Unit[Unit Tests]
    Unit -->|Pass| Invariant[Invariants & Proofs]
    Invariant -->|Pass| Determinism[Determinism Auditor]
    Determinism -->|Pass| Exploits[Security Exploit Simulations]
    Exploits -->|Pass| Report[Compliance Report Generated]
    Exploits -->|Fail| Alert[Escalate Failure Alert]
```

## Audit Log Assertion Flow
- Test orchestrators parse output JSON logs from the kernel telemetry.
- Assertions verify that event causality maps correctly to ledger hashes.
