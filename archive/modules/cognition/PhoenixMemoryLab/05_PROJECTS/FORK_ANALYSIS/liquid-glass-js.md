# liquid-glass-js

## Purpose
Glass-like distortion and layered effect utilities for web contexts.

## Core Architecture
- JS utilities + small shaders; effect layering with alpha/blur interplay.

## Interesting Systems
- Compositing order, blending modes, performance fallbacks for low-power devices.

## Reusable Ideas
- Distortion primitives, layered-effect pipeline, graceful degradation patterns.

## Reusable Components
- Distortion functions, blur+alpha composition helpers, canvas/GL interop helpers.

## Rendering / Runtime Concepts
- Effect layering as composable primitives; adapter layer for DOM vs WebGL.

## Patterns Worth Extracting
- Adapter interface (DOM↔GL), composable effect chain API, progressive enhancement.

## Weaknesses
- Likely lacks clear packaging; tests and docs thin.

## Integration Opportunities
- Extract into `modules/visual-effects` as `glass` submodule; document composition API in `04_ENGINEERING/pattern-library`.

## Delete / Archive Decision
- Classification: Reusable Component Source → extract code into your modules, then delete or convert fork into private module if heavy changes required.
