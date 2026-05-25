# GitHub Pull Model

## 1. Request
- Trigger: Module import or explicit pull request.
- Input: `[REPO_NAME]`.

## 2. Registry Lookup
- Query `github_execution/execution_registry.yaml`.
- Identify `policy` and `local` path.

## 3. Decision Logic

### Policy: KEEP_LOCAL
- **Action:** EDIT directly.
- **Verification:** Ensure local directory exists and is clean.

### Policy: CLONE_ON_DEMAND
- **Check Cache:** Does a valid cache exist in `github_cache/`?
- **YES:** Restore from cache to `local` path.
- **NO:** Perform `git clone` to `local` path.
- **Install:** Execute `install_mode` (e.g., `uv add --editable`).

### Policy: CACHE_TEMP
- **Check Cache:** Does a valid cache exist?
- **YES:** Restore to temporary workspace.
- **NO:** `git clone --depth 1` to temporary workspace.

### Policy: INSTALL_ONLY
- **Action:** Package manager install (e.g., `uv add git+url`).
- **Optimization:** Use cache if available to avoid network call.

## 4. Finalization
- Update `execution_metrics.yaml`.
- Log pull event in `github_execution/telemetry/`.
