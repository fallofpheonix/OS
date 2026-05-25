# Evolution Constraints

Constraints that limit adaptive evolution to preserve stability and identity.

Sections
--------
- Stability Preservation
- Entropy Bounds
- Goal Preservation
- Identity Continuity
- Branch Explosion Limits

Key Constraints
---------------
- Adaptive policies may optimize performance but must not weaken constitutional constraints.
- System-wide entropy (measured by drift metrics) must remain below configurable thresholds; otherwise adaptations are suspended.

Next steps
----------
- Define drift metrics and implement automated drift detectors in `OBSERVABILITY.md`.
