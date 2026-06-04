---\nStatus: Implemented\nImplementation: 100%\nConfidence: Proven\n---\n# Events — Failure Modes

> Last verified: 2026-06-04

## Known Failure Vectors

| ID | Failure Mode | Mitigation |
|----|--------------|------------|
| E-01 | JSON Parsing Error | Enforce strict schema checks during ingestion. |
| E-02 | Clock Drift / Invalid Logical Time | Reject events with logical timestamps behind the ledger head. |
