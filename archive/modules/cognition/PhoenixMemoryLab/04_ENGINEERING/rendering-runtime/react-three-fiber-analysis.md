# React-Three-Fiber — Runtime Analysis

Goal: analyze `react-three-fiber` for runtime control-plane semantics (scene lifecycle, scheduling, declarative→imperative boundary), not graphics features.

## Scene Lifecycle
- Declarative scene graph expressed in JSX maps to an imperative Three.js scene. Ownership must be explicit: which layer (React vs Three) owns an object at each lifecycle stage.
- Component lifecycle hooks (mount/unmount) correspond to object creation/destruction; `Canvas` drives the render loop.

## Update Scheduling
- `useFrame` subscribes components to the render loop; React scheduling and three's render loop interact — performance wins come from batching and scheduling updates outside React when appropriate.
- Backpressure: many `useFrame` subscribers can cause CPU/GPU contention; patterns include selective subscription and conditional updates.

## Declarative Runtime Coordination
- The reconciler turns JSX into create/update/delete operations for three objects — this is a canonical declarative→imperative adapter.
- Important pattern: small, self-contained components that communicate through props or external stores rather than implicit global threejs objects.

## State Propagation
- State can be kept in React, external stores (Zustand, Valtio), or per-object refs. Each choice trades locality vs global coordination: global stores risk centralization if used for many objects.

## Topology Preservation
- Scene graph topology must be preserved to keep invariants (transform hierarchies, materials sharing). Reconciler ensures stable identity mapping, but custom mutations can break topology.

## Failure & Recovery Semantics
- Imperative three.js errors (e.g., invalid geometry) can propagate silently; recommended: defensive wrappers that validate state before mutation.

## Coordination Gravity Signals (rendering context)
- Large numbers of shared global stores, cross-component direct mutations, or tightly coupled runtime effects (many components reading same global) indicate rising coordination mass.

## Architectural Weaknesses
- Hidden global mutation via `ref.current` can cause implicit coupling and surprising propagation.
- Overuse of global state stores for scene data leads to centralized update cycles and scheduling bottlenecks.

## Ideas Worth Preserving
- Declarative reconciler approach: clear separation of intent (JSX) and effect (three.js objects).
- `useFrame` subscription but with selective, conditional updates.

## Ideas Worth Avoiding
- Treating three.js objects as uncontrolled global mutable state; avoid wide-reaching direct mutation patterns.

---

Next actions (recommended):
- Identify example components in the fork that show safe reconciler boundaries and extract small architecture notes for `04_ENGINEERING/rendering-runtime/`.
- Measure subscription counts (`useFrame`) in large scenes to quantify scheduling pressure.
