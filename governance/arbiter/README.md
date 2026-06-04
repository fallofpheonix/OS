# Phoenix.Arbiter — Governance & Oversight

> Policy mapping, compliance scanning, and authority oversight for PhoenixOS.

## Overview

Phoenix.Arbiter is the governance domain of PhoenixOS. It ensures that the system operates within its defined policy boundaries and provides tools for auditing authority and detecting architectural drift.

## Sub-Packages

| Package | Purpose |
|---------|---------|
| **governance** | Policy definitions and FSM transition rules. |
| **scanner** | System-wide compliance scanning and audit reporting. |
| **mapper** | Authority mapping, lineage tracking, and dependency analysis. |

## Key Principles

1. **Truth of Governance**: Policy must be formally verifiable and grounded in the Genesis block.
2. **Deterministic Oversight**: Scanning and mapping results must be reproducible and tamper-evident.
3. **Zero AI Dependency**: The governance layer must function independently of the AI (Cognition) layer.

## Build & Test

```bash
# Build Arbiter packages
go build ./Phoenix.Arbiter/...

# Test Arbiter packages
go test ./Phoenix.Arbiter/...
```

## Related Documents

- [OWNERSHIP.md](OWNERSHIP.md)
- [ARCHITECTURE.md](../ARCHITECTURE.md)
