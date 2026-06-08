# shadergradient

## Purpose
Study shader composition and gradient-driven animation patterns.

## Core Architecture
- Small shader modules + composition layer; GPU-driven animation timing.

## Interesting Systems
- Gradient math, interpolation strategies, color-space handling, GPU resource lifecycle.

## Reusable Ideas
- Canonical gradient primitives, edge-blend helpers, time-based easing mapped to color ramps.

## Reusable Components
- Small shader snippets, uniform layout conventions, animation-timing helpers.

## Rendering / Runtime Concepts
- Minimal CPU control-plane; offload interpolation to GPU; deterministic time-base.

## Patterns Worth Extracting
- Shader organization, parameterized effects, GLSL function library, packaging for `modules/visual-effects`.

## Weaknesses
- Likely app-specific binding code; limited test coverage; tight coupling to demo UI.

## Integration Opportunities
- Extract GLSL snippets into `08_MODULES/visual-effects/` and document usage patterns in `04_ENGINEERING/rendering-runtime/`.

## Delete / Archive Decision
- Classification: Pattern Source → extract into brain, then delete fork unless heavily modified.
