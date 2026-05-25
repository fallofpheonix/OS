# PhoenixOS Contract Matrix

## Overview
The Contract Matrix defines the formal interfaces and semantic versioning rules for inter-component communication in PhoenixOS. This ensures that L1-L7 layers can evolve independently while maintaining system-wide determinism and safety.

## Active Contracts

| Contract ID | File | Description | Stability |
|-------------|------|-------------|-----------|
| VERSION | version.go | Semantic versioning structures and compatibility logic. | STABLE |
| POLICY | policy.go | Policy evaluation and enforcement interfaces. | EXPERIMENTAL |

## Interface Evolution Rules
To maintain the "Deterministic Cybernetic Security Runtime" integrity, the following rules apply to all contract changes:

1. **Additions are permitted:** New methods can be added to interfaces if they don't break existing implementations.
2. **Deletions are restricted:** Methods can only be removed in a MAJOR version bump.
3. **Signatures are immutable:** Once a method is part of a STABLE contract, its signature (parameters and return values) cannot change.
4. **Semantic Versioning:**
   - **MAJOR:** Incompatible API changes.
   - **MINOR:** Functionality added in a backwards-compatible manner.
   - **PATCH:** Backwards-compatible bug fixes.
5. **No Direct Internal Imports:** Components must depend on `phoenix_os/contracts/` interfaces rather than internal concrete implementations.
