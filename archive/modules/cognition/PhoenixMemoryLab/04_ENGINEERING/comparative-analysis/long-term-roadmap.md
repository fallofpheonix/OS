# Long-Term Ecosystem Roadmap

This document captures the desired evolution of forks into your engineering ecosystem, the states they pass through, extraction timing, deletion rules, and the final target structure.

## Six-State Evolution
1. Fork — initial clone for study.
2. Research Specimen — read/scan/experiment, no copying.
3. Pattern Source — canonical patterns documented in `pattern-library`.
4. Architectural Compression — distilled principles stored in `control-plane-analysis` and `rendering-runtime`.
5. Selective Extraction — reimplement small, constrained modules inside your `modules/` when and only when semantic convergence is proven.
6. Original Ecosystem Capability — modules live in `workspace/modules/*` under your ownership.

## When to Extract
- Only after semantic convergence (multiple repos agree on the same pattern or you have measured stable reuse pressure).
- Extraction triggers:
  - Repeated operational need across two or more projects
  - Clear mapping to your locality and containment invariants
  - Small, well-defined API that can be reimplemented without importing upstream architecture debt

## When to Delete Forks
- Delete or archive forks when:
  - All value has been converted to brain notes and reimplemented modules
  - The fork contains no unique evolution you plan to maintain
  - It is not the active canonical source for an evolved module

## How These Repos Connect to Your Runtime
- `agi` → observability & coordination patterns, trace semantics, pulse/leader analysis
- `react-three-fiber` → scheduler semantics, declarative→imperative boundary patterns
- `liquid-*` + `shadergradient` → visual runtime primitives and effect composition ideas (stay as pattern sources until validated)
- `the_well` → external dataset resource; keep as documented resource, not a code fork unless running benchmarks

## Final Target Structure (high-level)
- `brain/`
  - `control-plane-analysis/`
  - `rendering-runtime/`
  - `pattern-library/`
  - `comparative-analysis/`
  - `failure-library/`
- `workspace/`
  - `forge-agent/`
  - `runtime-lab/`
  - `render-lab/`
  - `governance-lab/`
- `modules/`
  - `runtime-governance/`
  - `trace-semantics/`
  - `visual-effects/`
  - `control-plane-boundaries/`

## One-Line Rules
- Do not absorb implementations; compress semantics.
- Delay extraction until patterns stabilize.
- Measure coordination gravity before adding a shared service.

## Immediate Next Steps (post-analysis)
1. Run snapshot concentration metrics on `agi` snapshots to quantify coordination gravity.
2. Measure `useFrame` subscription density in large `react-three-fiber` examples.
3. Produce a small report listing candidate patterns and their evidence strength (weak/medium/strong) before extracting.
