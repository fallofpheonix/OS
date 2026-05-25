# External Repository Registry & Merge Policies

This document registry tracks all external code imported and modified under the PhoenixOS project.

## Registry

1. **go-sqlite3**
   - **Purpose:** SQL database storage for L4 Trace process graphs.
   - **Version:** v1.14.22 (pinned).
   - **License:** MIT.
   - **Fork Status:** Unmodified upstream dependency.
   - **Risk Level:** Low.
   - **Replacement Plan:** Fallback to raw flat files or badger key-value store if memory overhead becomes restrictive.

## Merge Policy
- Upstream packages must be vendored or pinned to exact commit SHA-256 hashes.
- Direct code changes to external libraries are strictly prohibited without an RFC documenting safety impact and fork strategy.
