# Memory Subsystem

## Primary Responsibility
The Memory subsystem provides the tiered storage substrate for PhoenixOS. It is responsible for multi-version concurrency control (MVCC) of facts, chronological replay, and high-performance persistence via SQLite-Vec.

## System Architecture
1. **Fact Store:** A versioned key-value store where historical states are preserved rather than overwritten.
2. **Tiered Memory:** A 4-layer hierarchy (Working, Episodic, Semantic, Procedural) that manages the lifecycle of context from volatile ingress to permanent persistence.
3. **SQLite Vector Store:** A local persistence layer designed for O(1) retrieval and O(Log N) vector search.

## Tech Stack
- Go (MVCC logic)
- SQLite 3 (Persistence)
- github.com/mattn/go-sqlite3 (Driver)

## AI-Specific Context
- **System Boundaries:** Southbound persistence to disk. Northbound episodic retrieval for AI prompts.
- **Data Flow:** Telemetry -> Ingest (Working) -> Consolidation (Semantic/Episodic) -> SQLite Persistence.
