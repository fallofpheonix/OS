# Hydration Flow Simulation

This document simulates the dynamic hydration process for repositories within the empty skeleton structure.

## Scenario 1: Need `physics`
1.  **Registry Lookup:** Query `github_execution/execution_registry.yaml` for `physics`.
2.  **Policy Identification:** `policy: CACHE_TEMP`.
3.  **Cache Check:** Consult `github_execution/cache_registry.yaml`.
    - **Result:** Cache exists at `github_cache/physics`.
4.  **Hydration:** Restore cache to `workspace/active/physics`.
5.  **Runtime Attachment:** The core runtime detects the local presence and binds the scientific layer.

## Scenario 2: Need `simulation`
1.  **Registry Lookup:** Query registry for `simulation`.
2.  **Policy Identification:** `policy: CACHE_TEMP`.
3.  **Cache Check:** 
    - **Result:** Cache missing or expired.
4.  **Remote Fetch:** Perform `git clone --depth 1 https://github.com/fallofpheonix/simulation.git` into `workspace/active/simulation`.
5.  **Cache Update:** Populate `github_cache/simulation` with the new clone.
6.  **Readiness:** Project is now available for execution.

## Scenario 3: Need `RL` (Future System)
1.  **Registry Lookup:** Registry found in `scientific_preload_registry.yaml`.
2.  **Policy Identification:** `policy: CLONE_ON_DEMAND`.
3.  **Hydration:** Perform full clone to `research/rl-infrastructure`.
4.  **Installation:** `uv add --editable research/rl-infrastructure`.
5.  **Future Import:** Core systems can now import and utilize the RL modules.

## Summary
The hydration model relies on a tiered approach:
- **Local Active:** Always present (Core).
- **Cached:** Rapid restoration from local `github_cache/`.
- **On-Demand:** Fetched from GitHub when needed.
