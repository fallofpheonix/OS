# Remote Development Flow

## Overview
This document outlines the standard workflow for interacting with modules in a remote-first, cache-friendly manner.

## Workflow

### 1. Identify Need
- Identify the module or repository required for the task.

### 2. Registry Lookup
- Check `remote_execution_registry.yaml` for the module's status and policies.

### 3. Presence Check
- **Local Existence:**
  - If the module exists in its designated `local` path (e.g., `workspace/active/core/` or `modules/editable/`):
    - Proceed directly to **Step 6: Implementation**.

### 4. Cache Acquisition (If not local)
- If the module is NOT local:
  - Clone the repository into `github_cache/` or a temporary directory.
  - Perform a shallow clone (`--depth 1`) if full history is not required.

### 5. Environment Setup
- If modifications require specific dependencies:
  - Use `uv` or `pip` to install dependencies in a virtual environment.
  - Avoid global installs.

### 6. Implementation & Local Validation
- Perform the required edits or investigations.
- Run tests locally within the cache or local path.

### 7. Commit & Push
- If changes were made:
  - Stage and commit changes.
  - Push to the remote repository.

### 8. Cleanup (If cached)
- If the module was cloned into the cache:
  - Remove the cached repository to minimize local footprint.
  - `rm -rf github_cache/<repo_name>`

## Exception: Protected Core
- Protected core repositories (`KEEP_LOCAL` policy) are never removed from their local paths.
- Changes are pushed directly from their permanent locations.
