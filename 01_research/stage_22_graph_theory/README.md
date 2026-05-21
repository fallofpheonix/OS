# Stage 22: Graph Theory in Phoenix

## 1. Core Objective
To represent all OS telemetry as a dynamic, heterogeneous, directed graph of system entity interactions. By executing real-time graph algorithms directly on process lineage, file, and socket access logs, Phoenix isolates lateral movement, detects anomalous communities, and determines the optimal, minimal-impact containment actions during an active attack.

---

## 2. Mathematical Formulations & Entity Definitions

### 2.1. System Graph Formulation
Let $G(t) = (V(t), E(t))$ be a dynamic, heterogeneous directed multi-graph representing the state of the operating system at time $t$.

#### Entities (Vertices $V$):
The set of vertices is partitioned by entity type:
$$V(t) = V_P \cup V_T \cup V_S \cup V_F \cup V_C \cup V_M \cup V_{Cr} \cup V_U \cup V_H$$

Where:
*   $V_P$: **Process** nodes, identified by `(PID, start_time, container_id)`.
*   $V_T$: **Thread** nodes, identified by `(TID, start_time)`.
*   $V_S$: **Socket** nodes, identified by `(protocol, local_ip, local_port, remote_ip, remote_port)`.
*   $V_F$: **File** nodes, identified by `(canonical_path, mount_id)`.
*   $V_C$: **Container** nodes, identified by `(container_id, namespace_set)`.
*   $V_M$: **Memory Region** nodes, representing anonymous maps, shared memory segments, or mapped libraries.
*   $V_{Cr}$: **Credential** nodes, representing UIDs, GIDs, capabilities, or tokens.
*   $V_U$: **User** nodes, representing system and interactive user accounts.
*   $V_H$: **Host** nodes, representing physical or virtual OS environments.

#### Edges ($E$):
Each edge $e \in E(t)$ is a tuple $e = (u, v, l, \tau)$ where $u, v \in V(t)$ are the source and target nodes, $l$ is the edge label, and $\tau \le t$ is the event timestamp.
The label set is defined as:
$$l \in \{\text{spawn}, \text{exec}, \text{read}, \text{write}, \text{rename}, \text{connect}, \text{elevate}, \text{inject}\}$$

---

### 2.2. Graph Algorithms for Incident Detection & Containment

#### PageRank for Asset & Entity Importance:
To estimate the significance of process nodes and flag anomalous activity in highly active processes:
$$PR(u) = \frac{1 - d}{|V|} + d \sum_{v \in \text{In}(u)} \frac{PR(v)}{|\text{Out}(v)|}$$
Where $d = 0.85$ is the damping factor, and $\text{In}(u)$ represents processes reading from or spawning process $u$.

#### Betweenness Centrality ($C_B$) for Containment Target Isolation:
In multi-stage attacks, a compromised server (e.g., a web server) spawns shell code that fetches credentials and invokes a ransomware script. The "bridge" process that connects the entry point to the action-on-object (encryption) represents the high-value containment target. We run betweenness centrality to isolate this bridge:
$$C_B(v) = \sum_{s \neq v \neq t \in V} \frac{\sigma_{st}(v)}{\sigma_{st}}$$
Where $\sigma_{st}$ is the total number of shortest paths from source node $s$ to target node $t$, and $\sigma_{st}(v)$ is the number of those paths passing through $v$.

#### Louvain Clustering for Community Detection:
To group processes and files into functional compartments (e.g., Web Server community, Database community, System Update community). If a process transitions across communities or forms unauthorized edges, it is flagged. Modularity $Q$ is optimized:
$$Q = \frac{1}{2m} \sum_{i,j} \left[ A_{ij} - \frac{k_i k_j}{2m} \right] \delta(c_i, c_j)$$
Where $A_{ij}$ is the adjacency matrix, $k_i$ is the degree of node $i$, $m$ is total edges, and $\delta$ is the Kronecker delta indicating membership in communities $c_i, c_j$.

#### Traversal and Temporal Paths:
*   **BFS/DFS:** Used for real-time forward and backward lineage tracing (blast radius and root cause analysis).
*   **Temporal Shortest Path:** Shortest path computation restricted by causality ($\tau_1 \le \tau_2 \le \dots \le \tau_n$), ensuring no backward-in-time flows are analyzed.

---

## 3. Subsystem Mapping
*   **Engine Directory:** `07_security/graph/`
*   **Components:**
    *   `incident_graph/`: In-memory dynamic graph representing live system interactions.
    *   `attack_path/`: Algorithmic engine that traces path transitions and calculates betweenness centrality to isolate bridges.
    *   `community_detection/`: Periodic background worker running Louvain modularity clustering on active subgraphs.

---

## 4. Experiment Backlog

### Experiment R022: Ransomware Graph Location & Containment
*   **Objective:** Replay simulated ransomware (e.g., a rapid write/rename loop spanning multiple subdirectories) and construct the incident DAG. Calculate centrality to locate the attack center.
*   **Telemetry Source:** eBPF `sched_process_exec`, `vfs_write`, `vfs_rename`.
*   **Metrics:** 
    *   Graph build latency $\le 1.5$ ms per event.
    *   Bridge isolation accuracy $\ge 99\%$.
*   **Integration Target:** `07_security/graph/incident_graph`.
*   **Output Artifact:** Labeled Incident DAG.
