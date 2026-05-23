# Module Import Plan

## Overview
This plan defines the methods for importing and managing modules within the ecosystem, prioritizing remote-first access and localized editing.

## Methods

### 1. Git Clone (Standard)
- **Usage:** For full module access and permanent local presence.
- **Target:** `workspace/active/core/` for protected repos, `modules/` for general modules.
- **Workflow:** `git clone <repo_url> <target_path>`

### 2. Git Sparse-Checkout
- **Usage:** When only specific subdirectories of a large repository are needed.
- **Workflow:**
  ```bash
  git clone --no-checkout <repo_url>
  git sparse-checkout set <dir1> <dir2>
  git checkout
  ```

### 3. Pip Install -e (Editable)
- **Usage:** To install local modules in a way that changes are immediately reflected in the environment.
- **Workflow:** `pip install -e <path_to_module>`

### 4. UV Add
- **Usage:** For modern dependency management using `uv`.
- **Workflow:** `uv add --editable <path_to_module>` or `uv add <package_name>`

### 5. Temporary Cache Clone
- **Usage:** For quick modifications or inspections without permanent storage.
- **Workflow:**
  1. Clone to `github_cache/` or `/tmp/`.
  2. Perform operations.
  3. Push changes if needed.
  4. Remove clone.

### 6. Future Submodule
- **Usage:** For tight coupling between parent and child repositories.
- **Workflow:** `git submodule add <repo_url> <path>`

## Selection Criteria
| Priority | Method | Trigger |
|---|---|---|
| 1 | `KEEP_LOCAL` | Protected Core Repos |
| 2 | `CLONE_ON_DEMAND` | Active development on non-core modules |
| 3 | `INSTALL_ONLY` | Production-only dependencies |
| 4 | `ARCHIVE_ONLY` | Legacy or reference material |
