# Branching Laws

Governance for cognitive branches, isolation, and merging.

Sections
--------
- Branch Creation
- Branch Isolation
- Merge Rules
- Branch Replay
- Branch Pruning
- Branch Ranking

Key Laws
--------
- Every branch must preserve full causal history and be independently replayable.
- Branches that fail replay or violate critical invariants are quarantined and require remediation before merge.
- Merges into stable branches must include invariant validation, replay verification, and a rollback plan.

Next steps
----------
- Add automated branch trust scoring and CI-level merge blockers based on replay verification.
