---
Status: Planned
Implementation: 15%
Confidence: Conceptual
---
# PhoenixMind — Memory Model

Describes the partitioning of knowledge within the cognitive loop.

## Memory Classes

1. **Episodic Memory**: Raw logs of execution runs (what occurred and when).
2. **Semantic Memory**: Knowledge base containing libraries, system specifications, and tool contracts.
3. **Working Memory**: In-context task tracker mapping the current step inside execution plans.

## Storage Hierarchy

```mermaid
graph TD
    Working[Working Memory: Context In-Flight] -->|Prune| Semantic[Semantic Memory: Vector DB Embeddings]
    Semantic -->|Seal| Forensic[Forensic Memory: Ledger Log Blocks]
```
