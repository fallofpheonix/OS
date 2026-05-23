# Rollback Engine

Defines recovery and rollback procedures for failed mutations and catastrophic state errors.

Capabilities
------------
- Snapshot restore
- Event-driven reversal
- Partial rollback (subtree) and full rollback
- Post-rollback validation and quarantine

Risk levels
-----------
- Low: revertable changes with automated tests
- Medium: changes requiring human review after rollback
- High: core identity or constitutional changes — require multi-signer approval and forensic review

Next steps
----------
- Provide concrete rollback examples and automated verification scripts.
