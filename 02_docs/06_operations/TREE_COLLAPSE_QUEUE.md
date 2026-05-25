# TREE_COLLAPSE_QUEUE.md

## 1. Objective
Map and execute the collapse of the repository into the target 8 primary roots.

## 2. Collapse Queue

| Path | Type | Duplicate | Merge Target | Decision | Status |
| :--- | :--- | :--- | :--- | :--- | :--- |
| `phoenix_os/event_bus` | Directory | YES | `phoenix_os/bus` | **MERGED** | **COMPLETE** |
| `tests/unit/*_test.go` | Files | NO | `tests/unit/` | **KEEP** | **ACTIVE** |
| `tools/build/bin/` | Directory | YES | `build/bin/` | **ARCHIVE** | **OPEN** |
| `02_docs/identity/axioms.md` | File | YES | `02_docs/00_identity/SYSTEM_IDENTITY.MD` | **MERGE** | **COMPLETE** |
| `archive/custom/control_plane` | Directory | YES | `phoenix_os/nexus_coordination` | **REVIEW** | **OPEN** |

## 3. Target Tree
- `phoenix_os/`
- `tests/`
- `tools/`
- `research/`
- `external/`
- `experimental/`
- `archive/`
- `02_docs/`
