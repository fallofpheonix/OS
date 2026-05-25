# ILLEGAL_PATHS.md

## 1. Forbidden Dependencies
The following paths are strictly rejected by the architecture:

- Research → Runtime
- Quantum → Core
- Cognition → Warden
- External → Containment
- Memory → Runtime
- Archive → Runtime

## 2. Enforcement
Any attempt to import or call across these boundaries will result in a build failure or runtime panic.
