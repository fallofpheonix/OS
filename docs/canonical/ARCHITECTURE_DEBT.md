---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Architecture Debt Register

This file tracks temporary violations and exceptions that are accepted for progress but must be resolved.

## Fields
- Debt ID
- Description
- Affected Boundary
- Risk
- Owner
- Target Resolution

## Current Entries

| Debt ID | Description | Affected Boundary | Risk | Owner | Target Resolution |
| :--- | :--- | :--- | :--- | :--- | :--- |
| AD-001 | Validation tests use proto EventEnvelope while replay engine consumes internal Event | contracts/events, contracts/replay | High | Validation Team | Introduce contract adapter and remove duplicate event path |
| AD-002 | Replay tests expect Reconstruct/currentState APIs that no longer exist | contracts/replay | High | Validation Team | Publish replay contract and update tests to contract API |
| AD-003 | Guard test stub lacks parity with Actuator interface semantics | contracts/security | Medium | Security Team | Align stubs with contract version and add compatibility tests |