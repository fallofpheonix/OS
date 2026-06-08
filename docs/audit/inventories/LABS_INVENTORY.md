# Labs Inventory & Classification Audit

> **Authority Phase**: 4A.7 Labs Classification & Restoration Audit
> **Status**: COMPLETED
> **Last Updated**: 2026-06-04

## 1. Labs Inventory

| Module | Last Activity | Referenced By | Imports | Status | Action |
| :--- | :--- | :--- | :--- | :--- | :--- |
| `labs/formal` | 2026-06-04 | Docs/TLA+ | None (Standalone) | ACTIVE | **KEEP**: Critical for formal verification. |
| `labs/archive/research` | 2026-06-01 | None | Unknown | HISTORICAL | **ARCHIVE**: Speculative AI/ML research. |
| `labs/archive/g0dm0d3` | 2026-06-02 | None | Internal | HISTORICAL | **ARCHIVE**: Debugging/Root utilities. |
| `labs/archive/terminus` | 2026-06-04 | None | Internal | HISTORICAL | **ARCHIVE**: Superseded UI/CLI prototypes. |
| `labs/archive/nucleus` | 2026-06-04 | None | Internal | HISTORICAL | **ARCHIVE**: Legacy core experiments. |
| `platform/crucible` | 2026-06-04 | CLI, UI | Foundation | **TEMPORARY** | **RECLASSIFY**: Move to `assurance/simulation` or `game/crucible` post-audit. |

## 2. Decision Log

- **Labs Ownership**: `labs/` remains the primary location for high-entropy, low-trust research.
- **Formal Verification**: `labs/formal` is identified as the canonical location for TLA+ and formal specifications. It will remain in Labs to prevent polluting the production dependency graph with research scripts.
- **Crucible Reclassification**: Currently resides in `platform/crucible` to resolve the `platform/cli` dependency violation. However, it is architecturally a "Verification & Adversarial Simulation" asset. 
- **Archive Status**: Modules in `labs/archive/` are confirmed unreferenced by the current `go.work` and build pipelines. They are preserved for historical context but isolated from the active substrate.

## 3. Next Steps (Phase 4A.8)

1. Generate a full **Dead Code Register** for the Foundation and Assurance layers.
2. Verify all unreferenced packages in `labs/archive` can be safely removed from the filesystem if storage optimization is required.
3. Perform a **Dependency Truth Audit** to map the exact relationship between `platform/cli` and `platform/crucible`.

---
*Verified against Repository Constitution v1.0.*
