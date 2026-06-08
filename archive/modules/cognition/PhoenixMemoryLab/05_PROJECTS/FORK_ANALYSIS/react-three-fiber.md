# react-three-fiber

## Purpose
Understand declarative 3D rendering semantics and scene-graph control-plane.

## Core Architecture
- React reconciler for three.js; lifecycle mapping; resource management.

## Interesting Systems
- Reconciliation → imperative render updates, hooks for resource lifecycles, scheduling and batching.

## Reusable Ideas
- Declarative→imperative boundary rules, scheduling patterns, resource pooling.

## Reusable Components
- Example reconciler patterns, resource hooks, scene-graph traversal utilities.

## Rendering / Runtime Concepts
- Scene graph ownership, deterministic render loop integration with host framework.

## Patterns Worth Extracting
- Declarative resource hooks, minimal reconciler patterns, backpressure handling for large scenes.

## Weaknesses
- Large codebase; careful selective extraction recommended to avoid licensing/maintenance burden.

## Integration Opportunities
- Capture high-level patterns into `04_ENGINEERING/rendering-runtime/` and example reimplementations in `08_MODULES/visual-runtime/`.

## Delete / Archive Decision
- Classification: Knowledge + Pattern Source → do not keep fork; extract architecture notes and small helper reimplementations.
