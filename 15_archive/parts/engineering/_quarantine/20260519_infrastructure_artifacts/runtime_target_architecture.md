# Runtime Target Architecture

## Goal
Centralize and standardize runtime environments to reduce sprawl and ensure consistency.

## Target Structure
```
engineering/
runtime/
├── core/           # Authoritative environment for astraeus-core
├── research/       # Shared environment for active research projects
├── experiments/    # Isolated environments for specific experimental runs
├── archive/        # Frozen manifests (no active venvs)
└── shared/         # Common tools (ruff, mypy, etc.)
```

## Environment Specifications
- **Core Runtime**: Python 3.12 (Standardized).
- **Research Runtime**: Python 3.12 + ML Stack (Torch, Transformers).
- **Shared Tools**: Python 3.12 (Isolated via `uv tool` or similar).

## Migration Strategy
1. Collapse `astraeus-core/.venv` into `runtime/core`.
2. Consolidate `ai-system/venv` and other research manifests into `runtime/research`.
3. Clean up root `.venv` and manage via a root `pyproject.toml` in `infrastructure/`.
