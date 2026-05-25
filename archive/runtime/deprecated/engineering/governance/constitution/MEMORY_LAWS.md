# Memory Laws

Rules governing memory tiers, consolidation, and replay integrity.

Sections
--------
- Episodic Integrity
- Semantic Consolidation
- Replay Preservation
- Context Compression
- Memory Decay
- Memory Ownership

Key Laws
--------
- Historical events may be compressed but never silently deleted; compression must preserve causal lineage.
- Semantic memory may be derived from episodic traces but may not overwrite raw episodic data.
- Replay traces must reconstruct prior system state for any checkpointed interval.
- Memory ownership defaults to local runtime; sharing requires explicit transfer and recorded consent.

Next steps
----------
- Define storage schemas for each memory tier and implement retention / decay policies.
