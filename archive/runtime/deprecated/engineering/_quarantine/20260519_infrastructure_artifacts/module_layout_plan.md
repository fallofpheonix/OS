# Module Layout Plan

This plan defines the staging structure for the `modules/` directory, which will house all shared and installable components across the ecosystem.

## Target Structure

```
modules/
├── installable/   # Stable, third-party or internal packages for production use
├── editable/      # Internal modules under active development (pip install -e)
├── runtime/       # Modules required for the core Astraeus runtime
├── research/      # Modules specifically for research experiments and models
└── cache/         # Local cache for module artifacts and pre-built binaries
```

## Staging Rules

1. **Installable**: Only modules with a valid `setup.py` or `pyproject.toml` and passing tests.
2. **Editable**: Symbolic links or local clones of core modules that require frequent cross-repo updates.
3. **Runtime**: Hardcoded dependencies for the bootstrap and control-plane.
4. **Research**: Volatile modules that may be versioned but are not strictly "production" ready.
5. **Cache**: Ephemeral data; should be excluded from git.

## Next Steps (Phase G4)

- Create directory structure.
- Initialize `__init__.py` files where necessary.
- Prepare `requirements-modules.txt` for bulk mapping.
