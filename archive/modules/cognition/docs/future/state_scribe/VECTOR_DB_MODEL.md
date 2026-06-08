---
Status: Planned
Implementation: 10%
Confidence: Conceptual
---
# State Scribe — Vector DB Model

Defines memory storage properties.

## Database Schemas
- **Index**: HNSW (Hierarchical Navigable Small World).
- **Dimension**: 1536-dimensional float vector embeddings.
- **Fields**:
  - `embedding`
  - `timestamp`
  - `logical_tick`
  - `action_type`
  - `text_contents`
