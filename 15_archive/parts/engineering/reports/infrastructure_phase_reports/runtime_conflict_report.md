# Runtime Conflict Report

## Python Version Mismatch
- **Core (astraeus-core)**: Python 3.13.12
- **Research (ai-system)**: Python 3.14.3
- **Shared (root)**: Python 3.14.3
- *Impact*: Incompatibility if runtimes are shared or if core depends on 3.14 features (unlikely as it is on an older version).

## Package Overlap & Version Drift
- **requests**:
  - Shared/Research: 2.34.0
  - Core: 2.34.1
- **tokenizers**:
  - Shared/Research: 0.22.2
  - Core: 0.23.1
- **pip**:
  - Shared/Research: 26.1.1
  - Core: 25.3
- **Other Overlaps**: Significant overlap in common utilities (pydantic, anyio, etc.), but versions generally match.

## Lock Incompatibility
- `astraeus-core` uses `uv.lock`.
- Root and `ai-system` do not have visible lock files in their respective directories, though root might be managed by a global `uv` state.

## UV Drift
- Root `.venv` explicitly references `uv = 0.11.13` in `pyvenv.cfg`.
- Others do not, suggesting manual `venv` creation or different tooling.
