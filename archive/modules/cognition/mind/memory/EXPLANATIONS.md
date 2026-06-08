# PhoenixMemory Retrieval Subsystem

## Scope
Defines the mechanism for grounding AI reasoning in historical truth.

## Architectural Decision: SQLite-Vec Substrate
PhoenixOS has migrated from file-based Obsidian vaults to a programmatic SQLite-Vec substrate. This ensures O(1) retrieval time and O(Log N) vector search performance for semantic grounding.

## Memory Tiers
1. **Working:** Volatile ingress (In-memory map).
2. **Episodic:** Chronological ledger replay (SQLite `episodic_memory`).
3. **Semantic:** Knowledge and causal links (SQLite `semantic_memory` with vector support).
4. **Procedural:** System policies and reaction patterns (SQLite `procedural_memory`).

## Data Integrity
- **Persistence:** All consolidated facts are asynchronously persisted to the local SQLite database.
- **Rollback:** Rollback logic restores state by querying the last known-good Fact from the `StateActive` version slice.
- **Audit:** Every retrieval is tracked to maintain a verifiable lineage between history and AI directives.
