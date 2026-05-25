# Rename Migration Plan (Local)

## Phase 1: Analysis
- [x] Scan all local and remote repositories.
- [x] Classify repositories by category (CORE, SCIENTIFIC, etc.).
- [x] Generate `repo_identity_audit.yaml`.

## Phase 2: Proposal
- [x] Generate `rename_candidates.yaml`.
- [x] Define scientific stack target names in `scientific_registry.yaml`.
- [x] Identify protected repositories (`fallofpheonix`, `forge-agent`, etc.).

## Phase 3: Approval (WAITING)
- User review of `rename_candidates.yaml` and `scientific_registry.yaml`.

## Phase 4: GitHub Rename (PLANNED)
- Execute renames via GitHub UI or API.
- Verify automatic redirects are active.

## Phase 5: Local Update (PLANNED)
- Update local directory names in `workspace/active/`.
- Update Git remotes to point to new URLs.

## Phase 6: Manifest & Registry Update (PLANNED)
- Update `ENGINEERING_MANIFEST.md`.
- Update `github_execution/execution_registry.yaml`.
- Update `module_registry.yaml` and `install_registry.yaml`.
- Update internal dependency manifests within repositories.

## Phase 7: Validation (PLANNED)
- Run connectivity tests.
- Verify install and restore flows.
- Confirm CI/CD workflows are operational.
