### Traversal

- **BFS**: Explore nodes level-by-level using a queue for shortest path in unweighted graphs.
- **DFS**: Explore as deep as possible using recursion/stack for structure and path exploration.

---

### Cycle Detection

- **Directed**: Detect cycles using DFS with recursion stack or Kahn’s in-degree check.
- **Undirected**: Detect cycles using DFS with parent tracking or Union-Find.

---

### Topological Sort

- **Topological Sort (BFS/DFS)**: Order nodes in DAG such that all dependencies come before dependents.
- **Kahn’s Algorithm (BFS in-degree)**: Repeatedly remove nodes with zero in-degree to build order.
- **DFS-based topo sort**: Use post-order DFS stack to generate topological ordering.

---

### Shortest Path

- **Dijkstra**: Find shortest paths in weighted graphs with non-negative edges using a min-heap.
- **Bellman-Ford**: Handle negative weights and detect negative cycles via edge relaxation.
- **Floyd-Warshall**: Compute all-pairs shortest paths using dynamic programming.

---

### Spanning Tree

- **Kruskal**: Build MST by sorting edges and adding smallest non-cycling edges using DSU.
- **Prim’s**: Grow MST from a node by greedily adding minimum edge using a priority queue.

---

### Union-Find (DSU)

- **Union-Find (DSU)**: Track connected components and detect cycles via union and path compression.

---

### Special BFS Variants

- **Bipartite Check**: Use BFS/DFS coloring to verify if graph can be split into two sets.
- **Multi-source BFS**: Start BFS from multiple nodes simultaneously for minimum distance propagation.
- **0-1 BFS**: Use deque to handle edges with weights 0 or 1 in O(V + E).

## Related
- [[Dsa Patterns]]
