# RFC: Phoenix Trace (L4)

## 1. Data Model
Nodes: Process (PID), File (Path), Socket (IP:Port).
Edges: `forked`, `executed`, `wrote`, `connected`.

## 2. Algorithms
- **PageRank / Betweenness:** Identify critical "pivot" processes.
- **Reachability:** Trace root compromise from a leaf event.

## 3. Storage
In-memory Adjacency List with RWMutex protection.
