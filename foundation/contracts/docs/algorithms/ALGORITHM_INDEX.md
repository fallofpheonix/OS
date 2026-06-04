---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Contracts — Algorithm Index

> Last verified: 2026-06-04

Contracts define the interface specification for core algorithms. The implementation details are deferred to the implementing layers, but the contracts enforce their signatures.

## 1. Event Hashing Specification
- **Interface**: `EventEnvelope`
- **Output**: 256-bit SHA-256 hash.
- **Verification Rule**: The hash must match the payload, trace hash, state before/after, and sequence number.

## 2. Dynamic Escalation Logic
- **Interface**: `Actuator`
- **Inputs**: `ContainmentLevel` and `Escalation` reason.
- **Complexity**: $O(1)$ state modification.
