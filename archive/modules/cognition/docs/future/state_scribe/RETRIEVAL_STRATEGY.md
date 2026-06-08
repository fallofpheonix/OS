---
Status: Planned
Implementation: 10%
Confidence: Conceptual
---
# State Scribe — Retrieval Strategy

Queries vector stores using hybrid search (keyword BM25 + cosine vector similarity).

## Metrics
$$\text{Score} = \alpha \cdot \text{CosineSimilarity} + (1 - \alpha) \cdot \text{BM25Score}$$
