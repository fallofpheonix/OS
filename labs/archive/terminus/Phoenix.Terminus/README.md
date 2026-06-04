---\nStatus: Planned\nImplementation: 5%\nConfidence: Conceptual\n---\n# Phoenix.Terminus — Interface Layer

> Operator interfaces, documentation, external integrations, and the LLM Oracle for PhoenixOS.

## Overview

Phoenix.Terminus contains the interface components that connect PhoenixOS to the outside world. It includes the operator dashboard, documentation vault, external integrations, and the G0DM0D3 LLM Oracle.

**Invariant: All external interfaces must be authenticated and authorized.**

## Sub-Packages

| Package | Purpose | Key Files |
|---------|---------|-----------|
| **PhoenixOS** | Top-level orchestration and deployment | cmd/phoenixd/ |
| **PhoenixDocs** | Canonical documentation vault | architecture/, governance/ |
| **PhoenixDashboard** | Operator UI and visualization | core/, frontend/ |
| **PhoenixExternal** | Third-party integrations | core/logging/ |
| **PhoenixResearch** | Experimental systems | Various |
| **G0DM0D3** | LLM Oracle for cognitive queries | index.html, api/ |

## Dependency Graph

```
Phoenix.Terminus
    ↑
    ├── PhoenixOS (depends on all Nucleus packages)
    ├── PhoenixDocs (standalone documentation)
    ├── PhoenixDashboard (depends on Nucleus/Core)
    ├── PhoenixExternal (standalone integrations)
    ├── PhoenixResearch (experimental)
    └── G0DM0D3 (standalone LLM interface)
```

## Build & Test

```bash
# Go packages
go build ./Phoenix.Terminus/PhoenixOS/...
go test ./Phoenix.Terminus/PhoenixOS/...

# Python packages
cd Phoenix.Terminus/PhoenixDashboard && pip install -r requirements.txt

# Single-file app
open Phoenix.Terminus/G0DM0D3/index.html
```

## Invariants

1. **Authentication required** — All external interfaces
2. **Read-only by default** — Dashboard operations
3. **No sensitive data in UI** — Redaction required
4. **Documentation must be versioned** — All changes tracked

## Related Documents

- [PhoenixOS README](PhoenixOS/README.md)
- [PhoenixDocs README](PhoenixDocs/README.md)
- [G0DM0D3 README](G0DM0D3/README.md)
- [WORKING_MODEL.md](../WORKING_MODEL.md)
