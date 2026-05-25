# PhoenixOS: Import Rules

**Goal:** Prevent dependency cycles and preserve tiers.

1. **Substrate Layer:** `contracts`, `monitor`. No upward imports.
2. **Truth Layer:** `truth`, `state`. Imports `contracts`.
3. **Control Layer:** `warden`, `arbiter`. Imports `truth`, `state`, `contracts`.
4. **Cognitive Layer:** `ai`, `nexus`, `memory`, `swarm`, `sentinel`, `physics`, `learning`. **BLOCKED**.
