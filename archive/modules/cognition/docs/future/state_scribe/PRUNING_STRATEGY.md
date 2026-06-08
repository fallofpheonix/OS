---
Status: Planned
Implementation: 5%
Confidence: Conceptual
---
# State Scribe — Memory Pruning Strategy

Prunes low-utility memories to conserve context limits.

## Rules
- Prune memory elements with utility score $< 0.3$.
- Keep all files containing forensic ledger commit keys.
