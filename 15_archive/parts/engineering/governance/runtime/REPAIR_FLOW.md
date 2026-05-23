# Repair Flow

Defines detection, classification, planning, execution, and learning for repairs.

Flow
----
1. Detect failure via event or validator.
2. Classify failure (syntax, dependency, semantic, environment).
3. Plan repair using causal graph and simulation.
4. Simulate repair and evaluate side-effects.
5. Execute repair in staging with full event logging.
6. Validate; on success commit, on failure rollback and escalate.

Constraints
-----------
- Repair retries limited by invariant severity and configured thresholds.
- Repairs must include rollback plans and tests that validate no new invariant violations introduced.

Next steps
----------
- Add example repair scenarios and causal graph examples.
