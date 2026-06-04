---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Contracts — Failure Modes

> Last verified: 2026-06-04

## Known Failure Vectors

| ID | Failure Mode | Mitigation |
|----|--------------|------------|
| C-01 | Unrecognized Event Type | The serializer must fall back to generic JSON. |
| C-02 | Signature Verification Bypass | Enforce signature verification before envelope wrapping. |
| C-03 | Transition Blocked | Actuators must fail-closed and restrict environment rights. |
