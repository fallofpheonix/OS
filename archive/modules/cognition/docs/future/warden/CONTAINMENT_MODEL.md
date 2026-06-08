---
Status: Partial
Implementation: 45%
Confidence: Conceptual
---
# Warden Guard Rails — Containment namespaces

Uses Linux kernel namespaces to enforce isolation blocks.

## Isolation Scopes
- **Mounts**: Unshared root directory.
- **Net**: No egress routes except to consensus sockets.
- **Pid**: Root isolation inside the target sandbox context.
