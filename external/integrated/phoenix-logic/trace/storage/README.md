# Phoenix Trace: 3-Tier Storage (P0)

Solves the "Lineage Memory Explosion" problem by tiered lifecycle management.

## 1. HOT Tier (In-Memory)
- **Content:** Active processes and recent edges.
- **Format:** Pointer-heavy graph structure for fast traversal (BFS/DFS).
- **Eviction:** When process exits, move to WARM.

## 2. WARM Tier (Local Cache)
- **Content:** Lineage of processes that exited within the last N hours.
- **Format:** Protobuf or MsgPack compressed blobs on disk.
- **Eviction:** After expiry, move to COLD.

## 3. COLD Tier (Long-Term Archive)
- **Content:** "Skeleton Chain" for forensic integrity.
- **Format:** Minified tuples (PID, PPID, StartTime, EndTime, Hash) stored in a compressed archive.
- **Retention:** Permanent for critical nodes (`init`, `auth`, `kernel`).
