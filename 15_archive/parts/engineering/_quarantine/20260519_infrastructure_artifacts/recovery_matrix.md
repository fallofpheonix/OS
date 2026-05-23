# Recovery Matrix Simulation

## Failure Scenarios & Recovery Paths

| Scenario | Primary Failure | Fallback Path | Archive Path | Manual Recovery |
| --- | --- | --- | --- | --- |
| **Repo Missing** | Registry points to path that does not exist. | Trigger `PULL_MODEL` to re-clone from GitHub. | Check `archive/github_synced/` for local copy. | `git clone [GITHUB_URL]` |
| **Cache Missing** | `github_cache/` entry is empty or corrupt. | Bypass cache; perform full `git clone` to target path. | N/A | Clear cache and re-run hydration. |
| **Remote Unreachable** | GitHub is down or network is disconnected. | Attempt to restore from `archive/` or `backups/`. | Use `archive/deprecated/` snapshots. | Verify connectivity and proxy settings. |
| **Restore Failure** | Cache/Archive copy is corrupt or incompatible. | Purge local artifacts and re-clone from remote. | N/A | Rebuild environment from scratch. |
| **Divergence Lock** | Local core repo has unpushed changes that conflict with remote. | Stash local changes, pull remote, and re-apply/merge. | Use `git reflog` to recover lost commits. | Manual merge and conflict resolution. |

## Recovery Priority
1.  **Level 1 (Automatic)**: Pull from GitHub.
2.  **Level 2 (Semi-Automatic)**: Restore from local `github_cache/` or `archive/`.
3.  **Level 3 (Manual)**: Re-clone, rebuild, and re-synchronize.
