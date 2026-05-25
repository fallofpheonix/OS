# Runtime Integrity Report

| runtime | exists | state | rollback | status |
| :--- | :--- | :--- | :--- | :--- |
| core | YES | LOCKED | VALID | PASS |
| research | YES | FROZEN | VALID | PASS |
| shared | YES | LIVE | VALID | PASS |
| archive | YES | DEPRECATED | VALID | PASS |

## Verification Details

- **Core Runtime**: Path `runtime/core` exists. State matches `shared_manifest.yaml`.
- **Research Runtime**: Path `runtime/research` exists. State matches `shared_manifest.yaml`.
- **Shared Runtime**: Path `runtime/shared` exists. State matches `shared_manifest.yaml`.
- **Archive**: Archive directory and manifests (`archive_inventory.md`, `runtime/archive_manifest_index.json`) are present.
- **Rollback**: Rollback paths verified via `restore_manifest.json` (quarantine) and archive indices.
