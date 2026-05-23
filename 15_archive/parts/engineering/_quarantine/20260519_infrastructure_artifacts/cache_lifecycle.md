# Cache Lifecycle Model

## Overview
The cache lifecycle governs the temporary existence of repositories in the local environment, ensuring a minimal footprint while enabling full development capabilities.

## Lifecycle States

### 1. EMPTY
- **Description:** The repository exists only on GitHub. No local files are present in the cache.
- **Trigger:** Initial state or after a successful `PURGED` event.

### 2. CLONED
- **Description:** The repository has been cloned into `github_cache/` or `remote_runtime/cache/`.
- **Trigger:** Lookup failure in local directories for a `CLONE_ON_DEMAND` repo.
- **Action:** `git clone --depth 1 <repo_url> <cache_path>`

### 3. EDITING
- **Description:** Local modifications are being made within the cached repository.
- **Trigger:** User or agent begins work on the cached files.
- **Action:** File writes, dependency installs (into temp venv).

### 4. PUSHED
- **Description:** Changes have been committed and successfully pushed to the remote origin.
- **Trigger:** Completion of work and validation.
- **Action:** `git commit` followed by `git push`.

### 5. PURGED
- **Description:** The cached repository is removed from the local filesystem.
- **Trigger:** Successful `PUSHED` state or expiration of a session.
- **Action:** `rm -rf <cache_path>`

## Flow Diagram
```text
[EMPTY] -> (Lookup) -> [CLONED] -> (Work) -> [EDITING] -> (Success) -> [PUSHED] -> (Cleanup) -> [PURGED] -> [EMPTY]
```

## Policy Integration
| Policy | Lifecycle Enabled | Notes |
|---|---|---|
| `KEEP_LOCAL` | NO | Permanent local presence. |
| `CLONE_ON_DEMAND` | YES | Standard lifecycle apply. |
| `CACHE_TEMP` | YES | Aggressive purging after use. |
| `INSTALL_ONLY` | PARTIAL | Clone only for build/install, then purge. |
| `ARCHIVE_ONLY` | NO | Managed by archive restoration logic. |
