# liquid-logo

## Purpose
Fluid logo visuals and interaction choreography.

## Core Architecture
- Renderer + shader-based distortion + interaction hooks.

## Interesting Systems
- Interaction→shader parameter mapping, LOD for quality/performance, incremental morphing.

## Reusable Ideas
- Interaction-to-shader mapping, stateful visual transitions, deterministic morph easing.

## Reusable Components
- Distortion controllers, bezier-timed animation drivers, hit-test → visual-response glue.

## Rendering / Runtime Concepts
- Prioritize cheap CPU input processing and delegate heavy blending to GPU.

## Patterns Worth Extracting
- Input-driven visual state machine, reusable distortion controller API.

## Weaknesses
- Demo-centric structure, coupling between UI and rendering code.

## Integration Opportunities
- Reimplement controllers as small modules in `modules/visual-effects` and reference in `04_ENGINEERING/pattern-library/rendering-patterns`.

## Delete / Archive Decision
- Classification: Pattern Source → extract ideas + small controllers, then archive/delete fork.
