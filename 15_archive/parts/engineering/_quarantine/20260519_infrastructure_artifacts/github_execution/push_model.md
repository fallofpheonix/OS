# GitHub Push Model

## 1. Preparation
- **Edit:** Perform code changes in the `local` path defined in the registry.
- **Identify:** Determine target branch and upstream remote.

## 2. Pre-Push Validation
- Execute relevant checks from `github_execution/github_validation.yaml`.
- Ensure all tests pass (if applicable).
- Check `protection_rules.yaml` for immutability constraints.

## 3. Commit & Push
- `git add .`
- `git commit -m "[STATED_INTENT]"`
- `git push origin [BRANCH]`

## 4. Post-Push Synchronization
- Update `github_execution/execution_registry.yaml` with new commit hash/state.
- **Cache Invalidation:** Purge local cache for the repository to ensure next pull is fresh.
- Update `github_execution_report.md`.

## 5. Telemetry
- Log push metrics in `github_execution/telemetry/`.
- Update `active_repos` count in `execution_metrics.yaml`.
