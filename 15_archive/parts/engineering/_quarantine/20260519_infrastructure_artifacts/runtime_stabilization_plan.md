# Runtime Stabilization Plan

## Phase RS2: Runtime Collapse
- **Goal**: Move all active virtual environments into the centralized `runtime/` directory.
- **Actions**:
  - Relocate `astraeus-core/.venv` -> `runtime/core`.
  - Relocate `ai-system/venv` -> `runtime/research`.
  - Update all paths in scripts and manifests.

## Phase RS3: Dependency Freeze
- **Goal**: Standardize package versions across merged runtimes.
- **Actions**:
  - Resolve conflicts in `chromadb` and `pydantic`.
  - Generate a unified `uv.lock` for the `core` runtime.

## Phase RS4: Shared Runtime
- **Goal**: Enable sharing of heavy packages (torch, etc.) between research environments.
- **Actions**:
  - Explore `uv` workspace features for shared caches.
  - Implement a "base" research layer.

## Phase RS5: Validation
- **Goal**: Ensure all systems function correctly after relocation.
- **Actions**:
  - Run full test suite for `astraeus-core`.
  - Validate `ai-system` ingestion pipelines.
  - Verify CLI tool functionality in the new root env.
