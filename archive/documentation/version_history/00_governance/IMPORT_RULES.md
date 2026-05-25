# PhoenixOS: Import Rules

**Goal:** Prevent dependency cycles and preserve architectural tiers.

1. **Substrate Layer:** `contracts`, `monitor`. Cannot import higher layers.
2. **Truth Layer:** `truth`, `state`. Imports `contracts`.
3. **Control Layer:** `warden`, `arbiter`. Imports `truth`, `state`, `contracts`.
4. **Cognitive Layer:** `ai`, `nexus`, `memory`. **CURRENTLY BLOCKED (PHASE LOCK)**.

Enforcement is strictly managed via the CI script: `05_tools/validate_imports.py`
