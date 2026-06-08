# Failure Note: Executable Code in `brain/`

## Summary
Executable project files were created under `brain/`, violating the separation rule that cognition lives in `brain/` and executable systems live in `engineering/workspace/`.

## Impact
- Pollutes the vault with runtime artifacts.
- Blurs the boundary between notes and runnable systems.
- Makes future organization and reuse harder.

## Correction
- Moved the project into `~/engineering/workspace/forge-agent/`.
- Renamed the runtime executor to `runtime/shell/executor.py`.

## Prevention
- Create executable systems only under `engineering/workspace/`.
- Treat `brain/` as cognition, documentation, and failure capture only.
