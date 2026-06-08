# Labs Classification Matrix

> **Authority Phase**: 4A.6 Documentation Authority Cleanup
> **Status**: IN-PROGRESS
> **Last Updated**: 2026-06-04

## 1. Labs Classification

| Module | Last Activity | Referenced | Status | Action |
| :--- | :--- | :--- | :--- | :--- |
| `labs/crucible` | 2026-06-03 | YES (CLI) | ACTIVE | **VIOLATION**: Move to `platform/` or `game/` |
| `labs/formal` | 2026-06-02 | YES (Docs) | ACTIVE | KEEP in Labs |
| `labs/g0dm0d3` | 2026-06-02 | NO | ACTIVE | ARCHIVE if unreferenced |
| `labs/research` | 2026-06-01 | NO | ACTIVE | ARCHIVE |
| `labs/archive` | 2026-06-04 | NO | HISTORICAL | DELETE candidates inside |

## 2. Decision Log

- **Crucible**: Currently imported by `platform/cli/main.go`. This is a production boundary violation. Action: Classification as "Strategic Asset" and move out of Labs.
- **Formal**: TLA+ specs are valuable for verification but not imported. KEEP as research.
- **Research/G0dm0d3**: High entropy, low reference. Move to ARCHIVE to reduce cognitive load.

---
**Rule**: Production imports -> Labs MUST BE ZERO.
**Violation**: `platform/cli/main.go` -> `labs/crucible/...`
